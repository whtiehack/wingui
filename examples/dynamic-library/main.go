package main

import (
	"log"
	"syscall"
)

func main() {
	loaddll()
}

func loaddll() {
	//cw, _ := os.Getwd()
	h, err := syscall.LoadLibrary(`F:\gotest\wingui\examples\dynamic-library\dll\test.dll`)
	log.Println("load dll", h, err)
	proc, err := syscall.GetProcAddress(h, "Showui")
	log.Println("proc", proc, err)
	_, _, _ = syscall.Syscall(proc, 0, 0, 0, 0)
}
