package wingui

import "github.com/whtiehack/wingui/winapi"

// https://docs.microsoft.com/zh-cn/windows/win32/controls/trackbar-controls

//	TrackBar a widget for Dialog. Progress Bar
type TrackBar struct {
	WindowBase
}

func (tb *TrackBar) ClearSel(redraw bool) {
	rd := 0
	if redraw {
		rd = 1
	}
	tb.SendMessage(winapi.TBM_CLEARSEL, uintptr(rd), 0)
}

// NewTrackBar create a new TrackBar,need bind to Dialog before use.
func NewTrackBar(idd uintptr) *TrackBar {
	return &TrackBar{
		WindowBase: WindowBase{idd: idd},
	}
}

// BindNewTrackBar create a new TrackBar and bind to target dlg.
func BindNewTrackBar(idd uintptr, dlg *Dialog) (*TrackBar, error) {
	bb := NewTrackBar(idd)
	err := dlg.BindWidgets(bb)
	return bb, err
}
