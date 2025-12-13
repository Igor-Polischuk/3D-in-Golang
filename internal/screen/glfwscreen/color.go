package glfwscreen

type Color struct {
	R, G, B uint8
}

func GetColor(r, g, b uint8) Color {
	return Color{R: r, G: g, B: b}
}
