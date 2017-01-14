// Copyright © 2016 Tobias Wellnitz, DH1TW <Tobias.Wellnitz@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cskr/pubsub"
	"github.com/dh1tw/remoteAudio/audio"
	"github.com/dh1tw/remoteAudio/comms"
	"github.com/dh1tw/remoteAudio/events"
	"github.com/gogo/protobuf/proto"
	"github.com/gordonklaus/portaudio"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "net/http/pprof"

	sbAudio "github.com/dh1tw/remoteAudio/sb_audio"
)

// serveMqttCmd represents the mqtt command
var serveMqttCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "Server streaming Audio via MQTT",
	Long:  `Server streaming Audio via MQTT`,
	Run: func(cmd *cobra.Command, args []string) {
		mqttAudioServer()
	},
}

func init() {
	serveCmd.AddCommand(serveMqttCmd)
	serveMqttCmd.PersistentFlags().StringP("broker_url", "u", "localhost", "Broker URL")
	serveMqttCmd.PersistentFlags().StringP("client_id", "c", "", "MQTT Client Id")
	serveMqttCmd.PersistentFlags().IntP("broker_port", "p", 1883, "Broker Port")
	serveMqttCmd.PersistentFlags().StringP("station", "X", "mystation", "Your station callsign")
	serveMqttCmd.PersistentFlags().StringP("radio", "Y", "myradio", "Radio ID")
	viper.BindPFlag("mqtt.broker_url", serveMqttCmd.PersistentFlags().Lookup("broker_url"))
	viper.BindPFlag("mqtt.broker_port", serveMqttCmd.PersistentFlags().Lookup("broker_port"))
	viper.BindPFlag("mqtt.client_id", serveMqttCmd.PersistentFlags().Lookup("client_id"))
	viper.BindPFlag("mqtt.station", serveMqttCmd.PersistentFlags().Lookup("station"))
	viper.BindPFlag("mqtt.radio", serveMqttCmd.PersistentFlags().Lookup("radio"))
}

func mqttAudioServer() {

	// defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// viper settings need to be copied in local variables
	// since viper lookups allocate of each lookup a copy
	// and are quite inperformant

	mqttBrokerURL := viper.GetString("mqtt.broker_url")
	mqttBrokerPort := viper.GetInt("mqtt.broker_port")
	mqttClientID := viper.GetString("mqtt.client_id")

	baseTopic := viper.GetString("mqtt.station") +
		"/radios/" + viper.GetString("mqtt.radio") +
		"audio"

	serverRequestTopic := baseTopic + "/request"
	serverAudioOutTopic := baseTopic + "/audio_out"
	serverAudioInTopic := baseTopic + "/audio_in"
	// responseTopic := baseTopic + "/response"
	// errorTopic := baseTopic + "/error"
	// audioTopic := baseTopic + "/audio_data"

	mqttTopics := []string{serverRequestTopic, serverAudioInTopic}

	audioFrameLength := viper.GetInt("audio.frame_length")
	rxBufferLength := viper.GetInt("audio.rx_buffer_length")

	outputDeviceDeviceName := viper.GetString("output_device.device_name")
	outputDeviceSamplingrate := viper.GetFloat64("output_device.samplingrate")
	outputDeviceLatency := viper.GetDuration("output_device.latency")
	outputDeviceChannels := viper.GetString("output_device.channels")

	inputDeviceDeviceName := viper.GetString("input_device.device_name")
	inputDeviceSamplingrate := viper.GetFloat64("input_device.samplingrate")
	inputDeviceLatency := viper.GetDuration("input_device.latency")
	inputDeviceChannels := viper.GetString("input_device.channels")

	portaudio.Initialize()

	connStatus := pubsub.New(1)

	toWireCh := make(chan audio.AudioMsg, 20)
	toSerializeAudioDataCh := make(chan audio.AudioMsg, 20)
	toDeserializeAudioDataCh := make(chan audio.AudioMsg, rxBufferLength)
	toDeserializeAudioReqCh := make(chan audio.AudioMsg, 10)

	evPS := pubsub.New(1)

	settings := comms.MqttSettings{
		Transport:  "tcp",
		BrokerURL:  mqttBrokerURL,
		BrokerPort: mqttBrokerPort,
		ClientID:   mqttClientID,
		Topics:     mqttTopics,
		ToDeserializeAudioDataCh: toDeserializeAudioDataCh,
		ToDeserializeAudioReqCh:  toDeserializeAudioReqCh,
		ToWire:                   toWireCh,
		TxUserTopic:              evPS.Sub(events.TxUserTopic),
		ConnStatus:               *connStatus,
		InputBufferLength:        rxBufferLength,
	}

	player := audio.AudioDevice{
		ToWire:        nil,
		ToSerialize:   nil,
		ToDeserialize: toDeserializeAudioDataCh,
		EventChs: events.EventChs{
			RxAudioOn: nil,
		},
		AudioStream: audio.AudioStream{
			DeviceName:      outputDeviceDeviceName,
			FramesPerBuffer: audioFrameLength,
			Samplingrate:    outputDeviceSamplingrate,
			Latency:         outputDeviceLatency,
			Channels:        audio.GetChannel(outputDeviceChannels),
		},
	}

	recorder := audio.AudioDevice{
		ToWire:           toWireCh,
		ToSerialize:      toSerializeAudioDataCh,
		ToDeserialize:    nil,
		AudioToWireTopic: serverAudioOutTopic,
		EventChs: events.EventChs{
			RxAudioOn: evPS.Sub(events.RxAudioOn),
		},
		AudioStream: audio.AudioStream{
			DeviceName:      inputDeviceDeviceName,
			FramesPerBuffer: audioFrameLength,
			Samplingrate:    inputDeviceSamplingrate,
			Latency:         inputDeviceLatency,
			Channels:        audio.GetChannel(inputDeviceChannels),
		},
	}

	go audio.PlayerSync(player)
	go audio.RecorderAsync(recorder)

	go comms.MqttClient(settings)

	eventsConf := events.EventsConf{
		EventsPubSub: evPS,
	}

	go events.CaptureKeyboard(eventsConf)

	connectionStatusCh := connStatus.Sub(comms.CONNSTATUSTOPIC)

	for {
		select {
		case status := <-connectionStatusCh:
			fmt.Println(status)
		case data := <-toDeserializeAudioReqCh:

			msg := sbAudio.ClientRequest{}

			err := proto.Unmarshal(data.Data, &msg)
			if err != nil {
				fmt.Println(err)
			}

			if msg.RxAudioOn != nil {
				rxAudioOn := msg.GetRxAudioOn()
				evPS.Pub(rxAudioOn, events.RxAudioOn)
			}
		}
	}
}
