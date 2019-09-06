package wingui

import "github.com/lxn/win"

func rectFromRECT(r win.RECT) Rectangle {
	return Rectangle{
		X:      int(r.Left),
		Y:      int(r.Top),
		Width:  int(r.Right - r.Left),
		Height: int(r.Bottom - r.Top),
	}
}

type Rectangle struct {
	X, Y, Width, Height int
}
