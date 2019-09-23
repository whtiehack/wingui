package wingui

import (
	"github.com/lxn/win"
)

// Button a widget for Dialog.
type Button struct {
	WindowBase
	OnClicked func()
}

// WndProc Button window WndProc.
func (b *Button) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	// log.Println("btn wnd proc", b.hwnd, msg, wParam, lParam)
	switch msg {
	case win.WM_COMMAND:
		switch win.HIWORD(uint32(wParam)) {
		case win.BN_CLICKED:
			if b.OnClicked != nil && lParam == uintptr(b.hwnd) {
				b.OnClicked()
			}
		}
	}
	return b.AsWindowBase().WndProc(msg, wParam, lParam)
}

// NewButton create a new Button,need bind to Dialog before use.
func NewButton(idd uintptr) *Button {
	return &Button{WindowBase{idd: idd}, nil}
}

// BindNewButton create a new Button and bind to target dlg.
func BindNewButton(idd uintptr, dlg *Dialog) (*Button, error) {
	b := NewButton(idd)
	err := dlg.BindWidgets(b)
	return b, err
}
