package main

import (
	"fmt"
	"os"

	"gitlab.com/gomidi/midi/v2"
)

const device string = "UMC204HD"
var (
	msgCh chan midi.Message
	eventCh chan Event
)

type Event struct {
	key       uint8
	velocity  uint8
	timedelta int32
}

func (e *Event) String() string {
	return fmt.Sprintf("%v %v %v", e.key, e.velocity, e.timedelta)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("noop")
		return
	}
	msgCh = make(chan midi.Message, 5)
	eventCh = make(chan Event, 5)
	defer close(msgCh)
	defer close(eventCh)
	if os.Args[1] == "kill" {
		kill()
	}
	if os.Args[1] == "play" {
		play(os.Args[2])
	}
	if os.Args[1] == "record" {
		record(os.Args[2])
	}
	defer midi.CloseDriver()
}
