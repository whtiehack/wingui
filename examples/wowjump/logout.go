package main

import (
	"github.com/whtiehack/wingui/winapi"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/lxn/win"
)

// LogoutStatus logout status.
type LogoutStatus int

const (
	//NORMAL 正常
	NORMAL LogoutStatus = iota
	//INPUT inputing
	INPUT
	//CHAR char
	CHAR
	//ENTERING entering
	ENTERING
)

// Logout control Wow window status.
type Logout struct {
	prevTime     time.Time
	subTime      time.Duration
	currentState LogoutStatus
	hwnd         win.HWND
	count        int
}

// Update should be invoke by period.
func (l *Logout) Update() {
	switch l.currentState {
	case NORMAL:
		l.normal()
	case INPUT:
		l.input()
	case CHAR:
		l.char()
	case ENTERING:
		l.enter()
	}
}

func (l *Logout) normal() {
	if l.subTime == 0 {
		l.subTime = time.Duration(rand.Int()%(config.NormalTime)*int(time.Second)) + time.Duration(config.NormalTime*int(time.Second))
		l.prevTime = time.Now()
		log.Println(l.hwnd, "挂机一会", l.subTime)
		l.count++
	}
	if time.Now().Sub(l.prevTime) > l.subTime {
		l.subTime = 0
		l.currentState = INPUT
	}
}

func (l *Logout) checkFlashWindow() {
	if !config.LogoutFlash {
		return
	}
	log.Println("准备小退。。闪烁窗口..")
	for i := 0; i < 10; i++ {
		winapi.FlashWindow(config.flashHwnd, i%1+1)
		sleep(500)
	}
	randomSleep(3333, 6666)
}
func (l *Logout) input() {
	if l.subTime == 0 {
		l.checkFlashWindow()
		if !l.TryGetWindow() {
			randomSleep(1111, 1111)
		}
		l.subTime = time.Duration(rand.Int()%(config.InputTime)*int(time.Second)) + time.Duration(config.InputTime*int(time.Second))
		l.prevTime = time.Now()
		log.Println(l.hwnd, "小退中。。。", l.subTime)
		if !l.logout() {
			l.subTime = 0
		} else {
			stat.Stat("/logout/"+strconv.Itoa(l.count), "wowjump-logout")
		}
		return
	}
	if time.Now().Sub(l.prevTime) > l.subTime {
		l.subTime = 0
		l.currentState = CHAR
	}
}

func (l *Logout) char() {
	if l.subTime == 0 {
		if config.ChangeChar {
			if !l.TryGetWindow() {
				randomSleep(1111, 1111)
			}
			if !l.inputKey(win.VK_DOWN) {
				log.Println(l, l.hwnd, "角色界面换角色失败。。等一会重试。。。")
				return
			}
		}
		l.subTime = time.Duration(rand.Int()%(config.CharWaitTime)*int(time.Second)) + time.Duration(config.CharWaitTime*int(time.Second))
		l.prevTime = time.Now()
		log.Println(l.hwnd, "角色界面等待一会。。。", l.subTime)
	}
	if time.Now().Sub(l.prevTime) > l.subTime {
		l.subTime = 0
		l.currentState = ENTERING
	}
}
func (l *Logout) enter() {
	if l.subTime == 0 {
		if !l.TryGetWindow() {
			randomSleep(1111, 1111)
		}
		l.subTime = time.Duration(rand.Int()%(config.EnterTime)*int(time.Second)) + time.Duration(config.EnterTime*int(time.Second))
		l.prevTime = time.Now()
		log.Println(l.hwnd, "正在进入游戏。。。", l.subTime)
		if !l.inputEnter() {
			l.subTime = 0
		}
		return
	}
	if time.Now().Sub(l.prevTime) > l.subTime {
		l.subTime = 0
		l.currentState = NORMAL
	}
}

// Reset should be invoke by reset status.
func (l *Logout) Reset() {
	l.currentState = NORMAL
	l.subTime = 0
	l.prevTime = time.Now()
}

func (l *Logout) inputEnter() bool {
	if !l.CheckWindow() {
		return false
	}
	keybdInput(uint('\r'), false)
	randomSleep(211, 155)
	keybdInput(uint('\r'), true)
	return true
}

func (l *Logout) inputKey(key uint) bool {
	if !l.CheckWindow() {
		return false
	}
	keybdInput(key, false)
	randomSleep(211, 155)
	keybdInput(key, true)
	return true
}

func (l *Logout) logout() bool {
	if !l.CheckWindow() {
		return false
	}
	keybdInput(uint('\r'), false)
	randomSleep(211, 155)
	keybdInput(uint('\r'), true)
	randomSleep(211, 155)
	sendString("/logout", 333)
	randomSleep(211, 155)
	keybdInput(uint('\r'), false)
	randomSleep(111, 55)
	keybdInput(uint('\r'), true)
	return true
}

// IsValid check the windows is valid.
func (l *Logout) IsValid() bool {
	return win.IsWindowVisible(l.hwnd)
}

// CheckWindow check the window if in the foreground and in focus.
func (l *Logout) CheckWindow() bool {
	// 检查窗口
	if !l.IsValid() {
		return false
	}
	focus := win.GetForegroundWindow()
	if l.hwnd != focus {
		//log.Printf("FOCUS not WOW  focus:%d  WOW:%d  focus WOW", focus, logout.hwnd)
		win.ShowWindow(l.hwnd, win.SW_NORMAL)
		win.SetActiveWindow(l.hwnd)
		win.BringWindowToTop(l.hwnd)
		win.SetForegroundWindow(l.hwnd)
		randomSleep(2000, 1000)
	}
	return true
}

// TryGetWindow try focus window.
func (l *Logout) TryGetWindow() bool {
	if !l.IsValid() {
		return false
	}
	focus := win.GetForegroundWindow()
	if focus != l.hwnd {
		win.ShowWindow(l.hwnd, win.SW_NORMAL)
		randomSleep(500, 1000)
		win.BringWindowToTop(l.hwnd)
		randomSleep(500, 1000)
		winapi.SwitchToWindow(l.hwnd, 1)
		randomSleep(500, 1000)
		focus = win.GetForegroundWindow()
		if focus != l.hwnd {
			log.Println("当前窗口不是wow，小退过程中不要动游戏窗口", focus, l.hwnd)
			randomSleep(500, 1000)
			hwnd := config.flashHwnd
			if hwnd == 0 {
				hwnd = focus
			}
			winapi.SwitchToWindow(hwnd, 1)
			randomSleep(500, 1000)
			return l.TryGetWindow()
		}
		return false
	}
	return true
}
