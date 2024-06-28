package main

import (
	"encoding/json"
	"github.com/lxn/win"
	"github.com/whtiehack/wingui"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"syscall"
)

// Config is program config.
type Config struct {
	NormalTime          int
	InputTime           int
	CharWaitTime        int
	EnterTime           int
	ChangeChar          bool
	LogoutFlash         bool
	editEnterTime       *wingui.Edit   `json:"-"`
	editNormaltime      *wingui.Edit   `json:"-"`
	editInputTime       *wingui.Edit   `json:"-"`
	editCharWaitTime    *wingui.Edit   `json:"-"`
	btnCheckChangeChar  *wingui.Button `json:"-"`
	btnCheckLogoutFlash *wingui.Button `json:"-"`
	flashHwnd           win.HWND       `json:"-"`
}

var config = &Config{
	NormalTime:   800,
	InputTime:    30,
	CharWaitTime: 60,
	EnterTime:    30,
	ChangeChar:   false,
}

func init() {
	configDir := os.Getenv("AppData")
	log.Println("config", configDir)
	file, err := ioutil.ReadFile(configDir + "/wowjump/wowjump_config.txt")
	if err != nil {
		return
	}
	err = json.Unmarshal(file, config)
	if err != nil {
		log.Println("解析配置文件出错。")
		win.MessageBox(0, syscall.StringToUTF16Ptr("解析配置文件出错："+err.Error()), nil, 0)
		return
	}
	log.Println("解析配置文件成功", string(file))
}

// Save save config to file.
func (c *Config) Save() {
	config.EnterTime, _ = strconv.Atoi(c.editEnterTime.Text())
	config.NormalTime, _ = strconv.Atoi(c.editNormaltime.Text())
	config.InputTime, _ = strconv.Atoi(c.editInputTime.Text())
	config.CharWaitTime, _ = strconv.Atoi(c.editCharWaitTime.Text())
	config.ChangeChar = c.btnCheckChangeChar.GetCheck() == win.BST_CHECKED
	config.LogoutFlash = c.btnCheckLogoutFlash.GetCheck() == win.BST_CHECKED
	file, _ := json.MarshalIndent(c, "", "    ")
	configDir := os.Getenv("AppData")
	os.Mkdir(configDir+"/wowjump", 0777)
	ioutil.WriteFile(configDir+"/wowjump/wowjump_config.txt", file, 0777)
}

// InitVal put config val to edit.
func (c *Config) InitVal() {
	c.editEnterTime.SetText(strconv.Itoa(c.EnterTime))
	c.editNormaltime.SetText(strconv.Itoa(c.NormalTime))
	c.editInputTime.SetText(strconv.Itoa(c.InputTime))
	c.editCharWaitTime.SetText(strconv.Itoa(c.CharWaitTime))
	state := win.BST_UNCHECKED
	if config.ChangeChar {
		state = win.BST_CHECKED
	}
	c.btnCheckChangeChar.SetCheck(state)
	state = win.BST_UNCHECKED
	if config.LogoutFlash {
		state = win.BST_CHECKED
	}
	c.btnCheckLogoutFlash.SetCheck(state)
}

// EditEnable enable edits.
func (c *Config) EditEnable(enable bool) {
	c.editEnterTime.SetEnabled(enable)
	c.editNormaltime.SetEnabled(enable)
	c.editInputTime.SetEnabled(enable)
	c.editCharWaitTime.SetEnabled(enable)
	c.btnCheckChangeChar.SetEnabled(enable)
	c.btnCheckLogoutFlash.SetEnabled(enable)
}
