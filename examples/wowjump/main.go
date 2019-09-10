package main

import (
	"github.com/lxn/win"
	"github.com/whtiehack/wingui"
	"github.com/whtiehack/wingui/kernel32"
	"log"
	"os"
	"syscall"
)

var dlg *wingui.Dialog

// 防止进程多开，返回 true 表示进程已经开启
func ProcessMutex(name string) bool {
	r, err := kernel32.OpenMutex(2031617, 1, name)
	if err == nil || r != 0 {
		return true
	}
	r, err = kernel32.CreateMutex(name)
	if err != nil || r == 0 {
		return true
	}
	return false
}

func init() {
	log.SetOutput(os.Stdout)
	if ProcessMutex("@@wowjump@@") {
		win.MessageBox(0, &syscall.StringToUTF16("进程已经开启了，不可以多开")[0], nil, 0)
		os.Exit(-1)
	}
}

// optional  genereate resource IDs
//go:generate go run github.com/whtiehack/wingui/tools/genids -filename ui/resource.h -packagename main
func main() {
	var err error
	dlg, err = wingui.NewDialog(IDD_DIALOG_MAIN, 0, nil)
	if err != nil {
		log.Panic("main dialog create error", err)
	}
	dlg.SetIcon(IDI_ICON1)
	log.Println("dlg create end", dlg)
	config.editNormaltime, _ = dlg.NewEdit(IDE_NORMAL_TIME)
	config.editEnterTime, _ = dlg.NewEdit(IDE_ENTER_TIME)
	config.editInputTime, _ = dlg.NewEdit(IDE_INPUT_TIME)
	config.editCharWaitTime, _ = dlg.NewEdit(IDE_CHAR_WAIT_TIME)
	config.InitVal()
	dlg.Show()
	wingui.SetCurrentDialog(dlg.Handle())
	wingui.MessageLoop()
	log.Println("stoped")
}
