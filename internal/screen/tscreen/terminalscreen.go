package tscreen

import (
	"fmt"
	"gortex/internal/camera"
	"gortex/internal/material"
	"strings"
)

type TermScreen struct {
	buffer              []rune
	aspect, pixelAspect float64
	H, W                int
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
		W:           w, H: h,
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
	// Move cursor to top-left (0,0) to overwrite previous frame
	fmt.Print("\033[H") // Move cursor to home position (top-left)

	// Build output string row by row for efficient printing
	var sb strings.Builder
	sb.Grow(s.W*s.H + s.H) // Pre-allocate space for all characters + newlines

	for y := 0; y < s.H; y++ {
		start := y * s.W
		end := start + s.W
		if end > len(s.buffer) {
			end = len(s.buffer)
		}
		sb.WriteString(string(s.buffer[start:end]))
		if y < s.H-1 {
			sb.WriteByte('\n') // Newline between rows
		}
	}

	// Print everything at once to minimize flickering
	fmt.Print(sb.String())
}

func (s TermScreen) SetPixel(x, y int, pixel material.Pixel) {
	symbol := pixel.Symbol
	if x >= 0 && x < s.W && y >= 0 && y < s.H {
		s.buffer[x+y*s.W] = symbol
	}
}

func (s TermScreen) Width() int {
	return s.W
}

func (s TermScreen) Height() int {
	return s.H
}
