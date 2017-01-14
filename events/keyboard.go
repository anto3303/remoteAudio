package events

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cskr/pubsub"
)

const (
	RxAudioOn   = "RxAudioOn"
	TxUserTopic = "TxUserTopic"
	Ptt         = "Ptt"
)

type EventsConf struct {
	EventsPubSub *pubsub.PubSub
}

type EventChs struct {
	RxAudioOn   chan interface{} // bool
	TxUserTopic chan interface{} // string
}

func CaptureKeyboard(conf EventsConf) {

	ptt := false
	rxAudioOn := false

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			switch scanner.Text() {
			case "p":
				ptt = !ptt
				conf.EventsPubSub.Pub(ptt, Ptt)
				fmt.Println("keyboard - Ptt:", ptt)
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
