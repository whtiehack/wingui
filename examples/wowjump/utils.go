package main

import (
	"log"
	"math/rand"
	"time"
	"unsafe"

	"github.com/lxn/win"
)

var gScrx int32

var gScry int32

func init() {
	gScrx = win.GetSystemMetrics(win.SM_CXSCREEN)
	gScry = win.GetSystemMetrics(win.SM_CYSCREEN)
}
func sleep(t int) {
	time.Sleep(time.Duration(t) * time.Millisecond)
}

// low+randome(t) sleep Millisecond
func randomSleep(t int, low int) {
	sleep(rand.Int()%t + low)
}

func keybdInput(key uint, isUp bool) {
	//key = w32.MapVirtualKey(key,1)
	input := []win.KEYBD_INPUT{
		{
			Type: win.INPUT_KEYBOARD,
			Ki: win.KEYBDINPUT{
				WVk: uint16(key),
			},
		},
	}
	if isUp {
		input[0].Ki.DwFlags = win.KEYEVENTF_KEYUP
	}
	//log.Printf("SendInput(%d, %#v, %d)\n", len(input), input, unsafe.Sizeof(input[0]))
	i := win.SendInput(uint32(len(input)), unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))
	if i != 1 {
		log.Printf("keybd_input => %d (lastError: %d)\n", i, win.GetLastError())
	}
}

func sendString(str string, delay int) {
	//key = w32.MapVirtualKey(key,1)
	for _, v := range str {
		input := []win.KEYBD_INPUT{
			{
				Type: win.INPUT_KEYBOARD,
				Ki: win.KEYBDINPUT{
					WVk:     0,
					DwFlags: win.KEYEVENTF_UNICODE,
					WScan:   uint16(v),
				},
			},
		}
		win.SendInput(uint32(len(input)), unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))
		input[0].Ki.DwFlags = win.KEYEVENTF_KEYUP
		randomSleep(delay, 55)
		win.SendInput(uint32(len(input)), unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))
	}
}

func mouseMove(x, y uint, times uint32, dwflag uint32) {
	input := []win.MOUSE_INPUT{
		{
			Type: win.INPUT_MOUSE,
			Mi: win.MOUSEINPUT{
				Dx:          int32(65535 * int(x) / int(gScrx)),
				Dy:          int32(65535 * int(y) / int(gScry)),
				DwFlags:     dwflag,
				MouseData:   0,
				DwExtraInfo: 0,
				Time:        times,
			},
		},
	}
	i := win.SendInput(uint32(len(input)), unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))
	if i != 1 {
		log.Printf("mouse_move => %d (lastError: %d)\n", i, win.GetLastError())
	}
}

var frame int32
var caption int32

func init() {
	frame = win.GetSystemMetrics(win.SM_CXFRAME)
	caption = win.GetSystemMetrics(win.SM_CYCAPTION)
}
func randomMoveMouse() {
	//h := win.FindWindow(syscall.StringToUTF16Ptr("GxWindowClass"), nil)
	var rect win.RECT
	//	win.GetWindowRect(h, &rect)
	//log.Printf("边框宽: %d; 标题高: %d  rect:%#v", frame, caption, rect);
	rect.Right -= frame
	rect.Left += frame
	rect.Top += caption
	rect.Bottom -= caption
	x := uint(rect.Left + rand.Int31()%(rect.Right-rect.Left))
	y := uint(rect.Top + rand.Int31()%(rect.Bottom-rect.Top))
	//log.Printf("!! x y", x, y)
	mouseMove(x, y, 0, win.MOUSEEVENTF_ABSOLUTE|win.MOUSEEVENTF_MOVE)
	//mouse_move(x, y, 0, win.MOUSEEVENTF_LEFTDOWN)
	//mouse_move(x, y, 0, win.MOUSEEVENTF_LEFTUP)

}
