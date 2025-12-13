package main

import (
	"fmt"
	"gortex/internal/geom"
	"gortex/internal/mesh"
	"gortex/internal/render"
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

	trz := 0.0
	crz := 0.0
	srz := 0.0

	// screen := tscreen.InitTerminalScreen(w, h, nil, ' ')
	screen := glfwscreen.InitGLFWScreen(1080, 720, nil)

	for {
		screen.BeginFrame()

		triangle := mesh.NewTriangle()
		square := mesh.NewSquare()
		cube := mesh.NewCube()

		tModel := geom.Translate3D(0, 0, -3)
		tRotated := geom.RotateY(trz)
		tModel = tModel.Mul(&tRotated)

		cModel := geom.Translate3D(-2, 0, -3)
		cRotated := geom.RotateY(crz)
		cModel = cModel.Mul(&cRotated)

		sModel := geom.Translate3D(2, 0, -3)
		sRotated := geom.RotateY(srz)
		sModel = sModel.Mul(&sRotated)

		// VIEW → камера в (0,0,0), дивиться у -Z
		eye := geom.Vector3{X: 0, Y: 0, Z: 2}
		target := geom.Vector3{X: 0, Y: 0, Z: -1}
		up := geom.Vector3{X: 0, Y: 1, Z: 0}

		view := geom.LookAt(eye, target, up)

		// PROJECTION
		aspect := float64(screen.Width()) / float64(screen.Height())
		fov := 60 * math.Pi / 180
		proj := geom.Perspective(fov, aspect, 0.1, 100)

		// FINAL MVP
		tMVP := proj.Mul(&view).Mul(&tModel)
		cMVP := proj.Mul(&view).Mul(&cModel)
		sMVP := proj.Mul(&view).Mul(&sModel)

		render.RenderMesh(triangle, tMVP, screen)
		render.RenderMesh(square, sMVP, screen)
		render.RenderMesh(cube, cMVP, screen)

		screen.Present()

		trz += 0.03
		srz -= 0.01
		crz -= 0.01
	}
}
