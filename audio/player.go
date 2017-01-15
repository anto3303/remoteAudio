package audio

import (
	"fmt"
	"time"

	"github.com/dh1tw/gosamplerate"
	"github.com/dh1tw/opus"
	"github.com/dh1tw/remoteAudio/events"
	"github.com/gordonklaus/portaudio"
	"github.com/spf13/viper"
	ringBuffer "github.com/zfjagann/golang-ring"
)

//PlayerSync plays received audio on a local audio device
func PlayerSync(ad AudioDevice) {

	portaudio.Initialize()
	defer portaudio.Terminate()

	//out doesn't need to be initialized with a fixed buffer size
	//since the slice will be copied from the incoming data
	//and will therefore replay any buffer size

	var deviceInfo *portaudio.DeviceInfo
	var err error

	if ad.DeviceName == "default" {
		deviceInfo, err = portaudio.DefaultOutputDevice()
		if err != nil {
			fmt.Println("unable to find default playback sound device")
			fmt.Println(err)
			return // exit go routine
		}
	} else {
		if err := ad.IdentifyDevice(); err != nil {
			fmt.Printf("unable to find recording sound device %s\n", ad.DeviceName)
			fmt.Println(err)
			return
		}
	}

	ad.out = make([]float32, 500000)

	// streamParm := portaudio.LowLatencyParameters(nil, deviceInfo)

	streamDeviceParam := portaudio.StreamDeviceParameters{
		Device:   deviceInfo,
		Channels: ad.Channels,
		Latency:  ad.Latency,
	}

	streamParm := portaudio.StreamParameters{
		FramesPerBuffer: ad.FramesPerBuffer,
		Output:          streamDeviceParam,
		SampleRate:      ad.Samplingrate,
	}

	var stream *portaudio.Stream

	stream, err = portaudio.OpenStream(streamParm, &ad.out)
	if err != nil {
		fmt.Printf("unable to open playback audio stream on device %s\n", ad.DeviceName)
		fmt.Println(err)
		return // exit go routine
	}
	defer stream.Close()

	ad.Converter, err = gosamplerate.New(viper.GetInt("output_device.quality"), ad.Channels, 65536)
	if err != nil {
		fmt.Println("unable to create resampler")
		fmt.Println(err)
		return // exit go routine
	}
	defer gosamplerate.Delete(ad.Converter)

	if err = stream.Start(); err != nil {
		fmt.Printf("unable to start playback audio stream on device %s\n", ad.DeviceName)
		fmt.Println(err)
		return // exit go routine
	}

	defer stream.Stop()

	var d deserializer
	d.AudioDevice = &ad
	d.txTimestamp = time.Now()

	opusDecoder, err := opus.NewDecoder(int(ad.Samplingrate), ad.Channels)

	if err != nil || opusDecoder == nil {
		fmt.Println(err)
		return
	}
	d.opusDecoder = opusDecoder

	d.opusBuffer = make([]float32, 100000)

	r := ringBuffer.Ring{}
	r.SetCapacity(10)

	txUser := ""

	txUserResetTicker := time.NewTicker(1 * time.Second)
	txMonitorTicker := time.NewTicker(100 * time.Millisecond)

	for {
		select {

		// clear the tx user lock if nobody transmitted for the last second
		case <-txUserResetTicker.C:
			d.muTx.Lock()
			if time.Since(d.txTimestamp) > 500*time.Millisecond {
				d.txUser = ""
			}
			d.muTx.Unlock()

		case <-txMonitorTicker.C:
			d.muTx.Lock()

			if txUser != d.txUser {
				ad.Events.Pub(d.txUser, events.TxUser)
				txUser = d.txUser
			}
			d.muTx.Unlock()

		case msg := <-ad.ToDeserialize:
			r.Enqueue(msg.Data)

		default:
			data := r.Dequeue()
			if data != nil {
				err := d.DeserializeAudioMsg(data.([]byte))
				if err != nil {
					fmt.Println(err)
				} else {
					stream.Write()
				}
			} else {
				time.Sleep(time.Microsecond * 100)
			}
		}
	}
}
