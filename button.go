package wingui

// https:// docs.microsoft.com/zh-cn/windows/win32/controls/buttons?view=vs-2017

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

// GetCheck gets the check state of a radio button or check box.
// 
// checked return 1, unchecked return 0
// 
// if Button has BS_3STATE or BS_AUTO3STATE style, may be return  2
// 
// BST_CHECKED 1
// 
// BST_INDETERMINATE 2
// 
// BST_UNCHECKED 0
func (b *Button) GetCheck() int {
	return int(b.SendMessage(win.BM_GETCHECK, 0, 0))
}

// GetState retrieves the state of a button or check box.
// 
// The return value specifies the current state of the button. It is a combination of the following values.
// 
// Return code	Description
// 
// BST_CHECKED
// The button is checked.
// 
// BST_DROPDOWNPUSHED
// Windows Vista. The button is in the drop-down state. Applies only if the button has the TBSTYLE_DROPDOWN style.
// 
// BST_FOCUS
// The button has the keyboard focus.
// 
// BST_HOT
// The button is hot; that is, the mouse is hovering over it.
// 
// BST_INDETERMINATE
// The state of the button is indeterminate. Applies only if the button has the BS_3STATE or BS_AUTO3STATE style.
// 
// BST_PUSHED
// The button is being shown in the pushed state.
// 
// 
// BST_UNCHECKED
// No special state. Equivalent to zero.
func (b *Button) GetState() int {
	return int(b.SendMessage(win.BM_GETSTATE, 0, 0))
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
