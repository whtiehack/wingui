package wingui

//	https://	docs.microsoft.com/zh-cn/windows/win32/controls/buttons?view=vs-2017

import (
	"github.com/lxn/win"
	"github.com/whtiehack/wingui/winapi"
	"log"
	"syscall"
	"unsafe"
)

//	Button a widget for Dialog.  base of all button type.
type Button struct {
	WindowBase
	OnClicked       func()
	OnDoubleClicked func()
}

//	WndProc Button window WndProc.
func (b *Button) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	//	log.Println("btn wnd proc", b.hwnd, msg, wParam, lParam)
	switch msg {
	case win.WM_COMMAND:
		switch win.HIWORD(uint32(wParam)) {
		case win.BN_CLICKED:
			if b.OnClicked != nil && lParam == uintptr(b.hwnd) {
				b.OnClicked()
			}
		case win.BN_DBLCLK:
			if b.OnDoubleClicked != nil && lParam == uintptr(b.hwnd) {
				b.OnDoubleClicked()
			}
		}
	}
	return b.AsWindowBase().WndProc(msg, wParam, lParam)
}

//	GetNote gets the text of the note associated with a command link button.
//	CommandLink button
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

//	SetNote sets the text of the note associated with a command link button.
func (b *Button) SetNote(note string) {
	b.SendMessage(win.BCM_SETNOTE, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(note))))
}

//	SetDropDownState sets the drop down state for a button with style TBSTYLE_DROPDOWN.
//	state is a BOOL that is TRUE for state of BST_DROPDOWNPUSHED, or FALSE otherwise.
func (b *Button) SetDropDownState(state bool) {
	wParam := 0
	if state {
		wParam = 1
	}
	b.SendMessage(win.BCM_SETDROPDOWNSTATE, uintptr(wParam), 0)
}

//	SetShield sets the elevation required state for a specified button or command link to display an elevated icon.
//	state is a BOOL that is TRUE to draw an elevated icon, or FALSE otherwise.
func (b *Button) SetShield(state bool) {
	wParam := 0
	if state {
		wParam = 1
	}
	b.SendMessage(win.BCM_SETSHIELD, uintptr(wParam), 0)
}

//	Click simulates the user clicking a button. This message causes the button to receive the WM_LBUTTONDOWN and
//	WM_LBUTTONUP messages, and the button's parent window to receive a BN_CLICKED notification code.
func (b *Button) Click() {
	b.SendMessage(win.BM_CLICK, 0, 0)
}

//	GetCheck gets the check state of a radio button or check box.
//	
//	checked return 1, unchecked return 0
//	
//	if Button has BS_3STATE or BS_AUTO3STATE style, may be return  2
//	
//	BST_CHECKED 1
//	
//	BST_INDETERMINATE 2
//	
//	BST_UNCHECKED 0
func (b *Button) GetCheck() int {
	return int(b.SendMessage(win.BM_GETCHECK, 0, 0))
}

//	GetState retrieves the state of a button or check box.
//	
//	The return value specifies the current state of the button. It is a combination of the following values.
//
//	Return code	and description:
//
//	BST_CHECKED
//	The button is checked.
//	
//	BST_DROPDOWNPUSHED
//	Windows Vista. The button is in the drop-down state. Applies only if the button has the TBSTYLE_DROPDOWN style.
//	
//	BST_FOCUS
//	The button has the keyboard focus.
//	
//	BST_HOT
//	The button is hot; that is, the mouse is hovering over it.
//	
//	BST_INDETERMINATE
//	The state of the button is indeterminate. Applies only if the button has the BS_3STATE or BS_AUTO3STATE style.
//	
//	BST_PUSHED
//	The button is being shown in the pushed state.
//
//	BST_UNCHECKED
//	No special state. Equivalent to zero.
func (b *Button) GetState() int {
	return int(b.SendMessage(win.BM_GETSTATE, 0, 0))
}

//	SetCheck Sets the check state of a radio button or check box,
//	state is the check state. This parameter can be one of the following values.
//
//	Value	Meaning:
//	BST_CHECKED 1
//	Sets the button state to checked.
//
//	BST_INDETERMINATE 2
//	Sets the button state to grayed, indicating an indeterminate state. Use this value only
//	if the button has the BS_3STATE or BS_AUTO3STATE style.
//
//	BST_UNCHECKED 0
//	Sets the button state to cleared.
func (b *Button) SetCheck(state int) {
	b.SendMessage(win.BM_SETCHECK, uintptr(state), 0)
}

//	SetDontClick sets a flag on a radio button that controls the generation of BN_CLICKED messages when
//	the button receives focus. notice: WINVER >= 0x0600
func (b *Button) SetDontClick(state bool) {
	wParam := 0
	if state {
		wParam = 1
	}
	b.SendMessage(winapi.BM_SETDONTCLICK, uintptr(wParam), 0)
}

//	SetImage associates a new image (icon or bitmap) with the button.
//	t is the type of image to associate with the button. This parameter can be one of the following values:
//	- IMAGE_BITMAP
//	- IMAGE_ICON
//
//	handle is a handle (HICON or HBITMAP) to the image to associate with the button.
//
//	The return value is a handle to the image previously associated with the button, if any; otherwise, it is NULL.
func (b *Button) SetImage(t int, handle uintptr) uintptr {
	return b.SendMessage(win.BM_SETIMAGE, uintptr(t), handle)
}

//	GetImage retrieves a handle to the image (icon or bitmap) associated with the button.
//	t is the type of image to associate with the button. This parameter can be one of the following values:
//	- IMAGE_BITMAP
//	- IMAGE_ICON
//
//	The return value is a handle to the image, if any; otherwise, it is NULL.
func (b *Button) GetImage(t int) uintptr {
	return b.SendMessage(win.BM_GETIMAGE, uintptr(t), 0)
}

//	SetState sets the highlight state of a button.
//	The highlight state indicates whether the button is highlighted as if the user had pushed it.
//
//	Highlighting affects only the appearance of a button.
//	It has no effect on the check state of a radio button or check box.
//
// A button is automatically highlighted when the user positions the cursor over it
// and presses and holds the left mouse button. The highlighting is removed when the user releases the mouse button.
func (b *Button) SetState(highlight bool) {
	wParam := 0
	if highlight {
		wParam = 1
	}
	b.SendMessage(win.BM_SETSTATE, uintptr(wParam), 0)
}

//	SetStyle sets the style of a button.
//	The button style can be a combination of button styles. For a table of button styles,
//
//	Constant	Description:
//
//	BS_3STATE
//	Creates a button that is the same as a check box, except that the box can be grayed as
//	well as checked or cleared. Use the grayed state to show that the state of the check box is not determined.
//
//	BS_AUTO3STATE
//	Creates a button that is the same as a three-state check box, except that the box changes
//	its state when the user selects it. The state cycles through checked, indeterminate, and cleared.
//
//	BS_AUTOCHECKBOX
//	Creates a button that is the same as a check box, except that the check state automatically
//	toggles between checked and cleared each time the user selects the check box.
//
//	BS_AUTORADIOBUTTON
//	Creates a button that is the same as a radio button, except that when the user selects it,
//	the system automatically sets the button's check state to checked and automatically sets the
//	check state for all other buttons in the same group to cleared.
//
//	BS_BITMAP
//	Specifies that the button displays a bitmap. See the Remarks section for its interaction with BS_ICON.
//
//	BS_BOTTOM
//	Places text at the bottom of the button rectangle.
//
//	BS_CENTER
//	Centers text horizontally in the button rectangle.
//
//	BS_CHECKBOX
//	Creates a small, empty check box with text. By default, the text is displayed to the right
//	of the check box. To display the text to the left of the check box, combine this flag with
//	the BS_LEFTTEXT style (or with the equivalent BS_RIGHTBUTTON style).
//
//	BS_COMMANDLINK
//	Creates a command link button that behaves like a BS_PUSHBUTTON style button, but the command
//	link button has a green arrow on the left pointing to the button text. A caption for the button
//	text can be set by sending the BCM_SETNOTE message to the button.
//
//	BS_DEFCOMMANDLINK
//	Creates a command link button that behaves like a BS_PUSHBUTTON style button. If the button is
//	in a dialog box, the user can select the command link button by pressing the ENTER key,
//	even when the command link button does not have the input focus. This style is useful for
//	enabling the user to quickly select the most likely (default) option.
//
//	BS_DEFPUSHBUTTON
//	Creates a push button that behaves like a BS_PUSHBUTTON style button, but has a distinct appearance.
//	If the button is in a dialog box, the user can select the button by pressing the ENTER key,
//	even when the button does not have the input focus. This style is useful for enabling the user to
//	quickly select the most likely (default) option.
//
//	BS_DEFSPLITBUTTON
//	Creates a split button that behaves like a BS_PUSHBUTTON style button, but also has a distinctive appearance.
//	If the split button is in a dialog box, the user can select the split button by pressing the ENTER key,
//	even when the split button does not have the input focus. This style is useful for enabling the user to
//	quickly select the most likely (default) option.
//
//	BS_GROUPBOX
//	Creates a rectangle in which other controls can be grouped. Any text associated with this style is
//	displayed in the rectangle's upper left corner.
//
//	BS_ICON
//	Specifies that the button displays an icon. See the Remarks section for its interaction with BS_BITMAP.
//
//	BS_FLAT
//	Specifies that the button is two-dimensional; it does not use the default shading to create a 3-D image.
//
//	BS_LEFT
//	Left-justifies the text in the button rectangle. However, if the button is a check box or radio
//	button that does not have the BS_RIGHTBUTTON style, the text is left justified on the right side
//	of the check box or radio button.
//
//	BS_LEFTTEXT
//	Places text on the left side of the radio button or check box when combined with a radio button or
//	check box style. Same as the BS_RIGHTBUTTON style.
//
//	BS_MULTILINE
//	Wraps the button text to multiple lines if the text string is too long to fit on a single line
//	in the button rectangle.
//
//	BS_NOTIFY
//	Enables a button to send BN_KILLFOCUS and BN_SETFOCUS notification codes to its parent window.
//
//	Note that buttons send the BN_CLICKED notification code regardless of whether it has this style.
//	To get BN_DBLCLK notification codes, the button must have the BS_RADIOBUTTON or BS_OWNERDRAW style.
//
//	BS_OWNERDRAW
//	Creates an owner-drawn button. The owner window receives a WM_DRAWITEM message when a visual
//	aspect of the button has changed. Do not combine the BS_OWNERDRAW style with any other button styles.
//
//	BS_PUSHBUTTON
//	Creates a push button that posts a WM_COMMAND message to the owner window when the user selects the button.
//
//	BS_PUSHLIKE
//	Makes a button (such as a check box, three-state check box, or radio button) look and act
//	like a push button. The button looks raised when it isn't pushed or checked, and sunken when
//	it is pushed or checked.
//
//	BS_RADIOBUTTON
//	Creates a small circle with text. By default, the text is displayed to the right of the circle.
//	To display the text to the left of the circle, combine this flag with the BS_LEFTTEXT style
//	(or with the equivalent BS_RIGHTBUTTON style). Use radio buttons for groups of related,
//	but mutually exclusive choices.
//
//	BS_RIGHT
//	Right-justifies text in the button rectangle. However, if the button is a check box or
//	radio button that does not have the BS_RIGHTBUTTON style, the text is right justified on
//	the right side of the check box or radio button.
//
//	BS_RIGHTBUTTON
//	Positions a radio button's circle or a check box's square on the right side of the button rectangle.
//	Same as the BS_LEFTTEXT style.
//
//	BS_SPLITBUTTON
//	Creates a split button. A split button has a drop down arrow.
//
//	BS_TEXT
//	Specifies that the button displays text.
//
//	BS_TOP
//	Places text at the top of the button rectangle.
//
//	BS_TYPEMASK
//	Do not use this style. A composite style bit that results from using the OR operator on
//	BS_* style bits. It can be used to mask out valid BS_* bits from a given bitmask.
//	Note that this is out of date and does not correctly include all valid styles.
//	Thus, you should not use this style.
//
//	BS_USERBUTTON
//	Obsolete, but provided for compatibility with 16-bit versions of Windows.
//	Applications should use BS_OWNERDRAW instead.
//
//	BS_VCENTER
//	Places text in the middle (vertically) of the button rectangle.
//
//	https://docs.microsoft.com/zh-cn/windows/win32/controls/button-styles
func (b *Button) SetStyle(style int, redrawn bool) {
	lParam := 0
	if redrawn {
		lParam = 1
	}
	b.SendMessage(win.BM_SETSTYLE, uintptr(style), uintptr(lParam))
}

//	NewButton create a new Button,need bind to Dialog before use.
func NewButton(idd uintptr) *Button {
	return &Button{WindowBase: WindowBase{idd: idd}}
}

//	BindNewButton create a new Button and bind to target dlg.
func BindNewButton(idd uintptr, dlg *Dialog) (*Button, error) {
	b := NewButton(idd)
	err := dlg.BindWidgets(b)
	return b, err
}

func BindNewButtons(idds []uintptr, dlg *Dialog) ([]*Button, error) {
	btns := make([]*Button, len(idds))
	var err error
	for idx, idd := range idds {
		btns[idx], err = BindNewButton(idd, dlg)
	}
	return btns, err
}
