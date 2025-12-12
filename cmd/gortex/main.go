package main

import (
	"fmt"
	"gortex/internal/camera"
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
	cam := camera.GetCamera(0, 0, 1)
	s := screen.InitTerminalScreen(w, h, &cam, ' ')

	c := shapes.Circle{Pos: geom.Vector2{X: -1, Y: -1}, R: 0.5}
	r := shapes.Rectangle{Pos: geom.Vector2{X: 0, Y: 0}, Size: geom.Vector2{X: 0.5, Y: 0.5}, Angle: 0}

	zoomIn := false

	for {
		s.BeginFrame()
		s.RasterShape(r)
		s.RasterShape(c)
		s.Present()

		if cam.Zoom > 7 {
			zoomIn = true
		} else if cam.Zoom < 0.05 {
			zoomIn = false
		}

		if zoomIn {
			cam.Zoom -= 0.05
		} else {
			cam.Zoom += 0.05
		}
	}
}
