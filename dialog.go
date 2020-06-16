package wingui

import (
	"errors"
	"log"
	"strconv"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

// Widget inspect dialog item.
type Widget interface {
	WndProc(msg uint32, wParam, lParam uintptr) uintptr
	AsWindowBase() *WindowBase
}

// DialogConfig  TODO.
type DialogConfig struct {
	Style uint32
	//Widgets will be bind when dialog init.
	Widgets []Widget
}

// ModalDialogCallBack is modal dialog callback
type ModalDialogCallBack func(dlg *Dialog)

var dlgCount = 0

// Dialog is  main windows struct.
type Dialog struct {
	WindowBase
	items  map[win.HWND]Widget
	iddMap map[uintptr]Widget
	config *DialogConfig
	// Indicates whether it is a modal dialog
	cb          ModalDialogCallBack
	wndCallBack uintptr
	// TODO Support optionsl sub wnd class.
}

//func NewDialogIndirect(idd uintptr, parent win.HWND, dialogConfig *DialogConfig) (dlg *Dialog, err error) {
//	dlg, err = NewDialog(idd, parent, dialogConfig)
//	if err != nil {
//		return
//	}
//	var rect = dlg.GetWindowRect()
//	dlg.SetBounds(Rectangle{
//		X:      0,
//		Y:      0,
//		Width:  int(rect.Right - rect.Left),
//		Height: int(rect.Bottom - rect.Top),
//	})
//	return
//}

// NewDialog create a new Dialog.
func NewDialog(idd uintptr, parent win.HWND, dialogConfig *DialogConfig) (dlg *Dialog, err error) {
	if dialogConfig == nil {
		dialogConfig = &DialogConfig{}
	}
	dlg = &Dialog{
		items:  make(map[win.HWND]Widget),
		iddMap: make(map[uintptr]Widget),
		config: dialogConfig,
	}
	dlg.idd = idd
	dlg.wndCallBack = syscall.NewCallback(dlg.dialogWndProc)
	h := win.CreateDialogParam(hInstance, win.MAKEINTRESOURCE(idd), parent, dlg.wndCallBack, 0)
	if h == 0 {
		err = errors.New("Create Dialog error:" + strconv.Itoa(int(idd)))
		return
	}
	dlgCount++
	return
}

// NewModalDialog create a new modal Dialog.
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
	return win.DialogBoxParam(hInstance, win.MAKEINTRESOURCE(idd), parent, syscall.NewCallback(dlg.dialogWndProc), 0)
}

func (dlg *Dialog) dialogWndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	//log.Println("NewDialog.WndProc", hwnd, "msg:", msg, "wparam:", strconv.FormatInt(int64(win.HIWORD(uint32(wParam))), 16), win.LOWORD(uint32(wParam)), "lparam:", lParam)
	// process subclassing items
	if dlg.hwnd != hwnd {
		if item, ok := dlg.items[hwnd]; ok {
			return item.WndProc(msg, wParam, lParam)
		}
	}
	if lParam != 0 && dlg.hwnd == hwnd {
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
	//case win.WM_COMMAND:
	//	// process WM_COMMAND for dlg items.
	//	return 0
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
	case win.WM_DELETEITEM:
	// TODO
	// Sent to the owner of a list box or combo box when the list box or combo box is destroyed or
	// when items are removed by the LB_DELETESTRING, LB_RESETCONTENT, CB_DELETESTRING, or CB_RESETCONTENT message.
	// The system sends a WM_DELETEITEM message for each deleted item. The system sends the WM_DELETEITEM message
	// for any deleted list box or combo box item with nonzero item data.
	// Parameters
	//wParam
	//Specifies the identifier of the control that sent the WM_DELETEITEM message.
	//
	//lParam
	//Pointer to a DELETEITEMSTRUCT structure that contains information about the item deleted from a list box.
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-deleteitemstruct
	//
	//Return value
	//An application should return TRUE if it processes this message.
	case win.WM_NOTIFY:
		// trans message to child.
		nmhdr := (*win.NMHDR)(unsafe.Pointer(lParam))
		if item, ok := dlg.items[nmhdr.HwndFrom]; ok {
			return item.WndProc(msg, wParam, lParam)
		}
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
		// support custom widgets.
		h = base.Handle()
		if h == 0 {
			h, err = dlg.getDlgItem(base.idd)
		}
		if err != nil {
			log.Println("BindWidgets error", err, widgets)
			return err
		}
		if base.Subclassing {
			log.Println("sub classing idd", base.idd, h)
			base.lpPrevWndFunc = win.SetWindowLongPtr(h, win.GWL_WNDPROC, dlg.wndCallBack)
		}
		base.hwnd = h
		dlg.items[h] = w
		dlg.iddMap[base.idd] = w
	}

	return err
}

func (dlg *Dialog) GetWidget(idd uintptr) Widget {
	return dlg.iddMap[idd]
}
