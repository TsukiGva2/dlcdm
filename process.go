package main

import (
	"time"

	"aa2/lcdlogger"
)

func main() {

	display, err := lcdlogger.NewSerialDisplay()

	if err != nil {

		return
	}

	for {

		switch display.Screen {
		case lcdlogger.SCREEN_TEST:
			display.ScreenTest(int(time.Now().Second() % 2))
		}

		display.SwitchScreens()

		time.Sleep(100 * time.Millisecond)
	}
}
