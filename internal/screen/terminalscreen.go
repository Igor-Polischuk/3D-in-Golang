package screen

import (
	"fmt"
	"gortex/internal/drawable"
	"os"
	"os/exec"
	"runtime"
)

type TermScreen struct {
	buffer              []rune
	currentTexture      []rune
	aspect, pixelAspect float64
	height, width       int
	bgRune              rune
}

const PIXEL_ASPECT = 11.0 / 24.0

func InitTerminalScreen(w, h int, bgRune rune) TermScreen {
	screen := TermScreen{
		buffer:      make([]rune, h*w),
		aspect:      float64(w) / float64(h),
		pixelAspect: PIXEL_ASPECT,
		width:       w, height: h,
		bgRune:         bgRune,
		currentTexture: []rune{'@'},
	}

	for i := range screen.buffer {
		screen.buffer[i] = bgRune
	}

	return screen
}

func (s *TermScreen) BeginFrame() {
	for i := range s.buffer {
		s.buffer[i] = s.bgRune
	}
	ClearTerminal()
}

func (s *TermScreen) Draw(drawable drawable.Drawable) {
	drawable.Rasterize(s)
}

func (s *TermScreen) Present() {
	str := string(s.buffer)
	fmt.Println(str)
}

func (s *TermScreen) UseTexture(t []rune) {
	s.currentTexture = t
}

func ClearTerminal() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
