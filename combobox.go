package wingui

// ComboBox a ComboBox widget for Dialog.
type ComboBox struct {
	WindowBase
}

// WndProc ComboBox window WndProc.
func (b *ComboBox) WndProc(msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	}
	return b.AsWindowBase().WndProc(msg, wParam, lParam)
}

// NewComboBox create a new ComboBox,need bind to Dialog before use.
func NewComboBox(idd uintptr) *ComboBox {
	return &ComboBox{
		WindowBase: WindowBase{idd: idd},
	}
}

// BindNewStatic create a new Image and bind to target dlg.
func BindNewComboBox(idd uintptr, dlg *Dialog) (*ComboBox, error) {
	b := NewComboBox(idd)
	err := dlg.BindWidgets(b)
	return b, err
}
