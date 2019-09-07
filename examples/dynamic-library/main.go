package main

import (
	"github.com/lxn/win"
	"github.com/whtiehack/wingui"
	"log"
	"os"
	"syscall"
	"time"
)

func main() {
	win.MessageBox(0, nil, nil, 0)
	go func() {
		time.Sleep(2 * time.Second)
		loaddll()
	}()
	wingui.MessageLoop()
}

func loaddll() {
	cw, _ := os.Getwd()
	h, err := syscall.LoadLibrary(cw + "/dll/test.dll")
	log.Println("load dll", h, err)
}
