package main

import (
	"github.com/whtiehack/wingui/winapi"
	"log"
	"math/rand"
	"net"
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
	// tt   title
	// ep	停留时间,活跃时间
	// et	初始进入页面 0， 当前页面有时间 3
	// rnd  随机数
	// su  前一个页面  离开页面时
	// ct 新的一个页面开始   !!
	val, _ := url.ParseQuery(s.params.Encode())
	val.Set("rnd", strconv.Itoa(int(rand.Int31())))
	val.Set("sn", strconv.Itoa(int(time.Now().Unix()%65535)))
	if s.lt != 0 && int(time.Now().Unix())-s.lt > 2592E3 {
		s.lt = int(time.Now().Unix())
	}
	if s.lt != 0 {
		val.Set("lt", strconv.Itoa(s.lt))
	}
	if s.curPath == "" {
		// 第一次打开
		val.Set("ct", "!!")
		val.Set("tt", title)
		val.Set("et", "0")
		s.get(val.Encode(), s.baseUrl+path)
	} else {
		// 离开页面
		val.Set("et", "3")
		t := time.Now().Sub(s.prevTime).Milliseconds()
		ep := strconv.Itoa(int(t)) + "," + strconv.Itoa(int(t))
		val.Set("ep", ep)
		prevPath := s.baseUrl + s.curPath
		s.get(val.Encode(), prevPath)
		val.Set("rnd", strconv.Itoa(int(rand.Int31())))
		// 进入新页面
		val.Set("u", prevPath)
		s.get(val.Encode(), s.baseUrl+path)
		val.Set("rnd", strconv.Itoa(int(rand.Int31())))
		val.Del("u")
		val.Set("su", prevPath)
		val.Set("tt", title)
		val.Del("ep")
		val.Set("et", "0")
		val.Set("ct", "!!")

		if int(time.Now().Unix())-s.lt > 2592E3 {
			s.lt = int(time.Now().Unix())
		}
		val.Set("lt", strconv.Itoa(s.lt))
		s.get(val.Encode(), s.baseUrl+path)
	}
	s.prevTime = time.Now()
	s.curPath = path
}

func (s *Statistics) get(params string, referer string) error {
	conn, err := net.Dial("tcp", "hm.baidu.com:80")
	if err != nil {
		print("what err:", err, "\n")
		return err
	}
	defer conn.Close()
	_, err = conn.Write([]byte("GET /hm.gif?" + params + " HTTP/1.1\r\n" +
		"User-Agent: Wingui\r\n" +
		"Referer: " + referer + "\r\n" +
		"Host: hm.baidu.com\r\n" +
		"Connection: close\r\n" +
		"\r\n\r\n"))
	if err != nil {
		print("write err:", err, "\n")
		return err
	}
	return nil
}
