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

// connectMqttCmd represents the mqtt command
var connectMqttCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "Client streaming Audio via MQTT",
	Long:  `Client streaming Audio via MQTT`,
	Run: func(cmd *cobra.Command, args []string) {
		mqttAudioClient()
	},
}

func init() {
	connectCmd.AddCommand(connectMqttCmd)
	connectMqttCmd.PersistentFlags().StringP("broker_url", "u", "localhost", "Broker URL")
	connectMqttCmd.PersistentFlags().StringP("client_id", "c", "", "MQTT Client Id")
	connectMqttCmd.PersistentFlags().IntP("broker_port", "p", 1883, "Broker Port")
	connectMqttCmd.PersistentFlags().StringP("station", "X", "mystation", "Your station callsign")
	connectMqttCmd.PersistentFlags().StringP("radio", "Y", "myradio", "Radio ID")
	viper.BindPFlag("mqtt.broker_url", connectMqttCmd.PersistentFlags().Lookup("broker_url"))
	viper.BindPFlag("mqtt.broker_port", connectMqttCmd.PersistentFlags().Lookup("broker_port"))
	viper.BindPFlag("mqtt.client_id", connectMqttCmd.PersistentFlags().Lookup("client_id"))
	viper.BindPFlag("mqtt.station", connectMqttCmd.PersistentFlags().Lookup("station"))
	viper.BindPFlag("mqtt.radio", connectMqttCmd.PersistentFlags().Lookup("radio"))
}

func mqttAudioClient() {

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

	serverBaseTopic := viper.GetString("mqtt.station") +
		"/radios/" + viper.GetString("mqtt.radio") +
		"/audio"

	serverRequestTopic := serverBaseTopic + "/request"
	serverResponseTopic := serverBaseTopic + "/response"

	// errorTopic := baseTopic + "/error"
	serverAudioOutTopic := serverBaseTopic + "/audio_out"
	serverAudioInTopic := serverBaseTopic + "/audio_in"

	mqttTopics := []string{serverResponseTopic, serverAudioOutTopic}

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
	toDeserializeAudioRespCh := make(chan audio.AudioMsg, 10)

	evPS := pubsub.New(1)

	settings := comms.MqttSettings{
		Transport:  "tcp",
		BrokerURL:  mqttBrokerURL,
		BrokerPort: mqttBrokerPort,
		ClientID:   mqttClientID,
		Topics:     mqttTopics,
		ToDeserializeAudioDataCh: toDeserializeAudioDataCh,
		ToDeserializeAudioReqCh:  nil,
		ToDeserializeAudioRespCh: toDeserializeAudioRespCh,
		ToWire:            toWireCh,
		TxUserTopic:       evPS.Sub(events.TxUserTopic),
		ConnStatus:        *connStatus,
		InputBufferLength: rxBufferLength,
	}

	player := audio.AudioDevice{
		ToWire:           nil,
		ToSerialize:      nil,
		ToDeserialize:    toDeserializeAudioDataCh,
		AudioToWireTopic: serverAudioOutTopic,
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
		AudioToWireTopic: serverAudioInTopic,
		ToDeserialize:    nil,
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
		case msg := <-connectionStatusCh:
			s := msg.(comms.ConnectionStatus)

			if s.Status == comms.CONNECTED {
				req := sbAudio.ClientRequest{}
				on := true
				req.RxAudioOn = &on
				m, err := req.Marshal()
				if err != nil {
					fmt.Println(err)
				} else {
					wireMsg := audio.AudioMsg{
						Topic: serverRequestTopic,
						Data:  m,
					}
					toWireCh <- wireMsg
				}
			}

		// responses coming from server
		case data := <-toDeserializeAudioRespCh:

			msg := sbAudio.ServerResponse{}

			err := proto.Unmarshal(data.Data, &msg)
			if err != nil {
				fmt.Println(err)
			}

			if msg.RxAudioOn != nil {
				rxAudioOn := msg.GetRxAudioOn()
				fmt.Printf("Server Audio is %t", rxAudioOn)
			}

			if msg.TxUser != nil {
				txUser := msg.GetTxUser()
				fmt.Printf("Server: Current TX User: %s", txUser)
			}

			if msg.Tx != nil {
				tx := msg.GetTx()
				fmt.Printf("Server: TX %t", tx)
			}
		}
	}
}
