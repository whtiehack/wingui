// https://docs.microsoft.com/zh-cn/windows/win32/controls/list-boxes?view=vs-2017

package wingui

import (
	"errors"
	"github.com/lxn/win"
	"syscall"
	"unsafe"
)

// ListBox a ListBox widget for Dialog.
type ListBox struct {
	WindowBase
	// TODO: notify method
}

// WndProc ListBox window WndProc.
func (lb *ListBox) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_COMMAND:
		cmdCode := win.HIWORD(uint32(wParam))
		switch cmdCode {
		}
	}
	return lb.AsWindowBase().WndProc(msg, wParam, lParam)
}

// CurrentIndex gets the index of the currently selected item, if any, in a single-selection list box.
//  If there is no selection, the return value is -1.
func (lb *ListBox) CurrentIndex() int {
	return int(int32(lb.SendMessage(win.LB_GETCURSEL, 0, 0)))
}

// GetSel gets the selection state of an item. If selected , return true;otherwise return false
func (lb *ListBox) GetSel(idx int) bool {
	ret := lb.SendMessage(win.LB_GETSEL, uintptr(idx), 0)
	return int(ret) > 0
}

// SelectedIndexes gets the current selected items index.
// The return value is the array of index selected.
// If the list box is a single-selection list box, the return value is nil.
func (lb *ListBox) SelectedIndexes() []int {
	count := lb.SendMessage(win.LB_GETSELCOUNT, 0, 0)
	if count < 1 {
		return nil
	}
	index32 := make([]int32, count)
	if n := int(int32(lb.SendMessage(win.LB_GETSELITEMS, count, uintptr(unsafe.Pointer(&index32[0]))))); n == win.LB_ERR {
		return nil
	} else {
		indexes := make([]int, n)
		for i := 0; i < n; i++ {
			indexes[i] = int(index32[i])
		}
		return indexes
	}
}

// GetCount gets the number of items in a list box.
func (lb *ListBox) GetCount() int {
	return int(int32(lb.SendMessage(win.LB_GETCOUNT, 0, 0)))
}

// GetItemData gets the application-defined value associated with the specified list box item.
func (lb *ListBox) GetItemData(index int) (data uintptr, err error) {
	ret := lb.SendMessage(win.LB_GETITEMDATA, uintptr(index), 0)
	if int(ret) < 0 {
		err = errors.New("GetItemData error")
	} else {
		data = ret
	}
	return
}

// ResetContent removes all items from a list box.
func (lb *ListBox) ResetContent() int {
	return int(int32(lb.SendMessage(win.LB_RESETCONTENT, 0, 0)))
}

// AddString Adds a string to a list box. If the list box does not have the LBS_SORT style,
// the string is added to the end of the list. Otherwise, the string is inserted into the list and the list is sorted.
func (lb *ListBox) AddString(str string) (idx int, err error) {
	ret := lb.SendMessage(win.LB_ADDSTRING, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str))))
	if int(ret) < 0 {
		err = errors.New("add string to ListBox error")
	} else {
		idx = int(ret)
	}
	return
}

// DeleteString deletes a string in a list box.
// The return value leftCount is a count of the strings remaining in the list
func (lb *ListBox) DeleteString(idx int) (leftCount int, err error) {
	ret := lb.SendMessage(win.LB_DELETESTRING, 0, uintptr(idx))
	if int(ret) < 0 {
		err = errors.New("DeleteString ListBox error")
	} else {
		leftCount = int(ret)
	}
	return
}

// NewListBox create a new ListBox,need bind to Dialog before use.
func NewListBox(idd uintptr) *ListBox {
	return &ListBox{
		WindowBase: WindowBase{idd: idd},
	}
}

// BindNewListBox create a new ListBox and bind to target dlg.
func BindNewListBox(idd uintptr, dlg *Dialog) (*ListBox, error) {
	b := NewListBox(idd)
	err := dlg.BindWidgets(b)
	return b, err
}
