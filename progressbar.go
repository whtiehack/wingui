// Package wingui https://docs.microsoft.com/zh-cn/windows/win32/controls/progress-bar-control#using-progress-bars
package wingui

import (
	"github.com/lxn/win"
	"github.com/whtiehack/wingui/winapi"
)

// ProgressBar a widget for Dialog. Progress Bar
type ProgressBar struct {
	WindowBase
}

// DeltaPos Advances the current position of a progress bar by a specified
// increment and redraws the bar to reflect the new position.
//
// Returns the previous position.
//
// If the increment results in a value outside the range of the control,
// the position is set to the nearest boundary.
//
// The behavior of this message is undefined if it is sent to a control that has the PBS_MARQUEE style.
func (pb *ProgressBar) DeltaPos(delta int) int {
	return int(pb.SendMessage(win.PBM_DELTAPOS, uintptr(delta), 0))
}

// GetBarColor Gets the color of the progress bar.
//
// Returns the color of the progress bar.
//
// This is the color set by the PBM_SETBARCOLOR message.
// The default value is CLR_DEFAULT, which is defined in commctrl.h.
//
// This function only affects the classic mode, not any visual style.
func (pb *ProgressBar) GetBarColor() int {
	return int(pb.SendMessage(winapi.PBM_GETBARCOLOR, 0, 0))
}

// NewProgressBar create a new ProgressBar,need bind to Dialog before use.
func NewProgressBar(idd uintptr) *ProgressBar {
	return &ProgressBar{
		WindowBase: WindowBase{idd: idd},
	}
}

// BindNewProgressBar create a new ProgressBar and bind to target dlg.
func BindNewProgressBar(idd uintptr, dlg *Dialog) (*ProgressBar, error) {
	pb := NewProgressBar(idd)
	err := dlg.BindWidgets(pb)
	return pb, err
}
