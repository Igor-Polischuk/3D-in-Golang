package main

import (
	"fmt"
	"gortex/internal/geom"
	"gortex/internal/screen"
	"gortex/internal/shapes"

	"golang.org/x/term"
)

func main() {
	w, h, err := term.GetSize(0)

	if err != nil {
		fmt.Println(err)
		return
	}

	tScreen := screen.InitTerminalScreen(w, h, ' ')
	c := shapes.Circle{Pos: geom.Vector2{X: 0, Y: 0}, R: 0.25}

	tScreen.BeginFrame()
	tScreen.Draw(c)
	tScreen.Present()
}
