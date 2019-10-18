// Package wingui https://docs.microsoft.com/zh-cn/windows/win32/controls/progress-bar-control#using-progress-bars
package wingui

import (
	"github.com/lxn/win"
	"github.com/whtiehack/wingui/winapi"
	"unsafe"
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

// GetPos Retrieves the current position of the progress bar.
//
// Returns a UINT value that represents the current position of the progress bar.
func (pb *ProgressBar) GetPos() int {
	return int(pb.SendMessage(win.PBM_GETPOS, 0, 0))
}

// SetPos Sets the current position for a progress bar and redraws the bar to reflect the new position.
//
// pos:Signed integer that becomes the new position.
//
// Returns the previous position.
//
// 	Remarks
//	If pos is outside the range of the control, the position is set to the closest boundary.
//
//	Do not send this message to a control that has the PBS_MARQUEE style.
func (pb *ProgressBar) SetPos(pos int) int {
	return int(pb.SendMessage(win.PBM_SETPOS, uintptr(pos), 0))
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

// SetBarColor Sets the color of the progress indicator bar in the progress bar control.
//
// The COLORREF value that specifies the new progress indicator bar color.
// Specifying the CLR_DEFAULT value causes the progress bar to use its default progress indicator bar color.
//
// Returns the previous progress indicator bar color,
// or CLR_DEFAULT if the progress indicator bar color is the default color.
func (pb *ProgressBar) SetBarColor(color int) int {
	return int(pb.SendMessage(winapi.PBM_SETBARCOLOR, 0, uintptr(color)))
}

// GetBkColor Gets the background color of the progress bar.
//
// Returns the background color of the progress bar.
//
// This is the color set by the PBM_SETBKCOLOR message.
// The default value is CLR_DEFAULT, which is defined in commctrl.h.
//
// This function only affects the classic mode, not any visual style.
func (pb *ProgressBar) GetBkColor() int {
	return int(pb.SendMessage(winapi.PBM_GETBKCOLOR, 0, 0))
}

// SetBkColor Sets the background color in the progress bar.
//
// COLORREF value that specifies the new background color.
// Specify the CLR_DEFAULT value to cause the progress bar to use its default background color.
//
// Returns the previous background color, or CLR_DEFAULT if the background color is the default color.
//
// 	Remarks
//	When visual styles are enabled, this message has no effect.
func (pb *ProgressBar) SetBkColor(color int) int {
	return int(pb.SendMessage(winapi.PBM_SETBKCOLOR, 0, uintptr(color)))
}

// GetRange Retrieves information about the current high and low limits of a given progress bar control.
func (pb *ProgressBar) GetRange() (low int, high int) {
	var pbrange winapi.PBRANGE
	pb.SendMessage(winapi.PBM_GETRANGE, 0, uintptr(unsafe.Pointer(&pbrange)))
	return pbrange.Low, pbrange.High
}

// SetRange Sets the minimum and maximum values for a progress bar and redraws the bar to reflect the new range.
//
// The low specifies the minimum range value, and the high specifies the maximum range value.
// The minimum range value must not be negative. By default, the minimum value is zero.
// The maximum range value must be greater than the minimum range value. By default, the maximum range value is 100.
//
// 	Returns the previous range values if successful, or zero otherwise.
// 	The LOWORD specifies the previous minimum value, and the HIWORD specifies the previous maximum value.
//
//	Remarks
//	If you do not set the range values, the system sets the minimum value to 0 and the maximum value to 100.
//	Because this message expresses the range as a 16-bit unsigned integer, it can extend from 0 to 65,535.
//	The minimum value in the range can be from 0 to 65,535. Likewise, the maximum value can be from 0 to 65,535.
//
//	To set a larger range, call PBM_SETRANGE32.
func (pb *ProgressBar) SetRange(low, high int) int {
	return int(pb.SendMessage(winapi.PBM_SETRANGE, 0, uintptr(win.MAKELONG(uint16(low), uint16(high)))))
}

// SetRange32 Sets the minimum and maximum values for a progress bar to 32-bit values,
// and redraws the bar to reflect the new range.
func (pb *ProgressBar) SetRange32(low, high int) int {
	return int(pb.SendMessage(winapi.PBM_SETRANGE, uintptr(low), uintptr(high)))
}

// GetState Gets the state of the progress bar.
//
// Returns the current state of the progress bar. One of the following values.
//
//	Return code	Description
//	PBST_NORMAL
//	In progress.
//
//	PBST_ERROR
//	Error.
//
//	PBST_PAUSED
//	Paused.
func (pb *ProgressBar) GetState() int {
	return int(pb.SendMessage(winapi.PBM_GETSTATE, 0, 0))
}

// SetState  Sets the state of the progress bar.
//	State of the progress bar that is being set. One of the following values.
//
//	Value	Meaning
//	PBST_NORMAL
//	In progress.
//
//	PBST_ERROR
//	Error.
//
//	PBST_PAUSED
//	Paused.
//
// Returns the previous state.
func (pb *ProgressBar) SetState(state int) int {
	return int(pb.SendMessage(winapi.PBM_SETSTATE, uintptr(state), 0))
}

// GetStep Retrieves the step increment from a progress bar. The step increment is the amount by which
// the progress bar increases its current position whenever it receives a PBM_STEPIT message.
// By default, the step increment is set to 10.
//
// Returns the current step increment.
func (pb *ProgressBar) GetStep() int {
	return int(pb.SendMessage(winapi.PBM_GETSTEP, 0, 0))
}

// SetStep Specifies the step increment for a progress bar. The step increment is the amount by
// which the progress bar increases its current position whenever it receives a PBM_STEPIT message.
// By default, the step increment is set to 10.
//
// step: New step increment.
//
// Returns the previous step increment.
func (pb *ProgressBar) SetStep(step int) int {
	return int(pb.SendMessage(winapi.PBM_SETSTEP, uintptr(step), 0))
}

// SetPit Advances the current position for a progress bar by the step increment and redraws
// the bar to reflect the new position. An application sets the step increment by sending the PBM_SETSTEP message.
//
// Returns the previous position.
//
// When the position exceeds the maximum range value, this message resets the current position
// so that the progress indicator starts over again from the beginning.
func (pb *ProgressBar) SetPit() int {
	return int(pb.SendMessage(winapi.PBM_STEPIT, 0, 0))
}

// SetMarquee  Sets the progress bar to marquee mode. This causes the progress bar to move like a marquee.
//
// enable,indicates whether to turn the marquee mode on or off.
//
// Time, in milliseconds, between marquee animation updates. If this parameter is zero,
// the marquee animation is updated every 30 milliseconds.
//
//	Remarks
//	Use this message when you do not know the amount of progress toward completion
//	but wish to indicate that progress is being made.
//
//	Send the PBM_SETMARQUEE message to start or stop the animation.
func (pb *ProgressBar) SetMarquee(enable bool, time int) {
	turnOn := 1
	if !enable {
		turnOn = 0
	}
	pb.SendMessage(winapi.PBM_SETMARQUEE, uintptr(turnOn), uintptr(time))
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
