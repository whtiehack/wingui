package main

import "C"
import (
	"github.com/whtiehack/wingui"
	"log"
)

func init() {
	log.Println("dll init")
	wingui.InitHInstance("test.dll")
}

//export Showui
func Showui() {
	log.Println("show ui called")
	wingui.NewModalDialog(IDD_DIALOG1, 0, nil, dlgcb)
	log.Println("dlg end")
}

func dlgcb(dlg *wingui.Dialog) {
	// do others
	close := func() {
		dlg.Close()
	}
	btn1, _ := dlg.NewButton(IDB_OK)
	btn2, _ := dlg.NewButton(IDB_CANCEL)
	btn1.OnClicked, btn2.OnClicked = close, close
}

//go:generate go run github.com/whtiehack/wingui/tools/genids -filename ui/resource.h -packagename main
func main() {
	// Need a main function to make CGO compile package as C shared library
	Showui()
	log.Println("show ui end")
}
