# Gortex

A software-based 3D rendering engine written in Go. Gortex implements a complete 3D graphics pipeline including mesh transformation, perspective projection, triangle rasterization, depth buffering, and material shading.

## Features

- **3D Rendering Pipeline**: Complete implementation of the 3D graphics pipeline (model → view → projection → rasterization)
- **Mesh System**: Support for 3D meshes with vertices and indices (cubes, planes, and more)
- **Material System**: Multiple material types including:
  - Lambert shading with directional lighting
  - Color fill materials
  - ASCII fill materials
- **Scene Management**: Entity-based scene graph for organizing 3D objects
- **Transformations**: Full matrix-based transformation system (translation, rotation, scaling)
- **Depth Buffering**: Z-buffer for proper depth sorting and occlusion
- **Triangle Rasterization**: Software-based triangle filling with barycentric coordinates
- **Multiple Backends**: 
  - GLFW/OpenGL backend for windowed rendering
  - Terminal backend for ASCII art rendering

## Project Structure

```
go-3d/
├── cmd/
│   └── gortex/
│       └── main.go          # Application entry point
├── internal/
│   ├── camera/              # Camera and viewport management
│   ├── drawable/            # Drawable interface
│   ├── geom/                # Geometry and math utilities
│   │   ├── matrix.go        # Matrix operations
│   │   ├── vector2.go        # 2D vector operations
│   │   ├── vector3.go       # 3D vector operations
│   │   └── transformation.go # Transformation matrices
│   ├── material/            # Material and shading system
│   │   ├── material.go      # Material interface
│   │   ├── LambertMaterial.go # Lambert shading
│   │   ├── color.go         # Color utilities
│   │   └── textures.go      # Texture support
│   ├── mesh/                # Mesh definitions
│   │   ├── mesh.go          # Core mesh structure
│   │   ├── cube.go          # Cube mesh generator
│   │   └── plane.go         # Plane mesh generator
│   ├── render/              # Rendering engine
│   │   └── render.go        # Core rendering logic
│   ├── scene/               # Scene management
│   │   ├── scene.go         # Scene container
│   │   └── entity.go        # Scene entities
│   ├── screen/              # Screen backends
│   │   ├── glfwscreen/      # GLFW/OpenGL backend
│   │   └── tscreen/         # Terminal backend
│   ├── shapes/              # 2D shape utilities
│   └── utils/               # Utility functions
│       └── brezenhem.go     # Bresenham line algorithm
└── go.mod                   # Go module definition
```

## Requirements

- Go 1.25.3 or later
- OpenGL 4.1+ (for GLFW backend)
- GLFW 3.3+ libraries

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd go-3d
```

2. Install dependencies:
```bash
go mod download
```

3. Build the project:
```bash
go build ./cmd/gortex
```

4. Run the application:
```bash
./gortex
```

## Usage

The main application (`cmd/gortex/main.go`) demonstrates the rendering engine by creating a scene with a rotating cube and terrain. The renderer uses:

- **View Matrix**: Camera positioning using `LookAt`
- **Projection Matrix**: Perspective projection with configurable FOV
- **Model Matrix**: Per-entity transformations (translation, rotation, scaling)
- **Depth Buffer**: Z-buffering for correct depth sorting
- **Lambert Shading**: Directional lighting calculation

### Example Scene Setup

```go
// Create renderer and scene
screen := glfwscreen.InitGLFWScreen(1080, 720, nil)
renderer := render.NewRenderer(screen)
scn := scene.New()

// Setup camera
eye := geom.Vector3{X: 0, Y: -2, Z: 4}
target := geom.Vector3{X: 0, Y: 0, Z: -1}
up := geom.Vector3{X: 0, Y: 1, Z: 0}
view := geom.LookAt(eye, target, up)

// Setup projection
aspect := float64(screen.Width()) / float64(screen.Height())
fov := 60 * math.Pi / 180
proj := geom.Perspective(fov, aspect, 0.1, 100)

// Create entities
cube := scene.NewEntity(
    mesh.NewCube(geom.GetVector3(1, 1, 1)),
    material.NewLambert(color, lightDir),
    position,
)

scn.Add(&cube)

// Render loop
for {
    screen.BeginFrame()
    renderer.RenderScene(scn, view, proj)
    screen.Present()
    // Update entity rotations, etc.
}
```

## Architecture

### Rendering Pipeline

1. **Model Space**: Vertices defined in local coordinate space
2. **World Space**: Transform vertices using model matrix
3. **View Space**: Transform to camera coordinate space
4. **Clip Space**: Apply perspective projection
5. **NDC (Normalized Device Coordinates)**: Perspective divide
6. **Screen Space**: Map NDC to pixel coordinates
7. **Rasterization**: Fill triangles using barycentric coordinates
8. **Depth Test**: Z-buffer for occlusion culling

### Key Components

- **Renderer**: Handles the rendering loop, depth buffering, and triangle rasterization
- **Mesh**: Defines 3D geometry with vertices and indices
- **Material**: Shading interface for calculating pixel colors
- **Scene**: Container for entities with transformation matrices
- **Screen**: Abstract interface for pixel output (GLFW or terminal)

## Dependencies

- `golang.org/x/term` - Terminal size detection
- `github.com/go-gl/gl` - OpenGL bindings
- `github.com/go-gl/glfw/v3.3/glfw` - Window and context management


