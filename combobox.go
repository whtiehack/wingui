// https://docs.microsoft.com/zh-cn/windows/win32/controls/combo-boxes?view=vs-2017

package wingui

import (
	"errors"
	"github.com/lxn/win"
	"log"
	"syscall"
	"unsafe"
)

// ComboBox a ComboBox widget for Dialog.
type ComboBox struct {
	WindowBase

	// TODO: notify method
}

// WndProc ComboBox window WndProc.
func (cb *ComboBox) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	}
	return cb.AsWindowBase().WndProc(msg, wParam, lParam)
}

// CurrentIndex get currentIndex.
func (cb *ComboBox) CurrentIndex() int {
	return int(int32(cb.SendMessage(win.CB_GETCURSEL, 0, 0)))
}

// SetCurrentIndex set current index
func (cb *ComboBox) SetCurrentIndex(value int) error {
	index := int(int32(cb.SendMessage(win.CB_SETCURSEL, uintptr(value), 0)))
	if index != value {
		return errors.New("invalid index")
	}
	return nil
}

// GetCount get items count of combobox.
func (cb *ComboBox) GetCount() int {
	return int(int32(cb.SendMessage(win.CB_GETCOUNT, 0, 0)))
}

// SetItemData sends a CB_SETITEMDATA message to set the value associated with the specified item in a combo box.
func (cb *ComboBox) SetItemData(index int, data uintptr) error {
	ret := cb.SendMessage(win.CB_SETITEMDATA, uintptr(index), data)
	if ret == win.CB_ERR {
		return errors.New("set item data error")
	}
	return nil
}

// GetItemData  sends a CB_GETITEMDATA message to a combo box to retrieve the application-supplied value associated
// with the specified item in the combo box.
func (cb *ComboBox) GetItemData(index int) (data uintptr, err error) {
	ret := cb.SendMessage(win.CB_GETITEMDATA, uintptr(index), 0)
	if ret == win.CB_ERR {
		err = errors.New("GetItemData error")
	} else {
		data = ret
	}
	return
}

// ResetContent removes all items from the list box and edit control of a combo box.
func (cb *ComboBox) ResetContent() int {
	return int(int32(cb.SendMessage(win.CB_RESETCONTENT, 0, 0)))
}

// AddString adds a string to the list box of a combo box. If the combo box does not have the CBS_SORT style,
// the string is added to the end of the list. Otherwise, the string is inserted into the list, and the list is sorted.
// The return value idx is the zero-based index to the string in the list box of the combo box.
func (cb *ComboBox) AddString(str string) (idx int, err error) {
	ret := cb.SendMessage(win.CB_ADDSTRING, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str))))
	if ret < 0 {
		err = errors.New("add string to combobox error")
	} else {
		idx = int(ret)
	}
	return
}

// DeleteString deletes a string in the list box of a combo box.
// The return value leftCount is a count of the strings remaining in the list
func (cb *ComboBox) DeleteString(idx int) (leftCount int, err error) {
	ret := cb.SendMessage(win.CB_DELETESTRING, 0, uintptr(idx))
	if ret < 0 {
		err = errors.New("DeleteString combobox error")
	} else {
		leftCount = int(ret)
	}
	return
}

//InsertString  Inserts a string or item data into the list of a combo box. Unlike the CB_ADDSTRING message,
// the CB_INSERTSTRING message does not cause a list with the CBS_SORT style to be sorted.
func (cb *ComboBox) InsertString(idx int, str string) (err error) {
	ret := cb.SendMessage(win.CB_INSERTSTRING, uintptr(idx), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str))))
	if ret != uintptr(idx) {
		err = errors.New("InsertString to combobox error")
	}
	return
}

// GetLbText
func (cb *ComboBox) GetLbText(idx int) (str string) {
	textLength := cb.SendMessage(win.CB_GETLBTEXTLEN, uintptr(idx), 0)
	buf := make([]uint16, textLength+1)
	ret := cb.SendMessage(win.CB_GETLBTEXT, uintptr(idx), uintptr(unsafe.Pointer(&buf[0])))
	if ret < 0 {
		log.Println("GetLbText error:", cb, "idx:", idx)
	} else {
		str = syscall.UTF16ToString(buf)
	}
	return
}

// NewComboBox create a new ComboBox,need bind to Dialog before use.
func NewComboBox(idd uintptr) *ComboBox {
	return &ComboBox{
		WindowBase: WindowBase{idd: idd},
	}
}

// BindNewStatic create a new Image and bind to target dlg.
func BindNewComboBox(idd uintptr, dlg *Dialog) (*ComboBox, error) {
	b := NewComboBox(idd)
	err := dlg.BindWidgets(b)
	return b, err
}
