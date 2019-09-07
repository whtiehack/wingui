package wingui

import (
	"github.com/lxn/win"
	"log"
	"syscall"
	"unsafe"
)

// WindowBase
type WindowBase struct {
	hwnd   win.HWND
	parent win.HWND
	idd    uintptr
}

func (w *WindowBase) AsWindowBase() *WindowBase {
	return w
}

func (w *WindowBase) Handle() win.HWND {
	return w.hwnd
}

func (w *WindowBase) SetWindowText(title string) {
	log.Printf("SetCaption hwnd: %v, %s\n", w.hwnd, title)
	//win.SetWindowText(w.hwnd, title)
	win.SendMessage(w.hwnd, win.WM_SETTEXT, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
}

func (w *WindowBase) GetWindowText() string {
	textLength := win.SendMessage(w.hwnd, win.WM_GETTEXTLENGTH, 0, 0)
	buf := make([]uint16, textLength+1)
	win.SendMessage(w.hwnd, win.WM_GETTEXT, uintptr(textLength+1), uintptr(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf)
}

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

func (w *WindowBase) Show() {
	win.ShowWindow(w.hwnd, win.SW_SHOW)
}

func (w *WindowBase) Hide() {
	win.ShowWindow(w.hwnd, win.SW_HIDE)
}

func (w *WindowBase) ShowMinimized() {
	win.ShowWindow(w.hwnd, win.SW_MINIMIZE)
}

func (w *WindowBase) ShowMaximized() {
	win.ShowWindow(w.hwnd, win.SW_MAXIMIZE)
}

func (w *WindowBase) ShowFullScreen() {
	win.ShowWindow(w.hwnd, win.SW_SHOWMAXIMIZED)
}

func (w *WindowBase) ShowNormal() {
	win.ShowWindow(w.hwnd, win.SW_SHOWNORMAL)
}

func (w *WindowBase) IsEnabled() bool {
	return win.IsWindowEnabled(w.hwnd)
}

func (w *WindowBase) IsVisible() bool {
	return win.IsWindowVisible(w.hwnd)
}

func (w *WindowBase) SetVisible(value bool) {
	var cmd int32
	if value {
		cmd = win.SW_SHOW
	} else {
		cmd = win.SW_HIDE
	}
	win.ShowWindow(w.hwnd, cmd)
}

func (w *WindowBase) SetEnabled(b bool) {
	win.EnableWindow(w.hwnd, b)
}

func (w *WindowBase) SetDisabled(disable bool) {
	win.EnableWindow(w.hwnd, !disable)
}

func (w *WindowBase) Close() {
	win.SendMessage(w.hwnd, win.WM_CLOSE, 0, 0)
}

func (w *WindowBase) SetFocus() {
	win.SetFocus(w.hwnd)
}

func (w *WindowBase) GetWindowRect() win.RECT {
	var rect win.RECT
	win.GetWindowRect(w.hwnd, &rect)
	return rect
}

// BoundsPixels returns the outer bounding box Rectangle of the *WindowBase, including
// decorations.
//
// The coordinates are relative to the screen.
func (wb *WindowBase) BoundsPixels() Rectangle {
	var r win.RECT

	if !win.GetWindowRect(wb.hwnd, &r) {
		return Rectangle{}
	}
	return Rectangle{
		int(r.Left),
		int(r.Top),
		int(r.Right - r.Left),
		int(r.Bottom - r.Top),
	}
}

func (w *WindowBase) SetBounds(value Rectangle) {
	win.MoveWindow(w.hwnd, int32(value.X), int32(value.Y), int32(value.Width), int32(value.Height), true)
}

func (w *WindowBase) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	log.Println("WidgetBase.WndProc")
	return uintptr(0)
}
