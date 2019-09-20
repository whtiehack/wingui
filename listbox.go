// https://docs.microsoft.com/zh-cn/windows/win32/controls/list-boxes?view=vs-2017

package wingui

import (
	"errors"
	"github.com/lxn/win"
	"log"
	"syscall"
	"unsafe"
)

// ListBox a ListBox widget for Dialog.
type ListBox struct {
	WindowBase
	// TOxDO: notify method
	OnDoubleClick func()
	OnSelChange   func()
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

// WndProc ListBox window WndProc.
func (lb *ListBox) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_COMMAND:
		cmdCode := win.HIWORD(uint32(wParam))
		switch cmdCode {
		case win.LBN_DBLCLK:
			if lb.OnDoubleClick != nil {
				lb.OnDoubleClick()
			}
		case win.LBN_SELCHANGE:
			if lb.OnSelChange != nil {
				lb.OnSelChange()
			}
		}
	}
	return lb.AsWindowBase().WndProc(msg, wParam, lParam)
}

// GetCurSel gets the index of the currently selected item, if any, in a single-selection list box.
//  If there is no selection, the return value is -1.
func (lb *ListBox) GetCurSel() int {
	return int(int32(lb.SendMessage(win.LB_GETCURSEL, 0, 0)))
}

// SetCurSel selects a string and scrolls it into view, if necessary.
// When the new string is selected, the list box removes the highlight from the previously selected string.
// Use this message only with single-selection list boxes.
// You cannot use it to set or remove a selection in a multiple-selection list box.
func (lb *ListBox) SetCurSel(idx int) error {
	index := int(int32(lb.SendMessage(win.LB_SETCURSEL, uintptr(idx), 0)))
	if index != idx {
		return errors.New("invalid index")
	}
	return nil
}

// SetSel selects an item in a multiple-selection list box and, if necessary, scrolls the item into view.
// Use this message only with multiple-selection list boxes.
//  If idx parameter is -1, the selection is added to or removed from all items,
//  depending on the value of wParam, and no scrolling occurs.
func (lb *ListBox) SetSel(idx int, sel bool) error {
	ret := lb.SendMessage(win.LB_SETSEL, uintptr(win.BoolToBOOL(sel)), uintptr(idx))
	if int(ret) == win.LB_ERR {
		return errors.New("ListBox setsel error")
	}
	return nil
}

// GetSel gets the selection state of an item. If selected , return true;otherwise return false
func (lb *ListBox) GetSel(idx int) bool {
	ret := lb.SendMessage(win.LB_GETSEL, uintptr(idx), 0)
	return int(ret) > 0
}

// GetSelectedIndexes gets the current selected items index.
// The return value is the array of index selected.
// If the list box is a single-selection list box, the return value is nil.
func (lb *ListBox) GetSelectedIndexes() []int {
	count := lb.SendMessage(win.LB_GETSELCOUNT, 0, 0)
	if int(count) < 1 {
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

// SetItemData sets a value associated with the specified item in a list box.
func (lb *ListBox) SetItemData(index int, data uintptr) error {
	ret := lb.SendMessage(win.LB_SETITEMDATA, uintptr(index), data)
	if int(ret) == win.LB_ERR {
		return errors.New("set item data error")
	}
	return nil
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
	ret := lb.SendMessage(win.LB_DELETESTRING, uintptr(idx), 0)
	if int(ret) < 0 {
		err = errors.New("DeleteString ListBox error")
	} else {
		leftCount = int(ret)
	}
	return
}

// GetText gets a string from a list box.
func (lb *ListBox) GetText(idx int) (str string) {
	textLength := lb.SendMessage(win.LB_GETTEXTLEN, uintptr(idx), 0)
	buf := make([]uint16, textLength+1)
	ret := lb.SendMessage(win.LB_GETTEXT, uintptr(idx), uintptr(unsafe.Pointer(&buf[0])))
	if int(ret) < 0 {
		log.Println("GetText error:", lb, "idx:", idx)
	} else {
		str = syscall.UTF16ToString(buf)
	}
	return
}

// InsertString inserts a string or item data into a list box. Unlike the LB_ADDSTRING message,
// the LB_INSERTSTRING message does not cause a list with the LBS_SORT style to be sorted.
func (lb *ListBox) InsertString(idx int, str string) (err error) {
	ret := lb.SendMessage(win.LB_INSERTSTRING, uintptr(idx), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str))))
	if ret != uintptr(idx) {
		err = errors.New("InsertString to ListBox error")
	}
	return
}

// SelectString searches a list box for an item that begins with the characters in a specified string. If a matching item is found, the item is selected.
func (lb *ListBox) SelectString(str string, startIdx int) int {
	ret := int(lb.SendMessage(win.LB_SELECTSTRING, uintptr(startIdx), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))))
	if ret < 0 {
		ret = -1
	}
	return ret
}
