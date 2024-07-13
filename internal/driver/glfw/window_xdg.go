//go:build linux || freebsd || openbsd || netbsd
// +build linux freebsd openbsd netbsd

package glfw

import "github.com/gsmcwhirter/fyne/v2"

func (w *window) platformResize(canvasSize fyne.Size) {
	w.canvas.Resize(canvasSize)
}
