package tscreen

import (
	"fmt"
	"gortex/internal/camera"
)

type TermScreen struct {
	buffer              []rune
	aspect, pixelAspect float64
	Height, Width       int
	bgRune              rune
	Cam                 *camera.Camera
}

const PIXEL_ASPECT = 11.0 / 24.0

func InitTerminalScreen(w, h int, cam *camera.Camera, bgRune rune) TermScreen {
	screen := TermScreen{
		buffer:      make([]rune, h*w),
		aspect:      float64(w) / float64(h),
		Cam:         cam,
		pixelAspect: PIXEL_ASPECT,
		Width:       w, Height: h,
		bgRune: bgRune,
	}

	for i := range screen.buffer {
		screen.buffer[i] = bgRune
	}

	fmt.Print("\033[2J")   // Clear screen
	fmt.Print("\033[?25l") // Hide cursor

	return screen
}

func (s *TermScreen) BeginFrame() {
	for i := range s.buffer {
		s.buffer[i] = s.bgRune
	}
}

func (s *TermScreen) Present() {
	str := string(s.buffer)
	fmt.Println(str)
}

func (s *TermScreen) SetPixel(x, y int, pixel rune) {
	if x >= 0 && x < s.Width && y >= 0 && y < s.Height {
		s.buffer[x+y*s.Width] = pixel
	}
}
