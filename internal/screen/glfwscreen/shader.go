package glfwscreen

import (
	_ "embed"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

//go:embed shaders/quad.vert
var quadVert string

//go:embed shaders/quad.frag
var quadFrag string

func compileShader(src string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)
	csrc, free := gl.Strs(src)
	defer free()

	gl.ShaderSource(shader, 1, csrc, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		panic("Shader compile error: " + log)
	}

	return shader
}

func (s *GLScreen) initShaders() {
	vs := compileShader(quadVert+"\x00", gl.VERTEX_SHADER)
	fs := compileShader(quadFrag+"\x00", gl.FRAGMENT_SHADER)

	s.shader = gl.CreateProgram()
	gl.AttachShader(s.shader, vs)
	gl.AttachShader(s.shader, fs)
	gl.LinkProgram(s.shader)

	gl.UseProgram(s.shader)
}
