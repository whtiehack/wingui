package winapi

import (
	"unsafe"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

// Button message constants
const (
	// #if(WINVER >= 0x0600)
	//#define BM_SETDONTCLICK    0x00F8
	//#endif /* WINVER >= 0x0600 */
	BM_SETDONTCLICK = 0x00F8
)

var libuser32 *windows.LazyDLL
var findWindowEx *windows.LazyProc
var getNextWindow *windows.LazyProc
var flashWindow *windows.LazyProc

func init() {
	libuser32 = windows.NewLazySystemDLL("user32.dll")
	findWindowEx = libuser32.NewProc("FindWindowExW")
	getNextWindow = libuser32.NewProc("GetNextWindow")
	flashWindow = libuser32.NewProc("FlashWindow")
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

func FlashWindow(hWnd win.HWND, bInvert int) int {
	ret, _, _ := flashWindow.Call(uintptr(hWnd), uintptr(bInvert))
	return int(ret)
}
