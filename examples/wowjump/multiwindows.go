package main

import (
	"github.com/whtiehack/wingui/winapi"
	"log"
	"strconv"
	"syscall"

	"github.com/lxn/win"
)

func getHwnds() []win.HWND {
	hwndArr := make([]win.HWND, 0, 10)
	var curh win.HWND
	strArr := make([]string, 0, 10)
	for {
		// syscall.StringToUTF16Ptr("GxWindowClass")
		//curh = winapi.FindWindowEx(0, curh, nil, syscall.StringToUTF16Ptr("ssl.csr - 记事本"))
		curh = winapi.FindWindowEx(0, curh, syscall.StringToUTF16Ptr("GxWindowClass"), nil)
		if curh == 0 {
			break
		}
		hwndArr = append(hwndArr, curh)
		strArr = append(strArr, strconv.Itoa(int(curh)))
	}
	log.Println("hwndArr", hwndArr)
	return hwndArr
}

//GetMultiLogout get need control WOW windows and need user confirm.
func GetMultiLogout() []*Logout {
	arr := getHwnds()
	var logouts []*Logout
	for _, hwnd := range arr {
		win.ShowWindow(hwnd, win.SW_NORMAL)
		win.SetActiveWindow(hwnd)
		win.SetForegroundWindow(hwnd)
		c := make(chan struct{}, 5)
		r := win.MessageBox(hwnd, syscall.StringToUTF16Ptr("是否要操作这个WoW窗口?"),
			syscall.StringToUTF16Ptr("确认"), win.MB_YESNO)
		log.Println("select ret", r)
		c <- struct{}{}
		if r != win.IDYES {
			continue
		}
		logouts = append(logouts, &Logout{hwnd: hwnd})
	}
	return logouts
}
