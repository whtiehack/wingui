package wingui

// https://docs.microsoft.com/zh-cn/windows/win32/controls/buttons?view=vs-2017

import (
	"github.com/lxn/win"
	"log"
	"syscall"
	"unsafe"
)

// Button a widget for Dialog.  base of all button type.
type Button struct {
	WindowBase
	OnClicked func()
}

// WndProc Button window WndProc.
func (b *Button) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	// log.Println("btn wnd proc", b.hwnd, msg, wParam, lParam)
	switch msg {
	case win.WM_COMMAND:
		switch win.HIWORD(uint32(wParam)) {
		case win.BN_CLICKED:
			if b.OnClicked != nil && lParam == uintptr(b.hwnd) {
				b.OnClicked()
			}
		}
	}
	return b.AsWindowBase().WndProc(msg, wParam, lParam)
}

// GetNote gets the text of the note associated with a command link button.
// CommandLink button
func (b *Button) GetNote() (str string) {
	textLength := b.SendMessage(win.BCM_GETNOTELENGTH, 0, 0)
	buf := make([]uint16, textLength+1)
	ret := b.SendMessage(win.BCM_GETNOTE, textLength+1, uintptr(unsafe.Pointer(&buf[0])))
	if int(ret) == 0 {
		log.Println("GetNote error:", b, "len:", textLength)
	} else {
		str = syscall.UTF16ToString(buf)
	}
	return
}

// SetNote sets the text of the note associated with a command link button.
func (b *Button) SetNote(note string) {
	b.SendMessage(win.BCM_SETNOTE, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(note))))
}

// SetDropDownState sets the drop down state for a button with style TBSTYLE_DROPDOWN.
// state is a BOOL that is TRUE for state of BST_DROPDOWNPUSHED, or FALSE otherwise.
func (b *Button) SetDropDownState(state bool) {
	wParam := 0
	if state {
		wParam = 1
	}
	b.SendMessage(win.BCM_SETDROPDOWNSTATE, uintptr(wParam), 0)
}

// SetShield sets the elevation required state for a specified button or command link to display an elevated icon.
// state is a BOOL that is TRUE to draw an elevated icon, or FALSE otherwise.
func (b *Button) SetShield(state bool) {
	wParam := 0
	if state {
		wParam = 1
	}
	b.SendMessage(win.BCM_SETSHIELD, uintptr(wParam), 0)
}

// Click simulates the user clicking a button. This message causes the button to receive the WM_LBUTTONDOWN and
// WM_LBUTTONUP messages, and the button's parent window to receive a BN_CLICKED notification code.
func (b *Button) Click() {
	b.SendMessage(win.BM_CLICK, 0, 0)
}

// NewButton create a new Button,need bind to Dialog before use.
func NewButton(idd uintptr) *Button {
	return &Button{WindowBase{idd: idd}, nil}
}

// BindNewButton create a new Button and bind to target dlg.
func BindNewButton(idd uintptr, dlg *Dialog) (*Button, error) {
	b := NewButton(idd)
	err := dlg.BindWidgets(b)
	return b, err
}
