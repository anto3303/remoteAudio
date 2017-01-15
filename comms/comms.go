package comms

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/cskr/pubsub"
	"github.com/dh1tw/remoteAudio/audio"
	"github.com/dh1tw/remoteAudio/events"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttSettings struct {
	Transport                string
	BrokerURL                string
	BrokerPort               int
	ClientID                 string
	Topics                   []string
	ToDeserializeAudioDataCh chan audio.AudioMsg
	ToDeserializeAudioReqCh  chan audio.AudioMsg
	ToDeserializeAudioRespCh chan audio.AudioMsg
	ToWire                   chan audio.AudioMsg
	Events                   pubsub.PubSub
	LastWill                 *LastWill
}

type LastWill struct {
	Topic  string
	Data   []byte
	Qos    byte
	Retain bool
}

const (
	DISCONNECTED = 0
	CONNECTED    = 1
)

func MqttClient(s MqttSettings) {

	// mqtt.DEBUG = log.New(os.Stderr, "DEBUG - ", log.LstdFlags)
	// mqtt.CRITICAL = log.New(os.Stderr, "CRITICAL - ", log.LstdFlags)
	// mqtt.WARN = log.New(os.Stderr, "WARN - ", log.LstdFlags)
	// mqtt.ERROR = log.New(os.Stderr, "ERROR - ", log.LstdFlags)

	var msgHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

		if strings.Contains(msg.Topic(), "audio/audio") {
			audioMsg := audio.AudioMsg{
				Topic: msg.Topic(),
				Data:  msg.Payload()[:len(msg.Payload())],
			}
			s.ToDeserializeAudioDataCh <- audioMsg

		} else if strings.Contains(msg.Topic(), "request") {
			audioReqMsg := audio.AudioMsg{
				Data: msg.Payload()[:len(msg.Payload())],
			}
			s.ToDeserializeAudioReqCh <- audioReqMsg

		} else if strings.Contains(msg.Topic(), "response") {
			audioRespMsg := audio.AudioMsg{
				Data: msg.Payload()[:len(msg.Payload())],
			}
			s.ToDeserializeAudioRespCh <- audioRespMsg
		}
	}

	var connectionLostHandler = func(client mqtt.Client, err error) {
		log.Println("Connection lost to MQTT Broker; Reason:", err)
		s.Events.Pub(DISCONNECTED, events.MqttConnStatus)
	}

	// since we use SetCleanSession we have to subscribe on each
	// connect or reconnect to the channels
	var onConnectHandler = func(client mqtt.Client) {
		log.Println("Connected to MQTT Broker ")

		// Subscribe to Task Topics
		for _, topic := range s.Topics {
			if token := client.Subscribe(topic, 0, nil); token.Wait() &&
				token.Error() != nil {
				log.Println(token.Error())
			}
		}
		s.Events.Pub(CONNECTED, events.MqttConnStatus)
	}

	opts := mqtt.NewClientOptions().AddBroker(s.Transport + "://" + s.BrokerURL + ":" + strconv.Itoa(s.BrokerPort))
	opts.SetClientID(s.ClientID)
	opts.SetDefaultPublishHandler(msgHandler)
	opts.SetKeepAlive(time.Second * 5)
	opts.SetMaxReconnectInterval(time.Second)
	opts.SetCleanSession(true)
	opts.SetOnConnectHandler(onConnectHandler)
	opts.SetConnectionLostHandler(connectionLostHandler)
	opts.SetAutoReconnect(true)

	if s.LastWill != nil {
		opts.SetBinaryWill(s.LastWill.Topic, s.LastWill.Data, s.LastWill.Qos, s.LastWill.Retain)
	}

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	for {
		select {
		case msg := <-s.ToWire:
			token := client.Publish(msg.Topic, msg.Qos, msg.Retain, msg.Data)
			token.Wait()
		}
	}
}
