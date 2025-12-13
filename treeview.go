package wingui

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

// TreeView is a wrapper for the Win32 SysTreeView32 control.
type TreeView struct {
	WindowBase

	// OnSelChanged fires after selection has changed.
	OnSelChanged func(item win.HTREEITEM)
	// OnDblClick fires on NM_DBLCLK; item can be 0 if none.
	OnDblClick func(item win.HTREEITEM)
	// OnItemExpanding fires on TVN_ITEMEXPANDING. Return true to cancel.
	OnItemExpanding func(item win.HTREEITEM, action uint32) (cancel bool)
}

const (
	tvgnRoot     = 0
	tvgnNext     = 1
	tvgnPrevious = 2
	tvgnParent   = 3
	tvgnChild    = 4
)

func (tv *TreeView) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_NOTIFY:
		nmhdr := (*win.NMHDR)(unsafe.Pointer(lParam))
		code := uint32(nmhdr.Code)

		switch code {
		case win.TVN_SELCHANGED:
			if tv.OnSelChanged != nil {
				nmtv := (*win.NMTREEVIEW)(unsafe.Pointer(lParam))
				tv.OnSelChanged(nmtv.ItemNew.HItem)
			}
			return 0
		case win.TVN_ITEMEXPANDING:
			if tv.OnItemExpanding != nil {
				nmtv := (*win.NMTREEVIEW)(unsafe.Pointer(lParam))
				if tv.OnItemExpanding(nmtv.ItemNew.HItem, nmtv.Action) {
					return 1
				}
			}
			return 0
		case win.NM_DBLCLK:
			if tv.OnDblClick != nil {
				tv.OnDblClick(tv.GetSelection())
			}
			return 0
		}
		return 0
	}
	return tv.AsWindowBase().WndProc(msg, wParam, lParam)
}

// ItemCount returns the total number of items.
func (tv *TreeView) ItemCount() int {
	return int(int32(tv.SendMessage(win.TVM_GETCOUNT, 0, 0)))
}

// InsertItem inserts a new item and returns its handle (0 on failure).
// parent can be win.TVI_ROOT; insertAfter can be win.TVI_LAST / win.TVI_SORT, etc.
func (tv *TreeView) InsertItem(parent, insertAfter win.HTREEITEM, text string) win.HTREEITEM {
	utf16 := syscall.StringToUTF16(text)
	ins := win.TVINSERTSTRUCT{
		HParent:      parent,
		HInsertAfter: insertAfter,
		Item: win.TVITEM{
			Mask:    win.TVIF_TEXT,
			PszText: uintptr(unsafe.Pointer(&utf16[0])),
		},
	}
	ret := tv.SendMessage(win.TVM_INSERTITEM, 0, uintptr(unsafe.Pointer(&ins)))
	return win.HTREEITEM(ret)
}

// DeleteItem deletes an item (or all children if item is win.TVI_ROOT) and returns success.
func (tv *TreeView) DeleteItem(item win.HTREEITEM) bool {
	return tv.SendMessage(win.TVM_DELETEITEM, 0, uintptr(item)) != 0
}

// Expand expands/collapses an item. action is one of win.TVE_*.
func (tv *TreeView) Expand(item win.HTREEITEM, action uint32) bool {
	return tv.SendMessage(win.TVM_EXPAND, uintptr(action), uintptr(item)) != 0
}

// EnsureVisible scrolls the tree view so the item is visible.
func (tv *TreeView) EnsureVisible(item win.HTREEITEM) bool {
	return tv.SendMessage(win.TVM_ENSUREVISIBLE, 0, uintptr(item)) != 0
}

// GetSelection returns the currently selected item (0 if none).
func (tv *TreeView) GetSelection() win.HTREEITEM {
	return win.HTREEITEM(tv.SendMessage(win.TVM_GETNEXTITEM, win.TVGN_CARET, 0))
}

// SelectItem selects an item.
func (tv *TreeView) SelectItem(item win.HTREEITEM) bool {
	return tv.SendMessage(win.TVM_SELECTITEM, win.TVGN_CARET, uintptr(item)) != 0
}

// SetItemText updates the item text.
func (tv *TreeView) SetItemText(item win.HTREEITEM, text string) bool {
	utf16 := syscall.StringToUTF16(text)
	tvi := win.TVITEM{
		Mask:    win.TVIF_TEXT,
		HItem:   item,
		PszText: uintptr(unsafe.Pointer(&utf16[0])),
	}
	return tv.SendMessage(win.TVM_SETITEM, 0, uintptr(unsafe.Pointer(&tvi))) != 0
}

// GetItemText reads the item text.
func (tv *TreeView) GetItemText(item win.HTREEITEM) string {
	buf := make([]uint16, 512)
	tvi := win.TVITEM{
		Mask:       win.TVIF_TEXT,
		HItem:      item,
		PszText:    uintptr(unsafe.Pointer(&buf[0])),
		CchTextMax: int32(len(buf)),
	}
	if tv.SendMessage(win.TVM_GETITEM, 0, uintptr(unsafe.Pointer(&tvi))) == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf)
}

// ParentItem returns the parent item or 0.
func (tv *TreeView) ParentItem(item win.HTREEITEM) win.HTREEITEM {
	return win.HTREEITEM(tv.SendMessage(win.TVM_GETNEXTITEM, tvgnParent, uintptr(item)))
}

// NextSibling returns the next sibling item or 0.
func (tv *TreeView) NextSibling(item win.HTREEITEM) win.HTREEITEM {
	return win.HTREEITEM(tv.SendMessage(win.TVM_GETNEXTITEM, tvgnNext, uintptr(item)))
}

// PrevSibling returns the previous sibling item or 0.
func (tv *TreeView) PrevSibling(item win.HTREEITEM) win.HTREEITEM {
	return win.HTREEITEM(tv.SendMessage(win.TVM_GETNEXTITEM, tvgnPrevious, uintptr(item)))
}

// ChildItem returns the first child item or 0.
func (tv *TreeView) ChildItem(item win.HTREEITEM) win.HTREEITEM {
	return win.HTREEITEM(tv.SendMessage(win.TVM_GETNEXTITEM, tvgnChild, uintptr(item)))
}

// NewTreeView creates a new TreeView, need bind to Dialog before use.
func NewTreeView(idd uintptr) *TreeView {
	return &TreeView{WindowBase: WindowBase{idd: idd}}
}

// BindNewTreeView creates a new TreeView and bind to target dlg.
func BindNewTreeView(idd uintptr, dlg *Dialog) (*TreeView, error) {
	tv := NewTreeView(idd)
	err := dlg.BindWidgets(tv)
	return tv, err
}
