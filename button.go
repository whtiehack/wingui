package wingui

// Button a widget for Dialog.
type Button struct {
	WindowBase
	OnClicked func()
}

// WndProc Button window WndProc.
func (b *Button) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	// log.Println("btn wnd proc", b.hwnd, msg, wParam, lParam)
	if b.OnClicked != nil {
		b.OnClicked()
	}
	b.AsWindowBase().WndProc(msg, wParam, lParam)
	return 0
}

// NewButton create a new Button,need bind to Dialog for use.
func NewButton(idd uintptr) *Button {
	return &Button{WindowBase{idd: idd}, nil}
}
