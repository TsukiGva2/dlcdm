package lcdlogger

import (
	"fmt"
)

const (
	SCREEN_TEST = iota
	SCREEN_COUNT
)

const (
	HELLO = iota
	WORLD
)

const (
	HAPPY_FACE = iota
	SAD_FACE
)

func (display *SerialDisplay) ScreenTest(face int) {

	display.Forth.Send(
		fmt.Sprintf(
			"%d lbl %d val"+
				" %d lbl %d val"+
				" %d lbl %d val"+
				" %d lbl %d val",

			HELLO, SAD_FACE,
			WORLD, HAPPY_FACE,
			HELLO, HAPPY_FACE,
			WORLD, face,
		),
	)
}
