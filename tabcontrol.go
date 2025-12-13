package wingui

import (
	"github.com/lxn/win"
	"github.com/whtiehack/wingui/winapi"
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
	// OnSelChange fires after the current tab selection has changed.
	OnSelChange func(newIndex int)
	// OnSelChanging fires before the current selection changes. If it returns true,
	// the selection change is canceled.
	OnSelChanging func(oldIndex int) (cancel bool)
}

// GetCurSel returns the zero-based index of the selected tab, or -1 if none.
func (tc *TabControl) GetCurSel() int {
	return int(int32(tc.SendMessage(winapi.TCM_GETCURSEL, 0, 0)))
}

// SetCurSel sets the current selection and returns the previous selection index (or -1).
func (tc *TabControl) SetCurSel(index int) int {
	return int(int32(tc.SendMessage(winapi.TCM_SETCURSEL, uintptr(index), 0)))
}

// Select sets the current selection and updates page visibility accordingly.
func (tc *TabControl) Select(index int) {
	tc.SetCurSel(index)
	tc.setSelected(index)
}

func (tc *TabControl) contentRect() (win.RECT, bool) {
	if tc.hwnd == 0 {
		return win.RECT{}, false
	}
	var rc win.RECT
	if !win.GetClientRect(tc.hwnd, &rc) {
		return win.RECT{}, false
	}
	// Translate the tab control client rect into the display rect (excludes the tab headers).
	tc.SendMessage(winapi.TCM_ADJUSTRECT, 0, uintptr(unsafe.Pointer(&rc)))
	return rc, true
}

func (tc *TabControl) layoutPages() {
	rc, ok := tc.contentRect()
	if !ok {
		return
	}
	w, h := rc.Right-rc.Left, rc.Bottom-rc.Top
	if w <= 0 || h <= 0 {
		return
	}
	for _, dlg := range tc.dlgs {
		if dlg == nil || dlg.Handle() == 0 {
			continue
		}
		win.MoveWindow(dlg.Handle(), rc.Left, rc.Top, w, h, true)
	}
}

func (tc *TabControl) setSelected(index int) {
	if index < 0 || index >= len(tc.dlgs) {
		tc.curSel = -1
		for _, dlg := range tc.dlgs {
			if dlg == nil {
				continue
			}
			dlg.Hide()
		}
		return
	}

	tc.curSel = index
	tc.layoutPages()

	for i, dlg := range tc.dlgs {
		if dlg == nil {
			continue
		}
		if i == index {
			dlg.Show()
		} else {
			dlg.Hide()
		}
	}
}

// InsertItemText inserts a tab with the given text and binds it to a child dialog used as the page.
func (tc *TabControl) InsertItemText(text string, dlg *Dialog) {
	if dlg == nil {
		return
	}

	utf16Text := syscall.StringToUTF16(text)

	var tagItem tagTCITEM
	tagItem.mask = win.TCIF_TEXT | win.TCIF_IMAGE | win.TCIF_PARAM
	tagItem.iImage = -1
	tagItem.pszText = uintptr(unsafe.Pointer(&utf16Text[0]))
	tagItem.cchTextMax = int32(len(utf16Text))
	tagItem.lParam = uintptr(dlg.Handle())

	ret := tc.SendMessage(winapi.TCM_INSERTITEMW, uintptr(len(tc.dlgs)), uintptr(unsafe.Pointer(&tagItem)))
	newIdx := int(int32(ret))
	if newIdx < 0 {
		return
	}

	// Keep the page list aligned with tab indexes, even if Windows inserts at a different index.
	tc.dlgs = append(tc.dlgs, nil)
	copy(tc.dlgs[newIdx+1:], tc.dlgs[newIdx:])
	tc.dlgs[newIdx] = dlg

	tc.layoutPages()

	// Ensure only the selected page is visible.
	if tc.curSel < 0 {
		tc.SetCurSel(newIdx)
		tc.setSelected(newIdx)
	} else {
		tc.setSelected(tc.curSel)
	}
}

// WndProc TabControl window WndProc.
func (tc *TabControl) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_NOTIFY:
		nmhdr := (*win.NMHDR)(unsafe.Pointer(lParam))
		switch int(nmhdr.Code) {
		case win.TCN_SELCHANGING:
			oldIdx := tc.GetCurSel()
			if tc.OnSelChanging != nil && tc.OnSelChanging(oldIdx) {
				// Returning non-zero cancels the selection change.
				return 1
			}
			if oldIdx >= 0 && oldIdx < len(tc.dlgs) && tc.dlgs[oldIdx] != nil {
				tc.dlgs[oldIdx].Hide()
			}
			return 0
		case win.TCN_SELCHANGE:
			newIdx := tc.GetCurSel()
			tc.setSelected(newIdx)
			if tc.OnSelChange != nil {
				tc.OnSelChange(newIdx)
			}
			return 0
		}
		// WM_NOTIFY is routed by the parent dialog; do not forward it to the control's WndProc.
		return 0
	case win.WM_SIZE, win.WM_WINDOWPOSCHANGED:
		// Keep pages sized to the tab display area.
		tc.layoutPages()
	}
	return tc.AsWindowBase().WndProc(msg, wParam, lParam)
}

// NewTabControl create a new TabControl, need bind to Dialog before use.
func NewTabControl(idd uintptr) *TabControl {
	return &TabControl{
		WindowBase: WindowBase{idd: idd, Subclassing: true},
		curSel:     -1,
	}
}

// BindTabControl create a new TabControl and bind to target dlg.
func BindTabControl(idd uintptr, dlg *Dialog) (*TabControl, error) {
	tb := NewTabControl(idd)
	err := dlg.BindWidgets(tb)
	return tb, err
}
