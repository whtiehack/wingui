package wingui

import (
	"github.com/lxn/win"
)

// PushButton a widget for Dialog.
type PushButton struct {
	WindowBase
	OnClicked func()
}

// WndProc PushButton window WndProc.
func (b *PushButton) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	// log.Println("btn wnd proc", b.hwnd, msg, wParam, lParam)
	switch msg {
	case win.WM_COMMAND:
		if b.OnClicked != nil && lParam == uintptr(b.hwnd) {
			b.OnClicked()
		}
		return 1
	}
	return b.AsWindowBase().WndProc(msg, wParam, lParam)
}

// NewButton create a new PushButton,need bind to Dialog before use.
func NewButton(idd uintptr) *PushButton {
	return &PushButton{WindowBase{idd: idd}, nil}
}

// BindNewButton create a new PushButton and bind to target dlg.
func BindNewButton(idd uintptr, dlg *Dialog) (*PushButton, error) {
	b := NewButton(idd)
	err := dlg.BindWidgets(b)
	return b, err
}
