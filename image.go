package wingui

// https://www.daniweb.com/programming/software-development/code/216348/displaying-a-jpeg-image-using-windows-gui

import (
	"github.com/lxn/win"
)

// Image a static image widget for Dialog.
type Image struct {
	WindowBase
	// OnClicked must set appearance Notify to true before use.
	OnClicked func()
}

// WndProc Image window WndProc.
func (b *Image) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	// log.Print("static msg", msg, wParam, lParam)
	switch msg {
	case win.WM_CTLCOLORSTATIC:
	case win.WM_COMMAND:
		if b.OnClicked != nil {
			b.OnClicked()
		}
		break
	}
	return b.AsWindowBase().WndProc(msg, wParam, lParam)
}

// LoadBitmap show  Bitmap on window. Don't forget to set the type to Bitmap
func (b *Image) LoadBitmap(bitmap *Bitmap) uintptr {
	return b.SendMessage(0x172, win.IMAGE_BITMAP, uintptr(bitmap.hBmp))
}

// NewImage create a new Image,need bind to Dialog before use.
func NewImage(idd uintptr) *Image {
	return &Image{
		WindowBase: WindowBase{idd: idd},
		OnClicked:  nil,
	}
}

// BindNewStatic create a new Image and bind to target dlg.
func BindNewImage(idd uintptr, dlg *Dialog) (*Image, error) {
	b := NewImage(idd)
	err := dlg.BindWidgets(b)
	return b, err
}

/*

 // Static Control Mesages

#define STM_SETICON         0x0170
#define STM_GETICON         0x0171
#define STM_SETIMAGE        0x0172
#define STM_GETIMAGE        0x0173
#define STN_CLICKED         0
#define STN_DBLCLK          1
#define STN_ENABLE          2
#define STN_DISABLE         3
*/
