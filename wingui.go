/**
wingui Golang GUI library
*/
package wingui

/*

void MessageLoop();
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
