package glfwscreen

import (
	"gortex/internal/camera"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type GLScreen struct {
	buffer               []uint8
	W, H                 int
	Cam                  *camera.Camera
	window               *glfw.Window
	texture, vao, shader uint32
	aspect               float32
}

func InitGLFWScreen(w, h int, cam *camera.Camera) *GLScreen {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	// macOS requires these
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(w, h, "GL Screen", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	gl.Init()

	fw, fh := window.GetFramebufferSize()

	s := &GLScreen{
		W:      fw,
		H:      fh,
		Cam:    cam,
		buffer: make([]uint8, fw*fh*4), // RGBA buffer
		window: window,
		aspect: float32(w) / float32(h),
	}

	s.initTexture()
	s.initQuad()
	s.initShaders()

	return s
}

func (s *GLScreen) initTexture() {
	gl.GenTextures(1, &s.texture)
	gl.BindTexture(gl.TEXTURE_2D, s.texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
		int32(s.W), int32(s.H), 0,
		gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(s.buffer))
}

func (s *GLScreen) initQuad() {
	quadVertices := []float32{
		-1, -1, 0, 0,
		1, -1, 1, 0,
		1, 1, 1, 1,
		-1, -1, 0, 0,
		1, 1, 1, 1,
		-1, 1, 0, 1,
	}

	var vbo uint32
	gl.GenVertexArrays(1, &s.vao)
	gl.GenBuffers(1, &vbo)

	gl.BindVertexArray(s.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(quadVertices)*4, gl.Ptr(quadVertices), gl.STATIC_DRAW)

	// pos
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 4*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// uv
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 4*4, gl.PtrOffset(8))
	gl.EnableVertexAttribArray(1)
}

func (s *GLScreen) ShouldClose() bool {
	return s.window.ShouldClose()
}

func (s *GLScreen) BeginFrame() {
	for i := range s.buffer {
		s.buffer[i] = 0 // clear to black
	}
}

func (s *GLScreen) SetPixel(x, y int, color Color) {
	if x < 0 || y < 0 || x >= s.W || y >= s.H {
		return
	}
	i := (y*s.W + x) * 4
	s.buffer[i+0] = color.R
	s.buffer[i+1] = color.G
	s.buffer[i+2] = color.B
	s.buffer[i+3] = 255
}

func (s *GLScreen) Present() {
	if s.window.ShouldClose() {
		return
	}

	gl.Viewport(0, 0, int32(s.W), int32(s.H))
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.BindTexture(gl.TEXTURE_2D, s.texture)

	gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, int32(s.W), int32(s.H),
		gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(s.buffer))

	gl.BindVertexArray(s.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)

	s.window.SwapBuffers()
	glfw.PollEvents()
}
