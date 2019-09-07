package wingui

import (
	"errors"
	"github.com/lxn/win"
	"log"
	"strconv"
	"syscall"
)

type WndProc interface {
	WndProc(msg uint32, wParam, lParam uintptr) uintptr
}

var dlgCount = 0

type dialog struct {
	WindowBase
	items map[win.HWND]WndProc
}

func NewDialog(idd uintptr, parent win.HWND) (dlg *dialog, err error) {
	dlg = &dialog{
		items: make(map[win.HWND]WndProc),
	}
	h := win.CreateDialogParam(hInstance, win.MAKEINTRESOURCE(idd), parent, syscall.NewCallback(dlg.dialogWndProc), 0)
	if h == 0 {
		err = errors.New("Create dialog error:" + strconv.Itoa(int(idd)))
		return
	}
	dlg.hwnd = h
	dlg.idd = idd
	dlg.parent = parent
	dlgCount++
	return
}

func (dlg *dialog) dialogWndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	//log.Println("NewDialog.WndProc", dlg, hwnd, msg, wParam, lParam)
	switch msg {
	case win.WM_INITDIALOG:
		log.Println("wm init dialog")
		return 1
	case win.WM_COMMAND:
		// log.Printf("WM_COMMAND msg=%v wp %v lp %v\n", msg, wParam, lParam)
		if lParam != 0 {
			h := win.HWND(lParam)
			if item, ok := dlg.items[h]; ok {
				item.WndProc(msg, wParam, lParam)
			}
		}
		//log.Printf("WM_COMMAND msg=%v\n", msg)
		//if lParam != 0 { //Reflect message to control
		//	h := win.HWND(lParam)
		//	log.Printf("WM_COMMAND h=%v\n", h)
		//	if handler := GetMsgHandler(h); handler != nil {
		//		log.Println("WM_COMMAND handler.WndProc")
		//		ret := handler.WndProc(msg, wParam, lParam)
		//		if ret != 0 {
		//			//win.SetWindowLong(hwnd, win.DWL_MSGRESULT, int32(ret))
		//			log.Println("WM_COMMAND TRUE")
		//			return win.TRUE
		//		}
		//	}
		//}
		// log.Println("WM_COMMAND DONE")
		return 0
	case win.WM_CLOSE:
		log.Println("WM_CLOSE", hwnd)
		win.DestroyWindow(hwnd)
		return 0
	case win.WM_DESTROY:
		log.Println("WM_DESTROY", hwnd)
		dlgCount--
		if dlgCount == 0 {
			win.PostQuitMessage(0)
		}
		return 0
	}
	return uintptr(0)
}

func (dlg *dialog) getDlgItem(id uintptr) (h win.HWND, err error) {
	h = win.GetDlgItem(dlg.hwnd, int32(id))
	if h == 0 {
		err = errors.New("GetDlgItem Error:" + strconv.Itoa(int(dlg.hwnd)) + " id:" + strconv.Itoa(int(id)))
		return
	}
	return
}

// 绑定控件
func (dlg *dialog) NewButton(id uintptr, winBase **Button) (err error) {
	var h win.HWND
	h, err = dlg.getDlgItem(id)
	if err != nil {
		return
	}
	btn := &Button{
		WindowBase: WindowBase{
			hwnd:   h,
			idd:    id,
			parent: dlg.hwnd,
		},
		OnClicked: nil,
	}
	dlg.items[h] = btn
	*winBase = btn
	return
}
