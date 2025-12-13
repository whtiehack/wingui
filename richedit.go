package wingui

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

// RichEdit is a wrapper for the Win32 Rich Edit control.
// The class is provided by Msftedit.dll (RICHEDIT50W) or Riched20.dll (RichEdit20W).
type RichEdit struct {
	WindowBase
	OnChanged func()
	lastLen   int
}

func (re *RichEdit) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	maybeEmitChange := func() {
		if re.OnChanged == nil || re.hwnd == 0 {
			return
		}
		cur := re.TextLength()
		if re.lastLen != cur {
			re.lastLen = cur
			re.OnChanged()
		}
	}

	switch msg {
	case win.WM_COMMAND:
		switch win.HIWORD(uint32(wParam)) {
		case win.EN_CHANGE:
			if lParam == uintptr(re.hwnd) {
				maybeEmitChange()
			}
		case win.EN_UPDATE:
			if lParam == uintptr(re.hwnd) {
				maybeEmitChange()
			}
		}
	case win.WM_CHAR, win.WM_PASTE, win.WM_CUT, win.WM_CLEAR, win.WM_UNDO:
		ret := re.AsWindowBase().WndProc(msg, wParam, lParam)
		maybeEmitChange()
		return ret
	}
	return re.AsWindowBase().WndProc(msg, wParam, lParam)
}

// TextSelection returns the current selection range (character index).
func (re *RichEdit) TextSelection() (start, end int) {
	re.SendMessage(win.EM_GETSEL, uintptr(unsafe.Pointer(&start)), uintptr(unsafe.Pointer(&end)))
	return
}

// SetTextSelection sets the current selection range (character index).
func (re *RichEdit) SetTextSelection(start, end int) {
	re.SendMessage(win.EM_SETSEL, uintptr(start), uintptr(end))
}

// ReplaceSelectedText replaces the current selection with text.
func (re *RichEdit) ReplaceSelectedText(text string, canUndo bool) {
	re.SendMessage(win.EM_REPLACESEL,
		uintptr(win.BoolToBOOL(canUndo)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
	)
}

// AppendText appends text to the end of the control.
func (re *RichEdit) AppendText(value string) {
	s, e := re.TextSelection()
	l := re.TextLength()
	re.SetTextSelection(l, l)
	re.ReplaceSelectedText(value, false)
	re.SetTextSelection(s, e)
}

// TextLength returns current text length.
func (re *RichEdit) TextLength() int {
	return int(re.SendMessage(win.WM_GETTEXTLENGTH, 0, 0))
}

// SetReadOnly toggles read-only mode.
func (re *RichEdit) SetReadOnly(readOnly bool) {
	re.SendMessage(win.EM_SETREADONLY, uintptr(win.BoolToBOOL(readOnly)), 0)
}

// NewRichEdit creates a new RichEdit, need bind to Dialog before use.
func NewRichEdit(idd uintptr) *RichEdit {
	return &RichEdit{
		WindowBase: WindowBase{idd: idd, Subclassing: true},
		lastLen:    -1,
	}
}

// BindNewRichEdit creates a new RichEdit and bind to target dlg.
func BindNewRichEdit(idd uintptr, dlg *Dialog) (*RichEdit, error) {
	re := NewRichEdit(idd)
	err := dlg.BindWidgets(re)
	return re, err
}
