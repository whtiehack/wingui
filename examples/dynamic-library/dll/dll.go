package main

import "C"
import "log"

func init() {
	log.Println("dll init")
}

//export Showui
func Showui() {
	log.Println("show ui called")
	//wingui.InitHInstance("test.dll")
	//wingui.NewModalDialog(IDD_DIALOG1, 0, func(dlg *wingui.Dialog) {
	//	// do others
	//})

}

//go:generate go run github.com/whtiehack/wingui/tools/genids -filename ui/resource.h -packagename main
func main() {
	// Need a main function to make CGO compile package as C shared library

}
