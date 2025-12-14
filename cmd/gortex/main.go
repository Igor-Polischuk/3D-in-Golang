package main

import (
	"fmt"
	"gortex/internal/geom"
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

	cube1 := scene.NewEntity(mesh.NewCube(geom.GetVector3(1, 1, 1)), geom.GetVector3(-1, 0, 0))
	cube2 := scene.NewEntity(mesh.NewCube(geom.GetVector3(1, 1, 1)), geom.GetVector3(1, 0, 0))

	cubes := []scene.Entity{cube1, cube2}

	for {
		screen.BeginFrame()

		// VIEW → камера в (0,0,0), дивиться у -Z
		eye := geom.Vector3{X: 0, Y: 0, Z: 5}
		target := geom.Vector3{X: 0, Y: 0, Z: -1}
		up := geom.Vector3{X: 0, Y: 1, Z: 0}

		view := geom.LookAt(eye, target, up)

		// PROJECTION
		aspect := float64(screen.Width()) / float64(screen.Height())
		fov := 60 * math.Pi / 180
		proj := geom.Perspective(fov, aspect, 0.1, 100)

		for _, cube := range cubes {
			model := cube.ModelMatrix()

			// FINAL MVP
			MVP := proj.Mul(&view).Mul(&model)

			render.RenderMesh(cube.Mesh, MVP, screen)
		}

		cubes[0].Rot.X += 0.001
		cubes[0].Rot.Z -= 0.007

		cubes[1].Rot.Y -= 0.007
		cubes[1].Rot.X += 0.006

		screen.Present()

	}
}
