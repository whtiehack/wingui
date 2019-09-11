package main

import (
	"log"
	"os"

	"github.com/whtiehack/wingui"
)

func init() {
	log.SetOutput(os.Stdout)
}

var dlg *wingui.Dialog

// optional  genereate resource IDs
//go:generate go run github.com/whtiehack/wingui/tools/genids -filename ui/resource.h -packagename main
func main() {
	log.Printf("resource %v %#[1]v  \n", IDD_DIALOG)
	var err error
	dlg, err = wingui.NewDialog(IDD_DIALOG, 0, nil)
	if err != nil {
		log.Panic("main dialog create error", err)
	}
	dlg.SetIcon(IDI_ICON1)
	log.Println("dlg create end", dlg)
	btn := wingui.NewButton(IDB_OK)
	btn.OnClicked = modalBtnClicked
	closeBtn := wingui.NewButton(IDB_CANCEL)
	closeBtn.OnClicked = func() {
		dlg.Close()
	}
	dlg.BindWidgets(btn, closeBtn)
	dlg.Show()
	wingui.SetCurrentDialog(dlg.Handle())
	wingui.MessageLoop()
	log.Println("stoped")
}

func modalBtnClicked() {
	log.Println("btn clicked")
	wingui.NewModalDialog(IDD_DIALOG_OK, dlg.Handle(), nil, func(okdlg *wingui.Dialog) {
		okbtn := wingui.NewButton(IDB_OK)
		okbtn.OnClicked = func() {
			log.Println("modal btn click")
			okdlg.Close()
		}
		okdlg.BindWidgets(okbtn)
	})
}
