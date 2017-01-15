package events

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cskr/pubsub"
)

func CaptureKeyboard(evPS *pubsub.PubSub) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() {
			switch scanner.Text() {
			case "o":
				evPS.Pub(true, RxAudioOn)
			case "O":
				evPS.Pub(false, RxAudioOn)
			default:
				fmt.Println("keyboard input:", scanner.Text())
			}
		}
	}
}
