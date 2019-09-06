package wingui

/*

void MessageLoop();
*/
import "C"
import "github.com/lxn/win"

var hInstance win.HINSTANCE

func init() {
	hInstance = win.GetModuleHandle(nil)
}

// Message loop
func MessageLoop() {
	C.MessageLoop()
}
