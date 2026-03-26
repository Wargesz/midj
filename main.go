package main

import (
	"fmt"
	"os"

	"gitlab.com/gomidi/midi/v2"
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
	if len(os.Args) != 3 {
		fmt.Println("noop")
		return
	}
	if os.Args[1] == "play" {
		play(os.Args[2])
	}
	if os.Args[1] == "record" {
		record(os.Args[2])
	}
	defer midi.CloseDriver()
}
