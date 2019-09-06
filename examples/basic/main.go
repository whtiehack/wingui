package main

/*
// TODO: resource ID 转 GO代码
#include "ui/resource.h"
*/
import "C"
import (
	"github.com/whtiehack/wingui"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	log.Printf("resource %v %#[1]v  \n", C.IDD_DIALOG)
	dlg, err := wingui.NewDialog(C.IDD_DIALOG, 0)
	if err != nil {
		log.Panic("main dialog create error", err)
	}
	log.Println("dlg create end", dlg)
	dlg.Show()

	dlg, err = wingui.NewDialog(C.IDD_DIALOG, dlg.Handle())
	dlg.Show()
	wingui.MessageLoop()
	log.Println("stoped")
}
