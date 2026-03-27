package main

import (
	"fmt"

	"gitlab.com/gomidi/midi/v2"
)

func kill() {
	out, err := midi.FindOutPort(device)
	if err != nil {
		fmt.Println("cant find device")
		return
	}
	err = out.Open()
	defer out.Close()
	if err != nil {
		fmt.Println("cant open port")
		return
	}
	out.Send(midi.ControlChange(0, 123, 0))
}
