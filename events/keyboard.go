package events

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cskr/pubsub"
)

const (
	RxAudioOn = "audioOn"
	TxUser    = "txUser"
)

type EventsConf struct {
	EventsPubSub *pubsub.PubSub
}

type EventChs struct {
	RxAudioOn   chan interface{} // bool
	TxUserTopic chan interface{} // string
}

func CaptureKeyboard(conf EventsConf) {

	rxAudioOnCh := conf.EventsPubSub.Sub(RxAudioOn)

	rxAudioOn := false

	scanner := bufio.NewScanner(os.Stdin)
	for {

		select {
		case ev := <-rxAudioOnCh:
			rxAudioOn = ev.(bool)
		default:
			//pass
		}

		if scanner.Scan() {
			switch scanner.Text() {
			case "a":
				rxAudioOn = !rxAudioOn
				conf.EventsPubSub.Pub(rxAudioOn, RxAudioOn)
				fmt.Println("keyboard - Audio:", rxAudioOn)
			default:
				fmt.Println("keyboard input:", scanner.Text())
			}
		}
	}
}
