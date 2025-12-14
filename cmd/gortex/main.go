package main

import (
	"fmt"
	"gortex/internal/geom"
	"gortex/internal/material"
	"gortex/internal/mesh"
	"gortex/internal/render"
	"gortex/internal/scene"
	"gortex/internal/screen/glfwscreen"
	"math"

	"golang.org/x/term"
)

func main() {
	w, h, err := term.GetSize(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	geom.Rotate(float64(w + h))

	// screen := tscreen.InitTerminalScreen(w, h, nil, ' ')
	screen := glfwscreen.InitGLFWScreen(1080, 720, nil)
	renderer := render.NewRenderer(screen)
	scn := scene.New()

	whiteFront := scene.NewEntity(
		mesh.NewCube(geom.GetVector3(1, 1, 1)),
		material.NewColorFillMaterial(material.GetColor(255, 255, 255)),
		geom.GetVector3(0.4, 0, 1))

	yellowBack := scene.NewEntity(
		mesh.NewCube(geom.GetVector3(1, 1, 1)),
		material.NewColorFillMaterial(material.GetColor(252, 186, 3)),
		geom.GetVector3(0, 0, 0))

	eye := geom.Vector3{X: 0, Y: 0, Z: 2}
	target := geom.Vector3{X: 0, Y: 0, Z: -1}
	up := geom.Vector3{X: 0, Y: 1, Z: 0}

	view := geom.LookAt(eye, target, up)

	// PROJECTION
	aspect := float64(screen.Width()) / float64(screen.Height())
	fov := 60 * math.Pi / 180
	proj := geom.Perspective(fov, aspect, 0.1, 100)

	vp := proj.Mul(&view)

	scn.Add(&yellowBack, &whiteFront)

	for {
		screen.BeginFrame()
		renderer.Render(scn, vp)

		screen.Present()
		// cube1.Rot.X += 0.001

	}
}
