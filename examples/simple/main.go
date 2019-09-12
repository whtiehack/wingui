package main

import "github.com/whtiehack/wingui"

func main() {
	dlg, _ := wingui.NewDialog(101, 0, nil)
	dlg.SetIcon(105)
	btnok, _ := wingui.BindNewButton(1002, dlg)
	btncancel, _ := wingui.BindNewButton(1003, dlg)
	btnok.OnClicked = func() {
		dlg.Close()
	}
	btncancel.OnClicked = btnok.OnClicked
	dlg.Show()
	// This invoke is optional.
	wingui.SetCurrentDialog(dlg.Handle())
	wingui.MessageLoop()
}
