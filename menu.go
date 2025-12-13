package wingui

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"github.com/whtiehack/wingui/winapi"
)

/*
Manual menu creation example:

	const (
		idExit   = 60001
		idHello  = 60002
		idToggle = 60003
	)

	menu := wingui.NewMenu()
	fileMenu := wingui.NewPopupMenu()
	_ = fileMenu.AppendItem(idExit, "Exit")
	_ = menu.AppendSubMenu(fileMenu, "File")

	demoMenu := wingui.NewPopupMenu()
	_ = demoMenu.AppendItem(idHello, "Hello MessageBox")
	_ = demoMenu.AppendItem(idToggle, "Toggle Checked")
	_ = menu.AppendSubMenu(demoMenu, "Demo")

	_ = dlg.SetMenu(menu)

	checked := false
	dlg.OnCommand = func(id uint16) {
		switch id {
		case idExit:
			dlg.Close()
		case idHello:
			wingui.MessageBox(dlg.Handle(), "Hello from Menu", "Menu", win.MB_OK|win.MB_ICONINFORMATION)
		case idToggle:
			checked = !checked
			demoMenu.CheckItem(idToggle, checked)
		}
	}
*/

// Menu wraps a Win32 HMENU.
type Menu struct {
	hMenu win.HMENU
	count uint32
}

// NewMenuFromResource loads a menu resource by id (from the module hInstance).
func NewMenuFromResource(id uintptr) *Menu {
	h := win.LoadMenu(hInstance, win.MAKEINTRESOURCE(id))
	if h == 0 {
		return nil
	}
	return &Menu{hMenu: h}
}

// Handle returns the underlying HMENU.
func (m *Menu) Handle() win.HMENU {
	if m == nil {
		return 0
	}
	return m.hMenu
}

// SubMenu returns the submenu at the given position (0-based), or nil.
func (m *Menu) SubMenu(pos int32) *Menu {
	if m == nil || m.hMenu == 0 {
		return nil
	}
	h := winapi.GetSubMenu(m.hMenu, pos)
	if h == 0 {
		return nil
	}
	return &Menu{hMenu: h}
}

// Destroy releases the underlying HMENU.
func (m *Menu) Destroy() bool {
	if m == nil || m.hMenu == 0 {
		return false
	}
	ok := win.DestroyMenu(m.hMenu)
	if ok {
		m.hMenu = 0
	}
	return ok
}

// NewMenu creates a menu bar menu.
func NewMenu() *Menu {
	return &Menu{hMenu: win.CreateMenu()}
}

// NewPopupMenu creates a popup menu.
func NewPopupMenu() *Menu {
	return &Menu{hMenu: win.CreatePopupMenu()}
}

// AppendItem appends a clickable item.
func (m *Menu) AppendItem(id uint16, text string) bool {
	if m == nil || m.hMenu == 0 {
		return false
	}
	utf16 := syscall.StringToUTF16(text)
	mii := win.MENUITEMINFO{
		CbSize:     uint32(unsafe.Sizeof(win.MENUITEMINFO{})),
		FMask:      win.MIIM_ID | win.MIIM_STRING | win.MIIM_FTYPE | win.MIIM_STATE,
		FType:      win.MFT_STRING,
		FState:     win.MFS_ENABLED,
		WID:        uint32(id),
		DwTypeData: &utf16[0],
		Cch:        uint32(len(utf16) - 1),
	}
	ok := win.InsertMenuItem(m.hMenu, m.count, true, &mii)
	if ok {
		m.count++
	}
	return ok
}

// AppendSeparator appends a separator line.
func (m *Menu) AppendSeparator() bool {
	if m == nil || m.hMenu == 0 {
		return false
	}
	mii := win.MENUITEMINFO{
		CbSize: uint32(unsafe.Sizeof(win.MENUITEMINFO{})),
		FMask:  win.MIIM_FTYPE,
		FType:  win.MFT_SEPARATOR,
	}
	ok := win.InsertMenuItem(m.hMenu, m.count, true, &mii)
	if ok {
		m.count++
	}
	return ok
}

// AppendSubMenu appends a submenu.
func (m *Menu) AppendSubMenu(sub *Menu, text string) bool {
	if m == nil || m.hMenu == 0 || sub == nil || sub.hMenu == 0 {
		return false
	}
	utf16 := syscall.StringToUTF16(text)
	mii := win.MENUITEMINFO{
		CbSize:     uint32(unsafe.Sizeof(win.MENUITEMINFO{})),
		FMask:      win.MIIM_SUBMENU | win.MIIM_STRING | win.MIIM_FTYPE | win.MIIM_STATE,
		FType:      win.MFT_STRING,
		FState:     win.MFS_ENABLED,
		HSubMenu:   sub.hMenu,
		DwTypeData: &utf16[0],
		Cch:        uint32(len(utf16) - 1),
	}
	ok := win.InsertMenuItem(m.hMenu, m.count, true, &mii)
	if ok {
		m.count++
	}
	return ok
}

// EnableItem enables/disables an item.
func (m *Menu) EnableItem(id uint16, enabled bool) {
	if m == nil || m.hMenu == 0 {
		return
	}
	state := uint32(win.MFS_ENABLED)
	if !enabled {
		state = win.MFS_DISABLED
	}
	mii := win.MENUITEMINFO{
		CbSize: uint32(unsafe.Sizeof(win.MENUITEMINFO{})),
		FMask:  win.MIIM_STATE,
		FState: state,
	}
	_ = win.SetMenuItemInfo(m.hMenu, uint32(id), false, &mii)
}

// CheckItem checks/unchecks an item.
func (m *Menu) CheckItem(id uint16, checked bool) {
	if m == nil || m.hMenu == 0 {
		return
	}
	state := uint32(win.MFS_UNCHECKED)
	if checked {
		state = win.MFS_CHECKED
	}
	mii := win.MENUITEMINFO{
		CbSize: uint32(unsafe.Sizeof(win.MENUITEMINFO{})),
		FMask:  win.MIIM_STATE,
		FState: state,
	}
	_ = win.SetMenuItemInfo(m.hMenu, uint32(id), false, &mii)
}

// TrackPopup shows a popup menu at screen coordinates.
// Returns the selected command id, or 0 if canceled.
func (m *Menu) TrackPopup(owner win.HWND, x, y int32, alignRight bool) uint16 {
	if m == nil || m.hMenu == 0 {
		return 0
	}
	flags := uint32(win.TPM_RETURNCMD | win.TPM_NONOTIFY)
	if alignRight {
		flags |= win.TPM_RIGHTALIGN
	} else {
		flags |= win.TPM_LEFTALIGN
	}
	cmd := winapi.TrackPopupMenuEx(m.hMenu, flags, x, y, owner)
	return uint16(cmd)
}

// SetMenu sets the menu bar for the window and redraws it.
func (w *WindowBase) SetMenu(menu *Menu) bool {
	if w == nil || w.hwnd == 0 {
		return false
	}
	ok := win.SetMenu(w.hwnd, menu.Handle())
	if ok {
		win.DrawMenuBar(w.hwnd)
	}
	return ok
}
