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
	c1 := shapes.Circle{Pos: geom.Vector2{X: -1, Y: -1}, R: 0.5}
	c2 := shapes.Circle{Pos: geom.Vector2{X: 1, Y: 1}, R: 0.5}
	c3 := shapes.Circle{Pos: geom.Vector2{X: -1, Y: 1}, R: 0.5}
	c4 := shapes.Circle{Pos: geom.Vector2{X: 1, Y: -1}, R: 0.5}

	tScreen.BeginFrame()
	tScreen.Draw(c1, screen.GRADIENT)
	tScreen.Draw(c2, screen.GRADIENT)
	tScreen.Draw(c3, screen.GRADIENT)
	tScreen.Draw(c4, []rune{'@'})
	tScreen.Present()

}
