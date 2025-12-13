package winapi

import (
	"unsafe"

	"github.com/lxn/win"
)

var (
	messageBoxW = libuser32.NewProc("MessageBoxW")
)

// MessageBox wraps user32!MessageBoxW.
func MessageBox(hWnd win.HWND, lpText, lpCaption *uint16, uType uint32) int {
	ret, _, _ := messageBoxW.Call(
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		uintptr(uType),
	)
	return int(ret)
}
