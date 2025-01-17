package lcdlogger

import (
	"log"

	"github.com/TsukiGva2/flick"
)

type SerialDisplay struct {
	Forth  *flick.Forth
	Screen int

	switchButtonToggled bool
}

func NewSerialDisplay() (display SerialDisplay, err error) {

	f, err := flick.NewForth("/dev/ttyUSB0")

	if err != nil {

		log.Printf("Communication error: %v\n", err)

		return
	}

	f.Start()

	f.Query("WRD")

	display.Forth = &f

	f.Send("VAR bac")
	f.Send("VAR bst")
	f.Send(": btn 7 IN 0 = ;")
	f.Send(": chb bac @ NOT IF bst @ btn DUP ROT SWP NOT AND bac ! bst ! THN ;")
	f.Send("10 0 TMI chb 1 TME")

	return
}

func (display *SerialDisplay) SwitchScreens() {

	// TODO: onrelease actions

	res, err := display.Forth.Send("bac @ .")
	defer display.Forth.Send("0 bac !")

	if err != nil {

		return
	}

	if res[0] == '-' && !display.switchButtonToggled {

		display.Screen++
		display.Screen %= SCREEN_COUNT

		display.switchButtonToggled = true
	}

	if res[0] == '0' && display.switchButtonToggled {

		display.switchButtonToggled = false
	}
}
