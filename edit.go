package wingui

import (
	"github.com/lxn/win"
	"syscall"
	"unsafe"
)

type Edit struct {
	WindowBase
}

func (e *Edit) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	// log.Println("btn wnd proc", b.hwnd, msg, wParam, lParam)
	e.AsWindowBase().WndProc(msg, wParam, lParam)
	return 0
}

func (te *Edit) TextSelection() (start, end int) {
	te.SendMessage(win.EM_GETSEL, uintptr(unsafe.Pointer(&start)), uintptr(unsafe.Pointer(&end)))
	return
}

func (te *Edit) SetTextSelection(start, end int) {
	te.SendMessage(win.EM_SETSEL, uintptr(start), uintptr(end))
}

func (te *Edit) ReplaceSelectedText(text string, canUndo bool) {
	te.SendMessage(win.EM_REPLACESEL,
		uintptr(win.BoolToBOOL(canUndo)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))))
}

func (te *Edit) AppendText(value string) {
	s, e := te.TextSelection()
	l := te.TextLength()
	te.SetTextSelection(l, l)
	te.ReplaceSelectedText(value, false)
	te.SetTextSelection(s, e)
}
func (te *Edit) TextLength() int {
	return int(te.SendMessage(win.WM_GETTEXTLENGTH, 0, 0))
}
