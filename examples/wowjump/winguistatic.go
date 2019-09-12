package main

import (
	"github.com/lxn/win"
	"github.com/whtiehack/wingui"
	"os/exec"
	"syscall"
)

type winguiStatic struct {
	widget *wingui.Static
}

func (w *winguiStatic) AsWindowBase() *wingui.WindowBase {
	return w.widget.AsWindowBase()
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
	return w.widget.WndProc(msg, wParam, lParam)
}

func newMyWinguiStatic(idd uintptr) *winguiStatic {
	staticWingui := wingui.NewStatic(IDS_WINGUI)
	staticWingui.Subclassing = true
	staticWingui.Color = wingui.RGB(0, 0, 255)
	staticWingui.BkMode = win.TRANSPARENT
	staticWingui.OnClicked = openWinguiLink
	return &winguiStatic{widget: staticWingui}
}

func openWinguiLink() {
	cmd := exec.Command("cmd", "/c", "start", "https://github.com/whtiehack/wingui")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}
