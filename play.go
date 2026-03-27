package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

var wg int

func play(file string) {
	out, err := midi.FindOutPort(device)
	if err != nil {
		fmt.Println("cant find port")
		return
	}
	err = out.Open()
	defer out.Close()
	if err != nil {
		fmt.Println("cant open port")
		return
	}
	for _, event := range loadEvents(file) {
		go registerMsg(event)
	}
	for wg != 0 {
		out.Send(<-msgCh)
	}
}

func registerMsg(event Event) {
	wg++
	var msg midi.Message
	if event.velocity > 0 {
		msg = midi.NoteOn(0, event.key, event.velocity)
	} else {
		msg = midi.NoteOff(0, event.key)
	}
	time.Sleep(time.Millisecond * time.Duration(event.timedelta))
	msgCh <- msg
	wg--
}

func loadEvents(file string) []Event {
	var events []Event
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("failed to read file")
		os.Exit(1)
	}
	for eventString := range strings.SplitSeq(string(content), "\n") {
		if eventString == "" {
			continue
		}
		eventParams := strings.Split(eventString, " ")
		key, err := strconv.Atoi(eventParams[0])
		if err != nil {
			fmt.Println("failed to parse string to int")
			os.Exit(1)
		}
		velocity, err := strconv.Atoi(eventParams[1])
		if err != nil {
			fmt.Println("failed to parse string to int")
			os.Exit(1)
		}
		timedelta, err := strconv.Atoi(eventParams[2])
		if err != nil {
			fmt.Println("failed to parse string to float")
			os.Exit(1)
		}
		event := Event{key: uint8(key), velocity: uint8(velocity), timedelta: int32(timedelta)}
		events = append(events, event)
	}
	return events
}
