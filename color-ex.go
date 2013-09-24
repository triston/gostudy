package main

import (
	"fmt"
	"strings"
)

type Color int

const (
	Red   Color = 1 << iota // 1 << 0 = 1
	Green                   // 1 << 1 = 2
	Blue                    // 1 << 2 = 4
)

func (color Color) String() string {
	var colors []string
	if color&Red == Red {
		colors = append(colors, "red")
	}
	if color&Green == Green {
		colors = append(colors, "green")
	}
	if color&Blue == Blue {
		colors = append(colors, "blue")
	}
	if len(colors) > 0 {
		c := strings.Join(colors, "|")
		return fmt.Sprintf("[%s](%b)", c, int(color))
	}
	return "[](0)"
}

func main() {
	var color Color = 6
	fmt.Println(color)
}
