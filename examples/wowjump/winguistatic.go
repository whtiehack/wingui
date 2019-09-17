package main

import (
	"github.com/lxn/win"
	"github.com/whtiehack/wingui"
	"os/exec"
	"syscall"
)

// custom Widget.  implement AsWindowBase and WndProc method by Static.
type winguiStatic struct {
	wingui.Static
}

func (w *winguiStatic) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_MOUSEMOVE:
		//var tme win.TRACKMOUSEEVENT
		//tme.CbSize = uint32(unsafe.Sizeof(tme))
		//tme.DwFlags = win.TME_LEAVE | win.TME_HOVER
		//tme.HwndTrack = w.AsWindowBase().Handle()
		//tme.DwHoverTime = 10
		//win.TrackMouseEvent(&tme)
		win.SetCursor(win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_HAND)))
		break
	}
	return w.Static.WndProc(msg, wParam, lParam)
}

func newMyWinguiStatic(idd uintptr) *winguiStatic {
	staticWingui := &winguiStatic{}
	staticWingui.Init(0, idd)
	staticWingui.Subclassing = true
	staticWingui.Color = wingui.RGB(0, 0, 255)
	staticWingui.BkMode = win.TRANSPARENT
	staticWingui.OnClicked = openWinguiLink
	return staticWingui
}

func openWinguiLink() {
	cmd := exec.Command("cmd", "/c", "start", "https://github.com/whtiehack/wingui")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}
