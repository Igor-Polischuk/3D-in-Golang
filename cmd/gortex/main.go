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

	c := shapes.Circle{Pos: geom.Vector2{X: -1, Y: -1}, R: 0.5}
	r := shapes.Rectangle{Pos: geom.Vector2{X: 0, Y: 0}, Size: geom.Vector2{X: 0.5, Y: 0.5}, Angle: 45}

	for {
		tScreen.BeginFrame()
		tScreen.RasterShape(r)
		tScreen.RasterShape(c)
		tScreen.Present()
	}
}
