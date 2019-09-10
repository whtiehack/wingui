package main

import (
	"github.com/whtiehack/wingui"
	"log"
)

var dlg *wingui.Dialog

// optional  genereate resource IDs
//go:generate go run github.com/whtiehack/wingui/tools/genids -filename ui/resource.h -packagename main
func main() {
	var err error
	dlg, err = wingui.NewDialog(IDD_DIALOG_MAIN, 0)
	if err != nil {
		log.Panic("main dialog create error", err)
	}
	dlg.SetIcon(IDI_ICON1)
	log.Println("dlg create end", dlg)

	dlg.Show()
	wingui.MessageLoop()
	log.Println("stoped")
}
