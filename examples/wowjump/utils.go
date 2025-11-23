package main

import (
	"github.com/whtiehack/wingui/winapi"
	"log"
	"math/rand"
	"net/url"
	"strconv"
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

type Statistics struct {
	prevTime time.Time
	si       string
	curPath  string
	params   url.Values
	baseUrl  string
	lt       int
}

func NewStatistics(baseUrl string, si string) *Statistics {
	params := url.Values{}
	params.Add("cc", "1")
	params.Add("ck", "1")
	params.Add("cl", "24-bit")
	params.Add("ds", strconv.Itoa(int(win.GetSystemMetrics(win.SM_CXSCREEN)))+"x"+strconv.Itoa(int(win.GetSystemMetrics(win.SM_CYSCREEN))))
	params.Add("vl", "797")
	params.Add("et", "0")
	params.Add("ja", "0")
	params.Add("ln", "zh-cn")
	lang := winapi.GetSystemDefaultLocaleName()
	if lang != "" {
		params.Set("ln", lang)
	}
	params.Add("lo", "0")
	// params.Add("rnd", "0")
	params.Add("si", si)
	params.Add("v", "1.2.61")
	params.Add("lv", "2")
	// params.Add("sn", "30541")
	return &Statistics{baseUrl: baseUrl, lt: 0, si: si, params: params}
}

func (s *Statistics) Stat(path string, title string) {
}

func (s *Statistics) get(params string, referer string) error {
	return nil
}
