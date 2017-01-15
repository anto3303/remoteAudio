package events

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cskr/pubsub"
)

const (
	RxAudioOn      = "audioOn"
	TxUser         = "txUser"
	MqttConnStatus = "mqttConnStatus"
)

type EventsConf struct {
	EventsPubSub *pubsub.PubSub
}

func CaptureKeyboard(conf EventsConf) {

	// rxAudioOn := false

	// rxAudioOnCh := conf.EventsPubSub.Sub(RxAudioOn)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// select {
		// case ev := <-rxAudioOnCh:
		// 	rxAudioOn = ev.(bool)
		// 	fmt.Println("Audio is", rxAudioOn)
		// default:
		// 	// pass
		// }

		if scanner.Scan() {
			switch scanner.Text() {
			case "o":
				conf.EventsPubSub.Pub(true, RxAudioOn)
			case "O":
				conf.EventsPubSub.Pub(false, RxAudioOn)
			default:
				fmt.Println("keyboard input:", scanner.Text())
			}
		}
	}
}
