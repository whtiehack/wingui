package wingui

import (
	"github.com/lxn/win"
	"github.com/whtiehack/wingui/winapi"
	"log"
	"syscall"
	"unsafe"
)

type tagTCITEM struct {
	mask        uint32
	dwState     uint32
	dwStateMask uint32
	pszText     uintptr
	cchTextMax  int32
	iImage      int32
	lParam      uintptr
}

// TabControl a widget for Dialog. Tab Control
type TabControl struct {
	WindowBase
	curSel int
	dlgs   []*Dialog
}

// TODO: direct newdialog
// InsertItemText insert tab with text
func (tc *TabControl) InsertItemText(text string, dlg *Dialog) {
	var tagItem tagTCITEM
	tagItem.mask = win.TCIF_TEXT | win.TCIF_IMAGE
	tagItem.iImage = -1
	tagItem.pszText = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text)))
	tagItem.cchTextMax = int32(len(text))
	tagItem.lParam = uintptr(dlg.Handle())
	var newIdx = tc.SendMessage(winapi.TCM_INSERTITEMW, uintptr(len(tc.dlgs)), uintptr(unsafe.Pointer(&tagItem)))
	log.Println("insert item new idx:", newIdx)
	tc.dlgs = append(tc.dlgs, dlg)
	// TODO fix x y
	var rect = dlg.GetWindowRect()
	dlg.SetBounds(Rectangle{
		X:      5,
		Y:      25,
		Width:  int(rect.Right - rect.Left),
		Height: int(rect.Bottom - rect.Top),
	})
}

// WndProc TabControl window WndProc.
func (tc *TabControl) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_NOTIFY:
		nmhdr := (*win.NMHDR)(unsafe.Pointer(lParam))
		switch int(nmhdr.Code) {
		// TODO
		case win.TCN_SELCHANGE:
			log.Println(" TCN_SELCHANGE ", lParam)
			return 1
		}
	}
	return tc.AsWindowBase().WndProc(msg, wParam, lParam)
}

// NewTabControl create a new TabControl, need bind to Dialog before use.
func NewTabControl(idd uintptr) *TabControl {
	return &TabControl{
		WindowBase: WindowBase{idd: idd},
	}
}

// BindTabControl create a new TabControl and bind to target dlg.
func BindTabControl(idd uintptr, dlg *Dialog) (*TabControl, error) {
	tb := NewTabControl(idd)
	err := dlg.BindWidgets(tb)
	return tb, err
}
