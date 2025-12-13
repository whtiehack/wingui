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

// SendMessageTimeout flags.
const (
	SMTO_NORMAL             = 0x0000
	SMTO_BLOCK              = 0x0001
	SMTO_ABORTIFHUNG        = 0x0002
	SMTO_NOTIMEOUTIFNOTHUNG = 0x0008
)

var (
	libuser32          = windows.NewLazySystemDLL("user32.dll")
	findWindowEx       = libuser32.NewProc("FindWindowExW")
	getNextWindow      = libuser32.NewProc("GetNextWindow ")
	flashWindow        = libuser32.NewProc("FlashWindow")
	switchToWindow     = libuser32.NewProc("SwitchToThisWindow")
	sendMessageTimeout = libuser32.NewProc("SendMessageTimeoutW")
)

// FindWindowEx user32 API FindWindowEx
func FindWindowEx(hWndParent win.HWND, hWndChildAfter win.HWND, lpClassName, lpWindowName *uint16) win.HWND {
	ret, _, _ := findWindowEx.Call(
		uintptr(hWndParent),
		uintptr(hWndChildAfter),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
	)
	return win.HWND(ret)
}

// GetNextWindow user32 API GetNextWindow
func GetNextWindow(hWnd win.HWND, wCmd uintptr) win.HWND {
	ret, _, _ := getNextWindow.Call(uintptr(hWnd), wCmd)
	return win.HWND(ret)
}

func FlashWindow(hWnd win.HWND, bInvert int) int {
	ret, _, _ := flashWindow.Call(uintptr(hWnd), uintptr(bInvert))
	return int(ret)
}

func SwitchToWindow(hWnd win.HWND, bInvert int) int {
	ret, _, _ := switchToWindow.Call(uintptr(hWnd), uintptr(bInvert))
	return int(ret)
}

// SendMessageTimeout calls SendMessageTimeoutW and returns (result, ok).
func SendMessageTimeout(hWnd win.HWND, msg uint32, wParam, lParam uintptr, flags uint32, timeoutMs uint32) (uintptr, bool) {
	var result uintptr
	ret, _, _ := sendMessageTimeout.Call(
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		uintptr(flags),
		uintptr(timeoutMs),
		uintptr(unsafe.Pointer(&result)),
	)
	return result, ret != 0
}
