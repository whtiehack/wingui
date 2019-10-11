package wingui

import (
	"log"
	"unsafe"

	"github.com/lxn/win"
	"github.com/whtiehack/wingui/winapi"
)

// https://docs.microsoft.com/zh-cn/windows/win32/controls/trackbar-controls

//	TrackBar a widget for Dialog. TrackBar
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

//	SetBuddy assigns a window as the buddy window for a trackbar control.
//	Trackbar buddy windows are automatically displayed in a location relative to the control's
//	orientation (horizontal or vertical).
//	isLeftOrTop indicating which buddy window handle will be retrieved, by relative location.
//	This value can be one of the following:
//	- TRUE
//	- Retrieves the handle to the buddy to the left of the trackbar. If the trackbar control uses the TBS_VERT style, the message will retrieve the buddy above the trackbar.
//	- FALSE
//	- Retrieves the handle to the buddy to the right of the trackbar. If the trackbar control uses the TBS_VERT style, the message will retrieve the buddy below the trackbar.
//
//	hWnd is a Handle to the window that will be set as the trackbar control's buddy.
//	Returns the handle to the window that was previously assigned to the control at that location.
func (tb *TrackBar) SetBuddy(isLeftOrTop bool, hWnd win.HWND) win.HWND {
	rd := 0
	if isLeftOrTop {
		rd = 1
	}
	return win.HWND(tb.SendMessage(winapi.TBM_SETBUDDY, uintptr(rd), uintptr(hWnd)))
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

//	SetLineSize sets the number of logical positions the trackbar's slider moves in response to keyboard
//	input from the arrow keys, such as the or keys. The logical positions are the integer increments in the
//	trackbar's range of minimum to maximum slider positions.
//	size is new line size.
//	Returns a 32-bit value that specifies the previous line size.
//
//	Remarks
//	The default setting for the line size is 1.
//
//	The trackbar also sends a WM_HSCROLL or WM_VSCROLL message with the TB_LINEUP and TB_LINEDOWN notification codes
//	to its parent window when the user presses the arrow keys.
func (tb *TrackBar) SetLineSize(size int) int {
	return int(tb.SendMessage(winapi.TBM_SETLINESIZE, 0, uintptr(size)))
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

//	GetTic retrieves the logical position of a tick mark in a trackbar. The logical position can be any of the integer
//	values in the trackbar's range of minimum to maximum slider positions.
//	idx is a zero-based index identifying a tick mark. Valid indexes are in the range from zero to two less
//	than the tick count returned by the TBM_GETNUMTICS message.
//	Returns the logical position of the specified tick mark, or -1 if wParam does not specify a valid index.
func (tb *TrackBar) GetTic(idx int) int {
	return int(tb.SendMessage(winapi.TBM_GETTIC, uintptr(idx), 0))
}

//	SetTic sets a tick mark in a trackbar at the specified logical position.
//	position is position of the tick mark. This parameter can be any of the integer values
//	in the trackbar's range of minimum to maximum slider positions.
//	Returns TRUE if the tick mark is set, or FALSE otherwise.
func (tb *TrackBar) SetTic(position int) bool {
	return tb.SendMessage(winapi.TBM_SETTIC, 0, uintptr(position)) == 1
}

//	SetTicFreq Sets the interval frequency for tick marks in a trackbar. For example,
//	if the frequency is set to two, a tick mark is displayed for every other increment in the trackbar's range.
//	The default setting for the frequency is one; that is, every increment in the range is associated with a tick mark.
// 	frequency is frequency of the tick marks.
//	Remarks
//	The trackbar must have the TBS_AUTOTICKS style to use this message.
func (tb *TrackBar) SetTicFreq(frequency int) {
	tb.SendMessage(winapi.TBM_SETTIC, uintptr(frequency), 0)
}

//	GetTicPos retrieves the current physical position of a tick mark in a trackbar.
//	idx is a zero-based index identifying a tick mark. The positions of the first and last tick marks are not
//	directly available via this message.
//
//	Returns the distance, in client coordinates, from the left or top of the trackbar's client area to
//	the specified tick mark. The return value is the x-coordinate of the tick mark for a horizontal trackbar or
//	the y-coordinate for a vertical trackbar. If wParam is not a valid index, the return value is -1.
//
//	Remarks
//	Because the first and last tick marks are not available through this message, valid indexes are offset from
//	their tick position on the trackbar. If the difference between TBM_GETRANGEMIN and TBM_GETRANGEMAX is less than two,
//	then there is no valid index and this message will fail.
//
//	The following illustrates the relation between the ticks on a trackbar, the ticks available through this message,
//	and their zero-based indexes.
//
//	0 1 2 3 4 5 6 7 8 9    // Tick positions seen on the trackbar.
//    1 2 3 4 5 6 7 8      // Tick positions whose position can be identified.
//    0 1 2 3 4 5 6 7      // Index numbers for the identifiable positions.
func (tb *TrackBar) GetTicPos(idx int) int {
	return int(tb.SendMessage(winapi.TBM_GETTICPOS, uintptr(idx), 0))
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

//	SetPageSize sets the number of logical positions the trackbar's slider moves in response to keyboard input,
//	such as the or keys, or mouse input, such as clicks in the trackbar's channel. The logical positions are the
//	integer increments in the trackbar's range of minimum to maximum slider positions.
//	Returns a 32-bit value that specifies the previous page size.
//
//	Remarks
//	The trackbar also sends a WM_HSCROLL or WM_VSCROLL message with the TB_PAGEUP and TB_PAGEDOWN notification codes to
//	its parent window when it receives keyboard or mouse input that scrolls the page.
func (tb *TrackBar) SetPageSize(newSize int) int {
	return int(tb.SendMessage(winapi.TBM_SETPAGESIZE, 0, uintptr(newSize)))
}

//	GetPos retrieves the current logical position of the slider in a trackbar.
//	The logical positions are the integer values in the trackbar's range of minimum to maximum slider positions.
//	Returns a 32-bit value that specifies the current logical position of the trackbar's slider.
func (tb *TrackBar) GetPos() int {
	return int(tb.SendMessage(winapi.TBM_GETPOS, 0, 0))
}

//	SetPos sets the current logical position of the slider in a trackbar.
//
//	newPose is new logical position of the slider. Valid logical positions are the integer values
//	in the trackbar's range of minimum to maximum slider positions. If this value is outside the control's
//	maximum and minimum range, the position is set to the maximum or minimum value.
func (tb *TrackBar) SetPos(redraw bool, newPos int) {
	rd := 0
	if redraw {
		rd = 1
	}
	tb.SendMessage(winapi.TBM_SETPOS, uintptr(rd), uintptr(newPos))
}

//	SetPosNotify sets the current logical position of the slider in a trackbar.
//	newPost New logical position of the slider. Valid logical positions are the integer values in the trackbar's
//	range of minimum to maximum slider positions. If this value is outside the control's maximum and minimum range,
//	the position is set to the maximum or minimum value.
//
//	Remarks
//	Calling TBM_SETPOSNOTIFY will set the trackbar slider location like TBM_SETPOS would,
//	but it will also cause the trackbar to notify its parent of a move via a WM_HSCROLL or WM_VSCROLL message.
func (tb *TrackBar) SetPosNotify(newPos int) {
	tb.SendMessage(winapi.TBM_SETPOSNOTIFY, 0, uintptr(newPos))
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

//	SetRange sets the range of minimum and maximum logical positions for the slider in a trackbar.
//
//	Remarks
//	If the current slider position is outside the new range,
//	the TBM_SETRANGE message sets the slider position to the new maximum or minimum value.
//
//	Because this message takes two 16-bit unsigned integer values,
//	the maximum range that this message can specify is from 0 to 65,535.
//	To specify larger range values, use the TBM_SETRANGEMIN and TBM_SETRANGEMAX messages.
func (tb *TrackBar) SetRange(redraw bool, min, max uint16) {
	rd := 0
	if redraw {
		rd = 1
	}
	tb.SendMessage(winapi.TBM_SETRANGE, uintptr(rd), uintptr(win.MAKELONG(min, max)))
}

//	SetRangeMax Sets the maximum logical position for the slider in a trackbar.
//	max Maximum position for the slider.
//
//	Remarks
//	If the current slider position is greater than the new maximum,
//	the TBM_SETRANGEMAX message sets the slider position to the new maximum value.
func (tb *TrackBar) SetRangeMax(redraw bool, max int) {
	rd := 0
	if redraw {
		rd = 1
	}
	tb.SendMessage(winapi.TBM_SETRANGEMAX, uintptr(rd), uintptr(max))
}

//	SetRangeMin Sets the minimum logical position for the slider in a trackbar.
//	min Minimum position for the slider.
//
//	Remarks
//	If the current slider position is less than the new minimum,
//	the TBM_SETRANGEMIN message sets the slider position to the new minimum value.
func (tb *TrackBar) SetRangeMin(redraw bool, min int) {
	rd := 0
	if redraw {
		rd = 1
	}
	tb.SendMessage(winapi.TBM_SETRANGEMIN, uintptr(rd), uintptr(min))
}

//	GetSelEnd retrieves the ending position of the current selection range in a trackbar.
//	Returns a 32-bit value that specifies the ending position of the current selection range.
//	A trackbar can have a selection range only if you specified the TBS_ENABLESELRANGE style when you created it.
func (tb *TrackBar) GetSelEnd() int {
	return int(tb.SendMessage(winapi.TBM_GETSELEND, 0, 0))
}

//	GetSelStart retrieves the starting position of the current selection range in a trackbar.
//	Returns a 32-bit value that specifies the starting position of the current selection range.
//	A trackbar can have a selection range only if you specified the TBS_ENABLESELRANGE style when you created it.
func (tb *TrackBar) GetSelStart() int {
	return int(tb.SendMessage(winapi.TBM_GETSELSTART, 0, 0))
}

//	SetSel Sets the starting and ending positions for the available selection range in a trackbar.
func (tb *TrackBar) SetSel(redraw bool, start, end uint16) {
	rd := 0
	if redraw {
		rd = 1
	}
	tb.SendMessage(winapi.TBM_SETSEL, uintptr(rd), uintptr(win.MAKELONG(start, end)))
}

//	SetSelEnd Sets the ending logical position of the current selection range in a trackbar.
//	This message is ignored if the trackbar does not have the TBS_ENABLESELRANGE style.
//	end Ending logical position of the selection range.
func (tb *TrackBar) SetSelEnd(redraw bool, end int) {
	rd := 0
	if redraw {
		rd = 1
	}
	tb.SendMessage(winapi.TBM_SETSELEND, uintptr(rd), uintptr(end))
}

//	SetSelStart Sets the starting logical position of the current selection range in a trackbar.
//	This message is ignored if the trackbar does not have the TBS_ENABLESELRANGE style.
//	start Starting position of the selection range.
func (tb *TrackBar) SetSelStart(redraw bool, start int) {
	rd := 0
	if redraw {
		rd = 1
	}
	tb.SendMessage(winapi.TBM_SETSELSTART, uintptr(rd), uintptr(start))
}

//	GetThumbLength retrieves the length of the slider in a trackbar.
//	Returns the length, in pixels, of the slider.
func (tb *TrackBar) GetThumbLength() int {
	return int(tb.SendMessage(winapi.TBM_GETTHUMBLENGTH, 0, 0))
}

//	SetThumbLength sets the length of the slider in a trackbar.
//	This message is ignored if the trackbar does not have the TBS_FIXEDLENGTH style.\
//	length, in pixels, of the slider.
func (tb *TrackBar) SetThumbLength(length int) {
	tb.SendMessage(winapi.TBM_SETTHUMBLENGTH, uintptr(length), 0)
}

//	GetThumbRect retrieves the size and position of the bounding rectangle for the slider in a trackbar.
func (tb *TrackBar) GetThumbRect() win.RECT {
	var rect win.RECT
	tb.SendMessage(winapi.TBM_GETTHUMBRECT, 0, uintptr(unsafe.Pointer(&rect)))
	return rect
}

//	GetTooltips retrieves the handle to the tooltip control assigned to the trackbar, if any.
//	Returns the handle to the tooltip control assigned to the trackbar, or NULL if tooltips are not in use.
//	If the trackbar control does not use the TBS_TOOLTIPS style, the return value is NULL.
func (tb *TrackBar) GetTooltips() win.HWND {
	return win.HWND(tb.SendMessage(winapi.TBM_GETTOOLTIPS, 0, 0))
}

//	SetTooltips assigns a tooltip control to a trackbar control.
//	hWnd is a handle to an existing tooltip control.
//	Remarks
//	When a trackbar control is created with the TBS_TOOLTIPS style,
//	it creates a default tooltip control that appears next to the slider, displaying the slider's current position.
func (tb *TrackBar) SetTooltips(hWnd win.HWND) {
	tb.SendMessage(winapi.TBM_SETTOOLTIPS, uintptr(hWnd), 0)
}

//	SetTipSide positions a tooltip control used by a trackbar control.
//	Trackbar controls that use the TBS_TOOLTIPS style display tooltips.
//	flag representing the location at which to display the tooltip control. This value can be one of the following:
//	TBTS_TOP
//	The tooltip control will be positioned above the trackbar. This flag is for use with horizontal trackbars.
//	TBTS_LEFT
//	The tooltip control will be positioned to the left of the trackbar. This flag is for use with vertical trackbars.
//	TBTS_BOTTOM
//	The tooltip control will be positioned below the trackbar. This flag is for use with horizontal trackbars.
//	TBTS_RIGHT
//	The tooltip control will be positioned to the right of the trackbar. This flag is for use with vertical trackbars.
//
//	Returns a value that represents the tooltip control's previous location.
//	The return value equals one of the possible values for flag.
func (tb *TrackBar) SetTipSide(flag int) int {
	return int(tb.SendMessage(winapi.TBM_SETTIPSIDE, uintptr(flag), 0))
}

//	 SetUnicodeFormat sets the Unicode character format flag for the control.
//	 This message allows you to change the character set used by the control
//	 at run time rather than having to re-create the control.
//	unicode Determines the character set that is used by the control. If this value is nonzero,
//	the control will use Unicode characters. If this value is zero, the control will use ANSI characters.
//	Return value
//	Returns the previous Unicode format flag for the control.
func (tb *TrackBar) SetUnicodeFormat(unicode int) int {
	return int(tb.SendMessage(winapi.TBM_SETUNICODEFORMAT, uintptr(unicode), 0))
}

// WndProc TrackBar window WndProc.
func (tb *TrackBar) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_NOTIFY:
		nmhdr := (*win.NMHDR)(unsafe.Pointer(lParam))
		switch nmhdr.Code {
		// TODO
		case win.TRBN_THUMBPOSCHANGING:
			log.Println(" TRBN_THUMBPOSCHANGING ", lParam)
			return 1
		}
	}
	return tb.AsWindowBase().WndProc(msg, wParam, lParam)
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
