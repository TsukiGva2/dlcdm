package lcdlogger

import (
	"fmt"

	"github.com/TsukiGva2/flick"
)

const (
	SCREEN_TEST = iota
	SCREEN_COUNT
)

func (display *SerialDisplay) ScreenTest(foo, bar int) {

	display.Forth.Send(
		fmt.Sprintf(
			"%d lbl %d num"+
				" %d lbl %d num"+
				" %d lbl %d num"+
				" %d lbl %d val",

			flick.PORTAL, 701,
			flick.REGIST, foo,
			flick.UNICAS, bar,
			flick.COMUNICANDO, flick.WEB,
		),
	)
}
