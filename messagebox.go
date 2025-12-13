package wingui

import (
	"syscall"

	"github.com/lxn/win"
	"github.com/whtiehack/wingui/winapi"
)

// MessageBox shows a native Win32 message box.
func MessageBox(owner win.HWND, text, caption string, flags uint32) int {
	return winapi.MessageBox(owner, syscall.StringToUTF16Ptr(text), syscall.StringToUTF16Ptr(caption), flags)
}

// MessageBox shows a native Win32 message box using the window as the owner.
func (w *WindowBase) MessageBox(text, caption string, flags uint32) int {
	if w == nil {
		return MessageBox(0, text, caption, flags)
	}
	return MessageBox(w.Handle(), text, caption, flags)
}
