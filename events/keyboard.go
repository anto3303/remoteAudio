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

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			if scanner.Text() == "p" {
				ptt = !ptt
				conf.EventsPubSub.Pub(true, RxAudioOn)
				fmt.Println("keyboard - ptt:", ptt)
			} else {
				fmt.Println("keyboard input:", scanner.Text())
			}
		}
	}
}
