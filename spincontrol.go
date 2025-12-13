package wingui

import (
	"unsafe"

	"github.com/lxn/win"
)

// SpinControl is a wrapper for the Win32 Up-Down control (msctls_updown32).
type SpinControl struct {
	WindowBase

	// OnDeltaPos fires on UDN_DELTAPOS (before the position changes).
	// Return true to cancel the change.
	OnDeltaPos func(pos, delta int32) (cancel bool)
}

func (sc *SpinControl) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_NOTIFY:
		nmhdr := (*win.NMHDR)(unsafe.Pointer(lParam))
		code := uint32(nmhdr.Code)
		switch code {
		case win.UDN_DELTAPOS:
			if sc.OnDeltaPos != nil {
				nmud := (*win.NMUPDOWN)(unsafe.Pointer(lParam))
				if sc.OnDeltaPos(nmud.IPos, nmud.IDelta) {
					return 1
				}
			}
			return 0
		}
		return 0
	}
	return sc.AsWindowBase().WndProc(msg, wParam, lParam)
}

// SetBuddy sets the buddy window and returns the previous buddy.
func (sc *SpinControl) SetBuddy(buddy win.HWND) win.HWND {
	return win.HWND(sc.SendMessage(win.UDM_SETBUDDY, uintptr(buddy), 0))
}

// GetBuddy returns the current buddy window handle.
func (sc *SpinControl) GetBuddy() win.HWND {
	return win.HWND(sc.SendMessage(win.UDM_GETBUDDY, 0, 0))
}

// SetRange32 sets the 32-bit range.
func (sc *SpinControl) SetRange32(min, max int32) {
	sc.SendMessage(win.UDM_SETRANGE32, uintptr(min), uintptr(max))
}

// GetRange32 gets the 32-bit range.
func (sc *SpinControl) GetRange32() (min, max int32) {
	sc.SendMessage(win.UDM_GETRANGE32, uintptr(unsafe.Pointer(&min)), uintptr(unsafe.Pointer(&max)))
	return
}

// SetPos32 sets the current position and returns the previous position.
func (sc *SpinControl) SetPos32(pos int32) int32 {
	return int32(sc.SendMessage(win.UDM_SETPOS32, 0, uintptr(pos)))
}

// GetPos32 returns the current position.
func (sc *SpinControl) GetPos32() int32 {
	return int32(sc.SendMessage(win.UDM_GETPOS32, 0, 0))
}

// NewSpinControl creates a new SpinControl, need bind to Dialog before use.
func NewSpinControl(idd uintptr) *SpinControl {
	return &SpinControl{WindowBase: WindowBase{idd: idd}}
}

// BindNewSpinControl creates a new SpinControl and bind to target dlg.
func BindNewSpinControl(idd uintptr, dlg *Dialog) (*SpinControl, error) {
	sc := NewSpinControl(idd)
	err := dlg.BindWidgets(sc)
	return sc, err
}

