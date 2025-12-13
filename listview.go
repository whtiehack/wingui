package wingui

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"github.com/whtiehack/wingui/winapi"
)

// ListView is a wrapper for the Win32 SysListView32 control.
// Notes:
// - Most notifications are delivered as WM_NOTIFY to the parent dialog (routed by Dialog.dialogWndProc).
// - For report view with columns, prefer creating the control with LVS_REPORT style in resources.
type ListView struct {
	WindowBase

	// OnItemActivate fires on item activation (double click / Enter), if available.
	OnItemActivate func(index int)
	// OnItemChanged fires on LVN_ITEMCHANGED.
	OnItemChanged func(index int)
}

func (lv *ListView) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_NOTIFY:
		nmhdr := (*win.NMHDR)(unsafe.Pointer(lParam))
		code := uint32(nmhdr.Code)

		switch code {
		case win.LVN_ITEMCHANGED:
			if lv.OnItemChanged != nil {
				nmlv := (*win.NMLISTVIEW)(unsafe.Pointer(lParam))
				lv.OnItemChanged(int(nmlv.IItem))
			}
			return 0
		case win.NM_DBLCLK:
			if lv.OnItemActivate != nil {
				nmia := (*win.NMITEMACTIVATE)(unsafe.Pointer(lParam))
				lv.OnItemActivate(int(nmia.IItem))
			}
			return 0
		}
		return 0
	}
	return lv.AsWindowBase().WndProc(msg, wParam, lParam)
}

// SetExtendedStyle sets extended list-view styles (e.g. LVS_EX_FULLROWSELECT).
func (lv *ListView) SetExtendedStyle(style uint32) {
	lv.SendMessage(winapi.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, uintptr(style))
}

// InsertColumn inserts a column at index and returns the actual inserted index (or -1).
func (lv *ListView) InsertColumn(index int, text string, width int, fmt int32) int {
	utf16 := syscall.StringToUTF16(text)
	col := win.LVCOLUMN{
		Mask:    win.LVCF_TEXT | win.LVCF_WIDTH | win.LVCF_FMT | win.LVCF_SUBITEM,
		Fmt:     fmt,
		Cx:      int32(width),
		PszText: &utf16[0],
		ISubItem: int32(index),
	}
	ret := lv.SendMessage(winapi.LVM_INSERTCOLUMNW, uintptr(index), uintptr(unsafe.Pointer(&col)))
	return int(int32(ret))
}

// InsertItem inserts a new item at index (or appends if index < 0) and returns its index (or -1).
func (lv *ListView) InsertItem(index int, text string) int {
	if index < 0 {
		index = lv.ItemCount()
	}
	utf16 := syscall.StringToUTF16(text)
	item := win.LVITEM{
		Mask:     win.LVIF_TEXT,
		IItem:    int32(index),
		ISubItem: 0,
		PszText:  &utf16[0],
	}
	ret := lv.SendMessage(winapi.LVM_INSERTITEMW, 0, uintptr(unsafe.Pointer(&item)))
	return int(int32(ret))
}

// SetItemText sets the text for an item/subitem.
func (lv *ListView) SetItemText(itemIndex int, subItem int, text string) bool {
	utf16 := syscall.StringToUTF16(text)
	item := win.LVITEM{
		IItem:    int32(itemIndex),
		ISubItem: int32(subItem),
		PszText:  &utf16[0],
	}
	ret := lv.SendMessage(winapi.LVM_SETITEMTEXTW, uintptr(itemIndex), uintptr(unsafe.Pointer(&item)))
	return ret != 0
}

// GetItemText returns the text for an item/subitem.
func (lv *ListView) GetItemText(itemIndex int, subItem int) string {
	buf := make([]uint16, 512)
	item := win.LVITEM{
		IItem:       int32(itemIndex),
		ISubItem:    int32(subItem),
		CchTextMax:  int32(len(buf)),
		PszText:     &buf[0],
	}
	lv.SendMessage(winapi.LVM_GETITEMTEXTW, uintptr(itemIndex), uintptr(unsafe.Pointer(&item)))
	return syscall.UTF16ToString(buf)
}

// ItemCount returns the number of items in the list.
func (lv *ListView) ItemCount() int {
	return int(int32(lv.SendMessage(winapi.LVM_GETITEMCOUNT, 0, 0)))
}

// DeleteAllItems deletes all items.
func (lv *ListView) DeleteAllItems() bool {
	return lv.SendMessage(winapi.LVM_DELETEALLITEMS, 0, 0) != 0
}

// DeleteItem deletes an item at index.
func (lv *ListView) DeleteItem(index int) bool {
	return lv.SendMessage(winapi.LVM_DELETEITEM, uintptr(index), 0) != 0
}

// SelectedIndex returns the first selected item index, or -1 if none.
func (lv *ListView) SelectedIndex() int {
	ret := lv.SendMessage(winapi.LVM_GETNEXTITEM, ^uintptr(0), uintptr(win.LVNI_SELECTED))
	return int(int32(ret))
}

// NewListView creates a new ListView, need bind to Dialog before use.
func NewListView(idd uintptr) *ListView {
	return &ListView{WindowBase: WindowBase{idd: idd}}
}

// BindNewListView creates a new ListView and bind to target dlg.
func BindNewListView(idd uintptr, dlg *Dialog) (*ListView, error) {
	lv := NewListView(idd)
	err := dlg.BindWidgets(lv)
	return lv, err
}
