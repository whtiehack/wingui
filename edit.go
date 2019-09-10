package wingui

type Edit struct {
	WindowBase
}

func (b *Edit) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	// log.Println("btn wnd proc", b.hwnd, msg, wParam, lParam)
	b.AsWindowBase().WndProc(msg, wParam, lParam)
	return 0
}

