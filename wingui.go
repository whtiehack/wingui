/**
wingui Golang GUI library
*/
package wingui

/*
void MessageLoop();
void SetCurrentDialog(long long int h);
*/
import "C"
import (
	"github.com/lxn/win"
	"log"
	"os"
	"syscall"
)

var hInstance win.HINSTANCE

func init() {
	log.SetOutput(os.Stdout)
	InitHInstance("")
}

func InitHInstance(lpModuleName string) {
	var name *uint16
	if lpModuleName != "" {
		name = syscall.StringToUTF16Ptr(lpModuleName)
	}
	hInstance = win.GetModuleHandle(name)
	log.Println("hInstance", hInstance)
}

// Message loop
func MessageLoop() {
	C.MessageLoop()
}

// SetCurrentDialog  make sure Message Loop could process dialog msg correct
func SetCurrentDialog(h win.HWND) {
	C.SetCurrentDialog(C.longlong(uintptr(h)))
}
