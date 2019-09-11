package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"syscall"

	"github.com/lxn/win"
	"github.com/whtiehack/wingui"
)

type Config struct {
	NormalTime       int
	InputTime        int
	CharWaitTime     int
	EnterTime        int
	editEnterTime    *wingui.Edit
	editNormaltime   *wingui.Edit
	editInputTime    *wingui.Edit
	editCharWaitTime *wingui.Edit
}

var config = &Config{
	NormalTime:   800,
	InputTime:    30,
	CharWaitTime: 60,
	EnterTime:    30,
}

func init() {
	configDir, _ := os.UserConfigDir()
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

//Save save config to file.
func (c *Config) Save() {
	config.EnterTime, _ = strconv.Atoi(c.editEnterTime.Text())
	config.NormalTime, _ = strconv.Atoi(c.editNormaltime.Text())
	config.InputTime, _ = strconv.Atoi(c.editInputTime.Text())
	config.CharWaitTime, _ = strconv.Atoi(c.editCharWaitTime.Text())
	file, _ := json.MarshalIndent(c, "", "    ")
	configDir, _ := os.UserConfigDir()
	ioutil.WriteFile(configDir+"/wowjump/wowjump_config.txt", file, 777)
}

//InitVal put config val to edit.
func (c *Config) InitVal() {
	c.editEnterTime.SetText(strconv.Itoa(c.EnterTime))
	c.editNormaltime.SetText(strconv.Itoa(c.NormalTime))
	c.editInputTime.SetText(strconv.Itoa(c.InputTime))
	c.editCharWaitTime.SetText(strconv.Itoa(c.CharWaitTime))
}

//EditEnable enable edits.
func (c *Config) EditEnable(enable bool) {
	c.editEnterTime.SetEnabled(enable)
	c.editNormaltime.SetEnabled(enable)
	c.editInputTime.SetEnabled(enable)
	c.editCharWaitTime.SetEnabled(enable)
}
