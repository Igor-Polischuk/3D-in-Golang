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

	eye := geom.Vector3{X: 0, Y: -2, Z: 4}
	target := geom.Vector3{X: 0, Y: 0, Z: -1}
	up := geom.Vector3{X: 0, Y: 1, Z: 0}

	view := geom.LookAt(eye, target, up)

	// PROJECTION
	aspect := float64(screen.Width()) / float64(screen.Height())
	fov := 60 * math.Pi / 180
	proj := geom.Perspective(fov, aspect, 0.1, 100)

	lightDirWorld := geom.Vector3{X: 1, Y: 1, Z: -1}
	lightDirView := view.TransformVector(lightDirWorld).Normalize()

	lambert := material.NewLambert(
		material.GetColor(252, 186, 3),
		lightDirView,
	)

	yellowBack := scene.NewEntity(
		mesh.NewCube(geom.GetVector3(1, 1, 1)),
		lambert,
		// material.NewASCIIFillMaterial('.'),
		geom.GetVector3(0, 0.5, 0))

	yellowBack.Rot.Y = 1

	terrainMesh := mesh.NewSquare(4)
	terrainMat := material.NewColorFillMaterial(material.GetColor(25, 230, 60))
	terrainPos := geom.GetVector3(0, 0, 0)

	terrain := scene.NewEntity(terrainMesh, terrainMat, terrainPos)
	terrain.Rot.X = 1.57

	// vp := proj.Mul(&view)

	// scn.Add(&yellowBack)
	scn.Add(&yellowBack)

	for {
		screen.BeginFrame()
		renderer.RenderScene(scn, view, proj)

		screen.Present()
		// whiteFront.Rot.Y -= 0.001
		yellowBack.Rot.Y += 0.008
		yellowBack.Rot.X = 0.008
		yellowBack.Rot.Z += 0.002

	}
}
