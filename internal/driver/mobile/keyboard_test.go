package mobile

import (
	"testing"

	_ "github.com/gsmcwhirter/fyne/v2/test"
)

func TestDevice_HideVirtualKeyboard_BeforeRun(t *testing.T) {
	hideVirtualKeyboard() // should not crash!
}
