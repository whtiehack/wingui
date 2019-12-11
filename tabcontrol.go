package wingui

// TabControl a widget for Dialog. Tab Control
type TabControl struct {
	WindowBase
}

// NewTabControl create a new TabControl, need bind to Dialog before use.
func NewTabControl(idd uintptr) *TabControl {
	return &TabControl{
		WindowBase: WindowBase{idd: idd},
	}
}

// BindTabControl create a new TabControl and bind to target dlg.
func BindTabControl(idd uintptr, dlg *Dialog) (*TabControl, error) {
	tb := NewTabControl(idd)
	err := dlg.BindWidgets(tb)
	return tb, err
}
