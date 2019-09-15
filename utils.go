package wingui

import (
	"errors"
	"fmt"
	"github.com/lxn/win"
)

func RGB(r, g, b uintptr) win.COLORREF {
	return win.COLORREF(r | g<<8 | b<<16)
}

func lastError(win32FuncName string) error {
	if errno := win.GetLastError(); errno != win.ERROR_SUCCESS {
		return errors.New(fmt.Sprintf("%s: Error %d", win32FuncName, errno))
	}
	return errors.New(win32FuncName)
}
