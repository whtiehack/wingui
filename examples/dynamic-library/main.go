package main

import (
	"log"
	"syscall"
)

// May be cause error use Go as main program.
// details: https://github.com/golang/go/issues/22192
// https://github.com/golang/go/issues/34168
func main() {
	loaddll()
}

func loaddll() {
	//cw, _ := os.Getwd()
	h, err := syscall.LoadLibrary(`dll/test.dll`)
	log.Println("load dll", h, err)
	proc, err := syscall.GetProcAddress(h, "Showui")
	log.Println("proc", proc, err)
	_, _, _ = syscall.Syscall(proc, 0, 0, 0, 0)
}
