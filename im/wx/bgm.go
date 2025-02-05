package wx

import (
	"github.com/cdle/sillyGirl/develop/core"
)

var wx = core.MakeBucket("wx")
var api_url = func() string {
	if v := wx.GetString("vlw_addr"); v != "" {
		return v
	}
	if v := wx.GetString("api_url"); v != "" {
		return v
	}
	return ""
}
var robot_wxid = wx.GetString("robot_wxid")

type TextMsg struct {
	Event      string `json:"event"`
	ToWxid     string `json:"to_wxid"`
	Msg        string `json:"msg"`
	RobotWxid  string `json:"robot_wxid"`
	GroupWxid  string `json:"group_wxid"`
	MemberWxid string `json:"member_wxid"`
}

type OtherMsg struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Event      string `json:"event"`
	RobotWxid  string `json:"robot_wxid"`
	ToWxid     string `json:"to_wxid"`
	MemberWxid string `json:"member_wxid"`
	MemberName string `json:"member_name"`
	GroupWxid  string `json:"group_wxid"`
	Msg        Msg    `json:"msg"`
}

type Msg struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type wxmsg struct {
	content   string
	user_id   string
	chat_id   int
	user_name string
	chat_name string
}

type Sender struct {
	leixing int
	mtype   int
	deleted bool
	value   wxmsg
	core.BaseSender
}

type JsonMsg struct {
	Event         string      `json:"event"`
	RobotWxid     string      `json:"robot_wxid"`
	RobotName     string      `json:"robot_name"`
	Type          int         `json:"type"`
	FromWxid      string      `json:"from_wxid"`
	FromName      string      `json:"from_name"`
	FinalFromWxid string      `json:"final_from_wxid"`
	FinalFromName string      `json:"final_from_name"`
	ToWxid        string      `json:"to_wxid"`
	Msg           interface{} `json:"msg"`
}
