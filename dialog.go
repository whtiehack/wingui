package wingui

import (
	"errors"
	"log"
	"strconv"
	"syscall"

	"github.com/lxn/win"
)

// Widget inspect dialog item.
type Widget interface {
	WndProc(msg uint32, wParam, lParam uintptr) uintptr
	AsWindowBase() *WindowBase
	Handle() win.HWND
}

// DialogConfig  TODO.
type DialogConfig struct {
	Style   uint32
	Widgets []Widget
}

// ModalDialogCallBack is modal dialog callback
type ModalDialogCallBack func(dlg *Dialog)

var dlgCount = 0

// Dialog is  main windows struct.
type Dialog struct {
	WindowBase
	items  map[win.HWND]Widget
	config *DialogConfig
	// Indicates whether it is a modal dialog
	cb          ModalDialogCallBack
	wndCallBack uintptr
}

// NewDialog create a new Dialog.
func NewDialog(idd uintptr, parent win.HWND, dialogConfig *DialogConfig) (dlg *Dialog, err error) {
	if dialogConfig == nil {
		dialogConfig = &DialogConfig{}
	}
	dlg = &Dialog{
		items:  make(map[win.HWND]Widget),
		config: dialogConfig,
	}
	dlg.idd = idd
	dlg.parent = parent
	dlg.wndCallBack = syscall.NewCallback(dlg.dialogWndProc)
	h := win.CreateDialogParam(hInstance, win.MAKEINTRESOURCE(idd), parent, dlg.wndCallBack, 0)
	if h == 0 {
		err = errors.New("Create Dialog error:" + strconv.Itoa(int(idd)))
		return
	}
	dlgCount++
	return
}

// NewModalDialog create a new modal dialog
func NewModalDialog(idd uintptr, parent win.HWND, dialogConfig *DialogConfig, cb ModalDialogCallBack) int {
	if dialogConfig == nil {
		dialogConfig = &DialogConfig{}
	}
	dlg := &Dialog{
		items:  make(map[win.HWND]Widget),
		config: dialogConfig,
		cb:     cb,
	}
	dlg.idd = idd
	dlg.parent = parent
	return win.DialogBoxParam(hInstance, win.MAKEINTRESOURCE(idd), parent, syscall.NewCallback(dlg.dialogWndProc), 0)
}

func (dlg *Dialog) dialogWndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	//log.Println("NewDialog.WndProc", hwnd, "msg:", msg, "wparam:", strconv.FormatInt(int64(win.HIWORD(uint32(wParam))), 16), win.LOWORD(uint32(wParam)), "lparam:", lParam)
	// TODO：
	if item, ok := dlg.items[hwnd]; ok {
		return item.WndProc(msg, wParam, lParam)
	}
	if lParam != 0 {
		if item, ok := dlg.items[win.HWND(lParam)]; ok {
			return item.WndProc(msg, wParam, lParam)
		}
	}
	switch msg {
	//case win.WM_ACTIVATEAPP:
	//
	//	return 1
	case win.WM_INITDIALOG:
		log.Println("wm init Dialog", hwnd, dlg.idd)
		dlg.hwnd = hwnd
		if dlg.config.Widgets != nil {
			dlg.BindWidgets(dlg.config.Widgets...)
		}
		if dlg.cb != nil {
			dlgCount++
			dlg.cb(dlg)
			return 0
		}
		return 1
	case win.WM_COMMAND:
		// log.Printf("h:%v WM_COMMAND msg=%v wp %v lp %v   hiwp:%v  lowp:%v\n", dlg.hwnd, msg, wParam, lParam, win.HIWORD(uint32(wParam)), win.LOWORD(uint32(wParam)))
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
		if dlg.cb != nil {
			win.EndDialog(hwnd, 0)
		} else {
			win.DestroyWindow(hwnd)
		}
		return 0
	case win.WM_DESTROY:
		log.Println("WM_DESTROY", hwnd, dlgCount)
		dlgCount--
		if dlgCount == 0 {
			win.PostQuitMessage(0)
		}
		return 0
	}
	return uintptr(0)
}

func (dlg *Dialog) getDlgItem(id uintptr) (h win.HWND, err error) {
	h = win.GetDlgItem(dlg.hwnd, int32(id))
	if h == 0 {
		err = errors.New("GetDlgItem Error:" + strconv.Itoa(int(dlg.hwnd)) + " id:" + strconv.Itoa(int(id)))
		return
	}
	return
}

// SetIcon set Window Icon.
func (dlg *Dialog) SetIcon(id uintptr) {
	h := win.LoadIcon(hInstance, win.MAKEINTRESOURCE(id))
	dlg.AsWindowBase().SetIcon(0, uintptr(h), false)
	dlg.AsWindowBase().SetIcon(1, uintptr(h), false)
}

// BindWidgets bind dialog items.
func (dlg *Dialog) BindWidgets(widgets ...Widget) error {
	var h win.HWND
	var err error
	for _, w := range widgets {
		base := w.AsWindowBase()
		h, err = dlg.getDlgItem(base.idd)
		if err != nil {
			log.Println("BindWidgets error", err, widgets)
			return err
		}
		base.lpPrevWndFunc = win.SetWindowLongPtr(h, win.GWL_WNDPROC, dlg.wndCallBack)

		base.hwnd = h
		base.parent = dlg.hwnd
		dlg.items[h] = w
	}

	return err
}
