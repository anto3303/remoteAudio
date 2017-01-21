package audio

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/dh1tw/gosamplerate"
	"github.com/dh1tw/remoteAudio/events"
	"github.com/gordonklaus/portaudio"
	"github.com/spf13/viper"
	ringBuffer "github.com/zfjagann/golang-ring"
	"gopkg.in/hraban/opus.v2"
)

//PlayerASync plays received audio on a local audio device asynchronously
func PlayerASync(ad AudioDevice) {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	defer ad.WaitGroup.Done()

	portaudio.Initialize()
	defer portaudio.Terminate()

	// Subscribe on events
	shutdownCh := ad.Events.Sub(events.Shutdown)

	// give Portaudio a few milliseconds to initialize
	// this is necessary to avoid a SIGSEGV in case
	// DefaultOutputDevice is accessed without portaudio
	// being completely initialized
	time.Sleep(time.Millisecond * 200)

	ad.out = make([]float32, ad.FramesPerBuffer*ad.Channels)

	fmt.Println("Player Channels:", ad.Channels)
	fmt.Println("Player Frames:", ad.FramesPerBuffer)
	fmt.Println("Player Out Buffer:", len(ad.out))

	//ad.out doesn't need to be initialized with a fixed buffer size
	//since the slice will be copied from the incoming data
	//and will therefore replay any buffer size

	var deviceInfo *portaudio.DeviceInfo
	var err error

	audioBufferSize := viper.GetInt("audio.rx_buffer_length")

	// initialize audio (ring) buffer
	r := ringBuffer.Ring{}
	r.SetCapacity(audioBufferSize)

	// select Playback Audio Device
	if ad.DeviceName == "default" {
		deviceInfo, err = portaudio.DefaultOutputDevice()
		if err != nil {
			fmt.Println("unable to find default playback sound device")
			fmt.Println(err)
			ad.WaitGroup.Done()
			return // exit go routine
		}
	} else {
		if err := ad.IdentifyDevice(); err != nil {
			fmt.Printf("unable to find recording sound device %s\n", ad.DeviceName)
			fmt.Println(err)
			ad.WaitGroup.Done()
			return
		}
	}

	// setup Audio Stream
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

	// the deserializer struct is mainly used to cache variables which have
	// to be read / set during the deserialization
	var d deserializer
	d.AudioDevice = &ad
	d.txTimestamp = time.Now()
	d.toPlay = make(chan []float32, 5)
	d.ring = ringBuffer.Ring{}
	d.ring.SetCapacity(10)
	d.muRing = sync.Mutex{}

	// initialize the Opus Decoder
	opusDecoder, err := opus.NewDecoder(int(ad.Samplingrate), ad.AudioStream.Channels)

	if err != nil || opusDecoder == nil {
		fmt.Println(err)
		ad.WaitGroup.Done()
		return
	}
	d.opusDecoder = opusDecoder
	d.opusBuffer = make([]float32, 520000) //max opus message size

	// open the audio stream
	stream, err = portaudio.OpenStream(streamParm, d.playCb)
	if err != nil {
		fmt.Printf("unable to open playback audio stream on device %s\n", ad.DeviceName)
		fmt.Println(err)
		ad.WaitGroup.Done()
		return // exit go routine
	}
	defer stream.Close()

	// create the PCM samplerate converter
	ad.PCMSamplerateConverter, err = gosamplerate.New(viper.GetInt("output_device.quality"), ad.Channels, 65536)
	if err != nil {
		fmt.Println("unable to create resampler")
		fmt.Println(err)
		ad.WaitGroup.Done()
		return // exit go routine
	}
	defer gosamplerate.Delete(ad.PCMSamplerateConverter)

	// Start the playback audio stream
	if err = stream.Start(); err != nil {
		fmt.Printf("unable to start playback audio stream on device %s\n", ad.DeviceName)
		fmt.Println(err)
		ad.WaitGroup.Done()
		return // exit go routine
	}
	defer stream.Stop()

	// cache holding the id of user from which we currently receive audio
	txUser := ""

	// Tickers to check if we still receive audio from a certain user.
	// This is needed on the server to release the "lock" and allow
	// others to transmit
	txUserResetTicker := time.NewTicker(1 * time.Second)
	txMonitorTicker := time.NewTicker(100 * time.Millisecond)

	// Everything has been set up, let's start execution

	for {
		select {

		// shutdown application gracefully
		case <-shutdownCh:
			log.Println("Shutdown Player")
			ad.WaitGroup.Done()
			return

		// clear the tx user lock if nobody transmitted during the last 500ms
		case <-txUserResetTicker.C:
			d.muTx.Lock()
			if time.Since(d.txTimestamp) > 500*time.Millisecond {
				d.txUser = ""
			}
			d.muTx.Unlock()

		// check if the tx user has changed
		case <-txMonitorTicker.C:
			d.muTx.Lock()

			if txUser != d.txUser {
				ad.Events.Pub(d.txUser, events.TxUser)
				txUser = d.txUser
			}
			d.muTx.Unlock()

		// write received audio data into the ring buffer
		case msg := <-ad.ToDeserialize:
			// check if new data is available in the ring buffer
			// fmt.Println("av to write", av)

			// err := d.DeserializeAudioMsg(data.([]byte))
			err := d.DeserializeAudioMsg(msg.Data)
			if err != nil {
				fmt.Println(err)
			}
			// d.toPlay <- ad.out
		}
	}
}

func (d *deserializer) playCb(in []float32, iTime portaudio.StreamCallbackTimeInfo, iFlags portaudio.StreamCallbackFlags) {
	switch iFlags {
	case portaudio.OutputUnderflow:
		fmt.Println("OutputUnderflow")
		return // move on!
	case portaudio.OutputOverflow:
		fmt.Println("OutputOverflow")
		return // move on!
	}

	d.muRing.Lock()
	data := d.ring.Dequeue()
	d.muRing.Unlock()

	if data != nil {
		audioData := data.([]float32)
		for i := 0; i < len(in); i++ {
			in[i] = audioData[i]
		}
	} else {
		// log.Println("write silence")
		for i := 0; i < len(in); i++ {
			in[i] = 0
		}
	}

	// log.Println("len Play Buffer:", len(d.toPlay))

	// select {
	// case data := <-d.toPlay:
	// 	if len(in) != len(data) {
	// 		log.Printf("unequal buffers! in: % bytes, sample: % bytes\n", len(in), len(data))
	// 	}
	// 	for i := 0; i < len(in); i++ {
	// 		in[i] = data[i]
	// 	}
	// default:
	// 	// write silence
	// 	log.Println("write silence")
	// 	for i := 0; i < len(in); i++ {
	// 		in[i] = 0
	// 	}
	// }
}
