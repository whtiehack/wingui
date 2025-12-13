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

// ItemCount returns the number of tabs.
func (tc *TabControl) ItemCount() int {
	return int(int32(tc.SendMessage(winapi.TCM_GETITEMCOUNT, 0, 0)))
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
	if tc.OnSelChange != nil {
		tc.OnSelChange(index)
	}
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

// InsertDialogPage inserts a tab with the given text at index and binds it to a child dialog used as the page.
// If index is out of range, it appends at the end.
// Returns the actual inserted tab index, or -1 on failure.
func (tc *TabControl) InsertDialogPage(index int, text string, dlg *Dialog) int {
	if dlg == nil {
		return -1
	}

	utf16Text := syscall.StringToUTF16(text)

	var tagItem tagTCITEM
	tagItem.mask = win.TCIF_TEXT | win.TCIF_IMAGE | win.TCIF_PARAM
	tagItem.iImage = -1
	tagItem.pszText = uintptr(unsafe.Pointer(&utf16Text[0]))
	tagItem.cchTextMax = int32(len(utf16Text))
	tagItem.lParam = uintptr(dlg.Handle())

	if index < 0 || index > len(tc.dlgs) {
		index = len(tc.dlgs)
	}
	ret := tc.SendMessage(winapi.TCM_INSERTITEMW, uintptr(index), uintptr(unsafe.Pointer(&tagItem)))
	newIdx := int(int32(ret))
	if newIdx < 0 {
		return -1
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
	return newIdx
}

// AddDialogPage appends a new tab bound to dlg and returns its index (or -1 on failure).
func (tc *TabControl) AddDialogPage(text string, dlg *Dialog) int {
	return tc.InsertDialogPage(len(tc.dlgs), text, dlg)
}

// GetDialogPage returns the page dialog for the given tab index.
func (tc *TabControl) GetDialogPage(index int) *Dialog {
	if index < 0 || index >= len(tc.dlgs) {
		return nil
	}
	return tc.dlgs[index]
}

// CurrentDialogPage returns the currently selected page dialog (if any).
func (tc *TabControl) CurrentDialogPage() *Dialog {
	return tc.GetDialogPage(tc.GetCurSel())
}

// DeleteItem removes the tab at index and updates internal page mapping.
// It does not destroy the page dialog; it will be hidden.
func (tc *TabControl) DeleteItem(index int) bool {
	if index < 0 || index >= tc.ItemCount() {
		return false
	}
	if index >= 0 && index < len(tc.dlgs) && tc.dlgs[index] != nil {
		tc.dlgs[index].Hide()
	}

	ok := tc.SendMessage(winapi.TCM_DELETEITEM, uintptr(index), 0) != 0
	if !ok {
		return false
	}
	if index >= 0 && index < len(tc.dlgs) {
		copy(tc.dlgs[index:], tc.dlgs[index+1:])
		tc.dlgs[len(tc.dlgs)-1] = nil
		tc.dlgs = tc.dlgs[:len(tc.dlgs)-1]
	}

	// Clamp selection and refresh page visibility.
	newSel := tc.GetCurSel()
	if newSel < 0 && tc.ItemCount() > 0 {
		newSel = 0
		tc.SetCurSel(newSel)
	}
	tc.setSelected(newSel)
	return true
}

// WndProc TabControl window WndProc.
func (tc *TabControl) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_NOTIFY:
		nmhdr := (*win.NMHDR)(unsafe.Pointer(lParam))
		// NMHDR.Code is effectively an int32; many notification codes are negative.
		code := int32(nmhdr.Code)
		switch code {
		case int32(win.TCN_SELCHANGING):
			oldIdx := tc.GetCurSel()
			if tc.OnSelChanging != nil && tc.OnSelChanging(oldIdx) {
				// Returning non-zero cancels the selection change.
				return 1
			}
			if oldIdx >= 0 && oldIdx < len(tc.dlgs) && tc.dlgs[oldIdx] != nil {
				tc.dlgs[oldIdx].Hide()
			}
			return 0
		case int32(win.TCN_SELCHANGE):
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
