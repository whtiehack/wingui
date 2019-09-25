package wingui

// https://docs.microsoft.com/zh-cn/windows/win32/controls/trackbar-controls

//	TrackBar a widget for Dialog. Progress Bar
type TrackBar struct {
	WindowBase
}

func (tb *TrackBar) ClearSel(redraw bool) {
	//	tb.SendMessage(winapi.TBM_CLEARSEL)
}
