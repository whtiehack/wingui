package main

import (
	"github.com/whtiehack/wingui"
	"log"
)

func init() {
	wingui.InitHInstance("test.dll")
	log.Println("dll init")
	dlg, err := wingui.NewDialog(IDD_DIALOG1, 0)
	if err != nil {
		log.Panic("dll main dialog create error", err)
	}
	dlg.Show()
}

//go:generate go run github.com/whtiehack/wingui/tools/genids -filename ui/resource.h -packagename main
func main() {

}
