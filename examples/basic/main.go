package main

import (
	"github.com/whtiehack/wingui"
	"github.com/whtiehack/wingui/examples/basic/ui"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

var dlg *wingui.Dialog

// optional  genereate resource IDs
//go:generate go run github.com/whtiehack/wingui/tools/genids -filename ui/resource.h
func main() {
	log.Printf("resource %v %#[1]v  \n", ui.IDD_DIALOG)
	var err error
	dlg, err = wingui.NewDialog(ui.IDD_DIALOG, 0)
	if err != nil {
		log.Panic("main dialog create error", err)
	}
	log.Println("dlg create end", dlg)
	var btn *wingui.Button
	btn, _ = dlg.NewButton(ui.IDB_OK)
	btn.OnClicked = modalBtnClicked
	closeBtn, _ := dlg.NewButton(ui.IDB_CANCEL)
	closeBtn.OnClicked = func() {
		dlg.Close()
	}
	dlg.Show()
	wingui.MessageLoop()
	log.Println("stoped")
}

func modalBtnClicked() {
	log.Println("btn clicked")
	wingui.NewModalDialog(ui.IDD_DIALOG_OK, dlg.Handle(), func(okdlg *wingui.Dialog) {
		okbtn, _ := okdlg.NewButton(ui.IDB_OK)
		okbtn.OnClicked = func() {
			log.Println("modal btn click")
			okdlg.Close()
		}
	})
}
