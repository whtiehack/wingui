package wingui

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

// Edit a widget for Dialog.
type Edit struct {
	WindowBase
}

// WndProc Edit Window WndProc.
func (e *Edit) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	// log.Println("btn wnd proc", b.hwnd, msg, wParam, lParam)
	e.AsWindowBase().WndProc(msg, wParam, lParam)
	return 0
}

// TextSelection get Edit selection.
func (e *Edit) TextSelection() (start, end int) {
	e.SendMessage(win.EM_GETSEL, uintptr(unsafe.Pointer(&start)), uintptr(unsafe.Pointer(&end)))
	return
}

// SetTextSelection set Edit selection
func (e *Edit) SetTextSelection(start, end int) {
	e.SendMessage(win.EM_SETSEL, uintptr(start), uintptr(end))
}

// ReplaceSelectedText replace Text.
func (e *Edit) ReplaceSelectedText(text string, canUndo bool) {
	e.SendMessage(win.EM_REPLACESEL,
		uintptr(win.BoolToBOOL(canUndo)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))))
}

// AppendText append text to Edit
func (e *Edit) AppendText(value string) {
	s, end := e.TextSelection()
	l := e.TextLength()
	e.SetTextSelection(l, l)
	e.ReplaceSelectedText(value, false)
	e.SetTextSelection(s, end)
}

// TextLength get Edit text length.
func (e *Edit) TextLength() int {
	return int(e.SendMessage(win.WM_GETTEXTLENGTH, 0, 0))
}

// NewEdit create a new Edit ,need bind to Dialog for use.
func NewEdit(idd uintptr) *Edit {
	return &Edit{WindowBase{idd: idd}}
}
