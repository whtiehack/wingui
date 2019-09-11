package wingui

import "github.com/lxn/win"

func RGB(r, g, b uintptr) win.COLORREF {
	return win.COLORREF(r | g<<8 | b<<16)
}
