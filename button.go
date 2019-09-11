package wingui

import (
	"github.com/lxn/win"
	"log"
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
		if b.OnClicked != nil {
			b.OnClicked()
		}
		return 1
	case win.WM_MOUSEMOVE:
		log.Print("btn mouse move")
		return 0
	}
	return b.AsWindowBase().WndProc(msg, wParam, lParam)
}

// NewButton create a new Button,need bind to Dialog for use.
func NewButton(idd uintptr) *Button {
	return &Button{WindowBase{idd: idd}, nil}
}

// BindNewButton create a new Button and bind to target dlg.
func BindNewButton(idd uintptr, dlg *Dialog) (*Button, error) {
	b := &Button{WindowBase{idd: idd}, nil}
	err := dlg.BindWidgets(b)
	return b, err
}
