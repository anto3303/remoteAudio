package comms

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/cskr/pubsub"
	"github.com/dh1tw/remoteAudio/audio"
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
	ToWire                   chan audio.AudioMsg
	TxUserTopic              chan interface{}
	ConnStatus               pubsub.PubSub
	InputBufferLength        int
}

const (
	CONNECTED    = 1
	DISCONNECTED = 2
)

const (
	CONNSTATUSTOPIC = "ConnectionStatusTopic"
)

type ConnectionStatus struct {
	Status int
}

func MqttClient(s MqttSettings) {

	// mqtt.DEBUG = log.New(os.Stderr, "DEBUG - ", log.LstdFlags)
	// mqtt.CRITICAL = log.New(os.Stderr, "CRITICAL - ", log.LstdFlags)
	// mqtt.WARN = log.New(os.Stderr, "WARN - ", log.LstdFlags)
	// mqtt.ERROR = log.New(os.Stderr, "ERROR - ", log.LstdFlags)

	var msgHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

		if strings.Contains(msg.Topic(), "audio_data") {
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
		}
	}

	var connectionLostHandler = func(client mqtt.Client, err error) {
		log.Println("Connection lost to MQTT Broker; Reason:", err)
		status := ConnectionStatus{DISCONNECTED}
		s.ConnStatus.Pub(status, CONNSTATUSTOPIC)
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
		status := ConnectionStatus{CONNECTED}
		s.ConnStatus.Pub(status, CONNSTATUSTOPIC)
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

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	txUserTopic := ""

	for {
		select {
		case msg := <-s.ToWire:
			token := client.Publish(msg.Topic, 0, false, msg.Data)
			token.Wait()
		case msg := <-s.TxUserTopic:
			if client.IsConnected() {
				if len(txUserTopic) != 0 {
					if token := client.Unsubscribe(txUserTopic); token.Wait() && token.Error() != nil {
						log.Println(token.Error())
					}
					txUserTopic = msg.(string)
					token := client.Subscribe(txUserTopic, 0, nil)
					if token.Wait() && token.Error() != nil {
						log.Println(token.Error())
					}
				}
			}
		}
	}
}
