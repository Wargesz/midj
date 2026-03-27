package main

import (
	"fmt"
	"os"
	"os/signal"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

var events []Event
var lastTimestamp int32 = 0

func record(file string) {
	in, err := midi.FindInPort(device)
	if err != nil {
		fmt.Println("cant find port")
		return
	}
	err = in.Open()
	defer in.Close()
	if err != nil {
		fmt.Println("cant open port")
		return
	}
	stop, err := midi.ListenTo(in, callback)
	if err != nil {
		fmt.Println("cant start listener")
		return
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	stop()
	saveEvents(file, events)
}

func callback(msg midi.Message, timestampms int32) {
	delta := timestampms - lastTimestamp
	lastTimestamp = timestampms
	var channel, key, velocity uint8
	if msg.GetNoteOn(&channel, &key, &velocity) && channel == 0 {
		events = append(events, Event{key: key, velocity: velocity, timedelta: delta})
	}
}

func saveEvents(file string, events []Event) {
	var content string
	for _, event := range events {
		content = content + event.String() + "\n"
	}
	err := os.WriteFile(file, []byte(content), 0664)
	if err != nil {
		fmt.Println("failed to save events")
	}
}
