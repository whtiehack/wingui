package wingui

import (
	"errors"
	"github.com/lxn/win"
	"log"
	"strconv"
	"syscall"
)

var dlgCount = 0

type dialog struct {
	WindowBase
}

func NewDialog(idd uintptr, parent win.HWND) (dlg *dialog, err error) {
	dlg = &dialog{
	}
	h := win.CreateDialogParam(hInstance, win.MAKEINTRESOURCE(idd), parent, syscall.NewCallback(dlg.WndProc), 0)
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

func (dlg *dialog) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	//log.Println("NewDialog.WndProc", dlg, hwnd, msg, wParam, lParam)
	switch msg {
	case win.WM_INITDIALOG:
		log.Println("wm init dialog")
		return 1
	case win.WM_COMMAND:
		log.Printf("WM_COMMAND msg=%v\n", msg)
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
		log.Println("WM_COMMAND DONE")
		return 0
	case win.WM_CLOSE:
		log.Println("WM_CLOSE")
		win.DestroyWindow(hwnd)
		return 0
	case win.WM_DESTROY:
		log.Println("WM_DESTROY")
		dlgCount--
		if dlgCount == 0 {
			win.PostQuitMessage(0)
		}
		return 0
	}
	return uintptr(0)
}

// 绑定控件
