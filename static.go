package wingui

import (
	"github.com/lxn/win"
	"log"
)

// Static a static label widget for Dialog.
type Static struct {
	WindowBase
	// Color must set before Dialog init,should bind use DialogConfig.
	Color win.COLORREF
	//BkMode must set same as Color
	BkMode int32
}

// WndProc Button window WndProc.
func (b *Static) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	log.Print("static msg", msg)
	switch msg {
	case win.WM_CTLCOLORSTATIC:
		if b.Color != 0 && b.BkMode != 0 {
			win.SetTextColor(win.HDC(wParam), b.Color); //设置文本颜色
			win.SetBkMode(win.HDC(wParam), b.BkMode);   //设置背景透明
			//返回一个空画刷(必须)
			hb := win.GetStockObject(win.NULL_BRUSH)
			return uintptr(hb)
		}
		return 0
	}
	return b.AsWindowBase().WndProc(msg, wParam, lParam)
}

// NewStatic create a new Button,need bind to Dialog for use.
func NewStatic(idd uintptr) *Static {
	return &Static{
		WindowBase: WindowBase{idd: idd},
		Color:      0,
		BkMode:     0,
	}
}

// BindNewStatic create a new Button and bind to target dlg.
func BindNewStatic(idd uintptr, dlg *Dialog) (*Static, error) {
	b := &Static{WindowBase{idd: idd}, 0, 0}
	err := dlg.BindWidgets(b)
	return b, err
}
