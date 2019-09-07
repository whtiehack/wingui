package wingui

/*

void MessageLoop();
*/
import "C"
import (
	"github.com/lxn/win"
	"log"
	"os"
)

var hInstance win.HINSTANCE

func init() {
	log.SetOutput(os.Stdout)
	hInstance = win.GetModuleHandle(nil)
}

// Message loop
func MessageLoop() {
	C.MessageLoop()
}
