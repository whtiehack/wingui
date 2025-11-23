/*
*
wingui Golang GUI library

# Usage

### Simple usage:

Generate x.syso file from rc or res file use `windres.exe` tool.

genereate syso:

`windres -i emptyProject/Debug/resource.res -O coff -o vsui.syso`

or

`windres -i ui/ui.rc -O coff -o ui.syso`

main.go
```go
package main

import "github.com/whtiehack/wingui"

	func main() {
		dlg, _ := wingui.NewDialog(101, 0, nil)
		dlg.SetIcon(105)
		btnok, _ := wingui.BindNewButton(1002, dlg)
		btncancel, _ := wingui.BindNewButton(1003, dlg)
		btnok.OnClicked = func() {
			dlg.Close()
		}
		btncancel.OnClicked = btnok.OnClicked
		dlg.Show()
		// This invoke is optional.
		wingui.SetCurrentDialog(dlg.Handle())
		wingui.MessageLoop()
	}

```

run:
`go run .`

Don't use `go run main.go`, because golang can't load x.syso files.
*/
package wingui

/*
#include <windows.h>

HWND dlg;

// Message Loop
void MessageLoop(){
    MSG msg;
    while (GetMessage(&msg, NULL, 0, 0))
    {
        if(dlg){
                if(!IsDialogMessage(dlg, &msg))
                {
                    TranslateMessage(&msg);
                    DispatchMessage(&msg);
                }
        }else{
            TranslateMessage(&msg);
            DispatchMessage(&msg);
        }
    }

}

void SetCurrentDialog(long long int h){
    dlg = (HWND)h;
}
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

// InitHInstance init hInstance,used by Dialog APIs.
func InitHInstance(lpModuleName string) {
	var name *uint16
	if lpModuleName != "" {
		name = syscall.StringToUTF16Ptr(lpModuleName)
	}
	hInstance = win.GetModuleHandle(name)
	log.Println("hInstance", hInstance)
}

// MessageLoop start windows message loop.
func MessageLoop() {
	C.MessageLoop()
}

// SetCurrentDialog  make sure Message Loop could process dialog msg correct,such as Tabstop msg.
// This is a optional method.
func SetCurrentDialog(h win.HWND) {
	C.SetCurrentDialog(C.longlong(uintptr(h)))
}
