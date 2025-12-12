package main

import (
	"fmt"
	"gortex/internal/geom"
	"gortex/internal/screen"
	"gortex/internal/shapes"
	"math"
	"time"

	"golang.org/x/term"
)

func main() {
	w, h, err := term.GetSize(0)

	if err != nil {
		fmt.Println(err)
		return
	}

	tScreen := screen.InitTerminalScreen(w, h, ' ')
	c1 := shapes.Circle{Pos: geom.Vector2{X: -0.5, Y: -0.5}, R: 0.1}

	r := shapes.Rectangle{Pos: geom.Vector2{X: 0, Y: 0}, Size: geom.Vector2{X: 0.4, Y: 0.8}, Angle: 0}

	i := 0.0
	for {
		tScreen.BeginFrame()
		tScreen.Draw(c1, screen.GRADIENT)
		tScreen.Draw(r, []rune{'\u25AF', ' '})
		tScreen.Present()

		r.Angle++
		c1.R = max(min(math.Sin(i), 0.5), 0.001)

		i += 0.01
		time.Sleep(30 * time.Microsecond)
	}
}
