package winapi

import (
	"unsafe"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

var libuser32 *windows.LazyDLL
var findWindowEx *windows.LazyProc
var getNextWindow *windows.LazyProc

func init() {
	libuser32 = windows.NewLazySystemDLL("user32.dll")
	findWindowEx = libuser32.NewProc("FindWindowExW")
	getNextWindow = libuser32.NewProc("GetNextWindow ")
}

//FindWindowEx user32 API FindWindowEx
func FindWindowEx(hWndParent win.HWND, hWndChildAfter win.HWND, lpClassName, lpWindowName *uint16) win.HWND {
	ret, _, _ := findWindowEx.Call(
		uintptr(hWndParent),
		uintptr(hWndChildAfter),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
	)
	return win.HWND(ret)
}

//GetNextWindow user32 API GetNextWindow
func GetNextWindow(hWnd win.HWND, wCmd uintptr) win.HWND {
	ret, _, _ := getNextWindow.Call(uintptr(hWnd), wCmd)
	return win.HWND(ret)
}
