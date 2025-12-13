package wingui

import (
	"sync/atomic"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"github.com/whtiehack/wingui/winapi"
)

// WindowBase is an interface that provides operations common to all windows.
// if need subclass window ,must set Subclassing to true.
// hwnd,if custom window could set by Init.
// idd Reource Id. if custom window could set zero by Init
type WindowBase struct {
	hwnd win.HWND

	idd uintptr
	// lpPrevWndFunc is a WndProc point if Subclassing set to true.
	lpPrevWndFunc uintptr
	// Subclassing indicate that this window need Subclass,
	// make sure set this flag before bind to Dialog.
	Subclassing bool
}

// AsWindowBase  return a *WindowBase.
func (w *WindowBase) AsWindowBase() *WindowBase {
	return w
}

// Init could init new WindowBase by youself .
func (w *WindowBase) Init(hwnd win.HWND, idd uintptr) {
	w.hwnd = hwnd
	w.idd = idd
}

// Handle get hwnd.
func (w *WindowBase) Handle() win.HWND {
	return w.hwnd
}

// SetWindowText set text
func (w *WindowBase) SetWindowText(title string) {
	// log.Printf("SetCaption hwnd: %v, %s\n", w.hwnd, title)
	//win.SetWindowText(w.hwnd, title)
	win.SendMessage(w.hwnd, win.WM_SETTEXT, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
}

// GetWindowText get text.
func (w *WindowBase) GetWindowText() string {
	textLength := win.SendMessage(w.hwnd, win.WM_GETTEXTLENGTH, 0, 0)
	buf := make([]uint16, textLength+1)
	win.SendMessage(w.hwnd, win.WM_GETTEXT, uintptr(textLength+1), uintptr(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf)
}

// Text alias to GetWindowText
func (w *WindowBase) Text() string {
	return w.GetWindowText()
}

// SetText alias to SetWindowText
func (w *WindowBase) SetText(str string) {
	w.SetWindowText(str)
}

// SetIcon set window icon.
// IconType: 1 - ICON_BIG; 0 - ICON_SMALL
// Icon: Resource Id or Icon Handle
// LoadIcon: If Icon is ResourceId then invoke LoadIcon
func (w *WindowBase) SetIcon(iconType int, icon uintptr, loadIcon bool) {
	if iconType > 1 {
		panic("IconType is invalid")
	}
	if loadIcon {
		icon = uintptr(win.LoadIcon(hInstance, win.MAKEINTRESOURCE(icon)))
	}
	win.SendMessage(w.hwnd, win.WM_SETICON, uintptr(iconType), icon)
}

// Show set window show.
func (w *WindowBase) Show() {
	win.ShowWindow(w.hwnd, win.SW_SHOW)
}

// Hide set window hide.
func (w *WindowBase) Hide() {
	win.ShowWindow(w.hwnd, win.SW_HIDE)
}

// ShowMinimized show minimized btn.
func (w *WindowBase) ShowMinimized() {
	win.ShowWindow(w.hwnd, win.SW_MINIMIZE)
}

// ShowMaximized show maximized btn.
func (w *WindowBase) ShowMaximized() {
	win.ShowWindow(w.hwnd, win.SW_MAXIMIZE)
}

// ShowFullScreen ShowFullScreen
func (w *WindowBase) ShowFullScreen() {
	win.ShowWindow(w.hwnd, win.SW_SHOWMAXIMIZED)
}

// ShowNormal ShowNormal
func (w *WindowBase) ShowNormal() {
	win.ShowWindow(w.hwnd, win.SW_SHOWNORMAL)
}

// IsEnabled check windows is enabled.
func (w *WindowBase) IsEnabled() bool {
	return win.IsWindowEnabled(w.hwnd)
}

// IsVisible check window is visible.
func (w *WindowBase) IsVisible() bool {
	return win.IsWindowVisible(w.hwnd)
}

// SetVisible set window visible status.
func (w *WindowBase) SetVisible(value bool) {
	var cmd int32
	if value {
		cmd = win.SW_SHOW
	} else {
		cmd = win.SW_HIDE
	}
	win.ShowWindow(w.hwnd, cmd)
}

// SetEnabled set window enable status.
func (w *WindowBase) SetEnabled(b bool) {
	win.EnableWindow(w.hwnd, b)
}

// SetDisabled reverse of SetEnabled
func (w *WindowBase) SetDisabled(disable bool) {
	win.EnableWindow(w.hwnd, !disable)
}

// Close close window.
func (w *WindowBase) Close() {
	win.SendMessage(w.hwnd, win.WM_CLOSE, 0, 0)
}

// SetFocus set focus.
func (w *WindowBase) SetFocus() {
	win.SetFocus(w.hwnd)
}

// GetWindowRect get window rect
func (w *WindowBase) GetWindowRect() win.RECT {
	var rect win.RECT
	win.GetWindowRect(w.hwnd, &rect)
	return rect
}

// BoundsPixels returns the outer bounding box Rectangle of the *WindowBase, including
// decorations.
// The coordinates are relative to the screen.
func (w *WindowBase) BoundsPixels() Rectangle {
	var r win.RECT

	if !win.GetWindowRect(w.hwnd, &r) {
		return Rectangle{}
	}
	return Rectangle{
		int(r.Left),
		int(r.Top),
		int(r.Right - r.Left),
		int(r.Bottom - r.Top),
	}
}

// SetBounds set window rect.
func (w *WindowBase) SetBounds(value Rectangle) {
	win.MoveWindow(w.hwnd, int32(value.X), int32(value.Y), int32(value.Width), int32(value.Height), true)
}

// WndProc process window message.
func (w *WindowBase) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	// log.Println("WidgetBase.WndProc")
	if w.lpPrevWndFunc != 0 {
		return win.CallWindowProc(w.lpPrevWndFunc, w.hwnd, msg, wParam, lParam)
	}
	return uintptr(0)
}

// SendMessage sends a message to the window and returns the result.
func (w *WindowBase) SendMessage(msg uint32, wParam, lParam uintptr) uintptr {
	if w.hwnd == 0 {
		return 0
	}
	// Calling SendMessage across threads can deadlock (common with controls + modal loops).
	// If we don't know the UI thread yet (uiThreadID==0) or we're not on it, use SendMessageTimeout to avoid hanging the process.
	uiTID := atomic.LoadUint32(&uiThreadID)
	if uiTID == 0 || win.GetCurrentThreadId() != uiTID {
		if ret, ok := winapi.SendMessageTimeout(w.hwnd, msg, wParam, lParam, winapi.SMTO_ABORTIFHUNG|winapi.SMTO_BLOCK, 2000); ok {
			return ret
		}
		return 0
	}
	return win.SendMessage(w.hwnd, msg, wParam, lParam)
}
