package wingui

import (
	"github.com/lxn/win"
	"github.com/whtiehack/wingui/winapi"
	"unsafe"
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

//	Retrieves the size and position of the bounding rectangle for a trackbar's channel.
//	(The channel is the area over which the slider moves. It contains the highlight when a range is selected.)
//	Return value is a RECT structure. This structure with the channel's bounding rectangle,
//	in client coordinates of the trackbar's window.
func (tb *TrackBar) GetChannelRect() win.RECT {
	var rect win.RECT
	tb.SendMessage(winapi.TBM_GETCHANNELRECT, 0, uintptr(unsafe.Pointer(&rect)))
	return rect
}

//	GetLineSize retrieves the number of logical positions the trackbar's slider moves
//	in response to keyboard input from the arrow keys, such as the or keys.
//	The logical positions are the integer increments in the trackbar's range of minimum to maximum slider positions.
//	Returns a 32-bit value that specifies the line size for the trackbar.
//	The default setting for the line size is 1.
//	The trackbar also sends a WM_HSCROLL or WM_VSCROLL message with the TB_LINEUP and TB_LINEDOWN notification codes
//	to its parent window when the user presses the arrow keys.
func (tb *TrackBar) GetLineSize() int {
	return int(tb.SendMessage(winapi.TBM_GETLINESIZE, 0, 0))
}

//	GetNumTics retrieves the number of tick marks in a trackbar.
//	If no tick flag is set, it returns 2 for the beginning and ending ticks.
//	If TBS_NOTICKS is set, it returns zero. Otherwise,
//	it takes the difference between the range minimum and maximum, divides by the tick frequency, and adds 2.
//
//	The TBM_GETNUMTICS message counts all of the tick marks,
//	including the first and last tick marks created by the trackbar.
func (tb *TrackBar) GetNumTics() int {
	return int(tb.SendMessage(winapi.TBM_GETNUMTICS, 0, 0))
}

//	GetPageSize retrieves the number of logical positions the trackbar's slider moves in response to keyboard input,
//	such as the or keys, or mouse input, such as clicks in the trackbar's channel.
//	The logical positions are the integer increments in the trackbar's range of minimum to maximum slider positions.
//
//	Returns a 32-bit value that specifies the page size for the trackbar.
//	Remarks:
//	The trackbar also sends a WM_HSCROLL or WM_VSCROLL message with the TB_PAGEUP and TB_PAGEDOWN notification codes to
//	its parent window when it receives keyboard or mouse input that scrolls the page.
func (tb *TrackBar) GetPageSize() int {
	return int(tb.SendMessage(winapi.TBM_GETPAGESIZE, 0, 0))
}

//	GetPos retrieves the current logical position of the slider in a trackbar.
//	The logical positions are the integer values in the trackbar's range of minimum to maximum slider positions.
//	Returns a 32-bit value that specifies the current logical position of the trackbar's slider.
func (tb *TrackBar) GetPos() int {
	return int(tb.SendMessage(winapi.TBM_GETPOS, 0, 0))
}

// TODO: TBM_GETPTICS

//	GetRangeMax retrieves the maximum position for the slider in a trackbar.
//	Returns a 32-bit value that specifies the maximum position in the trackbar's range of minimum to
//	maximum slider positions.
func (tb *TrackBar) GetRangeMax() int {
	return int(tb.SendMessage(winapi.TBM_GETRANGEMAX, 0, 0))
}

//	GetRangeMin retrieves the minimum position for the slider in a trackbar.
//	Returns a 32-bit value that specifies the minimum position in the trackbar's range of minimum to
//	maximum slider positions
func (tb *TrackBar) GetRangeMin() int {
	return int(tb.SendMessage(winapi.TBM_GETRANGEMIN, 0, 0))
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
