package screen

import (
	"fmt"
	"gortex/internal/camera"
	"os"
	"os/exec"
	"runtime"
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

	return screen
}

func (s *TermScreen) BeginFrame() {
	for i := range s.buffer {
		s.buffer[i] = s.bgRune
	}
	ClearTerminal()
}

func (s *TermScreen) Draw() {

}

func (s *TermScreen) Present() {
	str := string(s.buffer)
	fmt.Println(str)
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
