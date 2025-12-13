package main

import (
	"fmt"
	"gortex/internal/camera"
	"gortex/internal/geom"

	"gortex/internal/screen/glfwscreen"
	"gortex/internal/screen/tscreen"
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
	gl := glfwscreen.InitGLFWScreen(1080, 720, &cam)
	ts := tscreen.InitTerminalScreen(w, h, &cam, ' ')

	c := shapes.Circle{Pos: geom.Vector2{X: 1, Y: 1}, R: 0.5}
	r := shapes.Rectangle{Pos: geom.Vector2{X: 0, Y: 0}, Size: geom.Vector2{X: 1, Y: 0.1}, Angle: 45}
	r2 := shapes.Rectangle{Pos: geom.Vector2{X: -1, Y: 1}, Size: geom.Vector2{X: 1, Y: 0.1}, Angle: 0}

	for {
		gl.BeginFrame()
		ts.BeginFrame()

		gl.RasterShape(c)
		ts.RasterShape(c)
		gl.RasterShape(r)
		gl.RasterShape(r2)
		ts.RasterShape(r)
		ts.RasterShape(r2)

		ts.DrawLine(0, 0, ts.Width, ts.Height, '~')
		gl.DrawLine(0, 0, gl.W, gl.H, glfwscreen.GetColor(100, 100, 100))

		gl.Present()
		ts.Present()
	}
}
