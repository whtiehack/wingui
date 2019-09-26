package wingui

import (
	"github.com/lxn/win"
	"github.com/whtiehack/wingui/winapi"
)

// https://docs.microsoft.com/zh-cn/windows/win32/controls/trackbar-controls

//	TrackBar a widget for Dialog. Progress Bar
type TrackBar struct {
	WindowBase
}

// ClearSel clears the current selection range in a trackbar.
// redraw if this parameter is TRUE, the trackbar is redrawn after the selection is cleared.
func (tb *TrackBar) ClearSel(redraw bool) {
	rd := 0
	if redraw {
		rd = 1
	}
	tb.SendMessage(winapi.TBM_CLEARSEL, uintptr(rd), 0)
}

// ClearTics removes the current tick marks from a trackbar.
// This message does not remove the first and last tick marks,
// which are created automatically by the trackbar.
// redraw if this parameter is TRUE, the trackbar is redrawn after the selection is cleared.
func (tb *TrackBar) ClearTics(redraw bool) {
	rd := 0
	if redraw {
		rd = 1
	}
	tb.SendMessage(winapi.TBM_CLEARTICS, uintptr(rd), 0)
}

//	GetBuddy retrieves the handle to a trackbar control buddy window at a given location.
//	The specified location is relative to the control's orientation (horizontal or vertical).
//	isLeftOrTop indicating which buddy window handle will be retrieved, by relative location.
//	This value can be one of the following:
//	- TRUE
//	- Retrieves the handle to the buddy to the left of the trackbar. If the trackbar control uses the TBS_VERT style, the message will retrieve the buddy above the trackbar.
//	- FALSE
//	- Retrieves the handle to the buddy to the right of the trackbar. If the trackbar control uses the TBS_VERT style, the message will retrieve the buddy below the trackbar.
func (tb *TrackBar) GetBuddy(isLeftOrTop bool) win.HWND {
	rd := 0
	if isLeftOrTop {
		rd = 1
	}
	return win.HWND(tb.SendMessage(winapi.TBM_GETBUDDY, uintptr(rd), 0))
}

// NewTrackBar create a new TrackBar,need bind to Dialog before use.
func NewTrackBar(idd uintptr) *TrackBar {
	return &TrackBar{
		WindowBase: WindowBase{idd: idd},
	}
}

// BindNewTrackBar create a new TrackBar and bind to target dlg.
func BindNewTrackBar(idd uintptr, dlg *Dialog) (*TrackBar, error) {
	bb := NewTrackBar(idd)
	err := dlg.BindWidgets(bb)
	return bb, err
}
