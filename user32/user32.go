package user32

import (
	"github.com/lxn/win"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

var libuser32 *windows.LazyDLL
var findWindowEx *windows.LazyProc

func init() {
	libuser32 = windows.NewLazySystemDLL("user32.dll")
	findWindowEx = libuser32.NewProc("FindWindowExW")
}
func FindWindowEx(hWndParent win.HWND, hWndChildAfter win.HWND, lpClassName, lpWindowName *uint16) win.HWND {
	ret, _, _ := syscall.Syscall6(findWindowEx.Addr(), 4,
		uintptr(hWndParent),
		uintptr(hWndChildAfter),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		0, 0,
	)
	return win.HWND(ret)
}
