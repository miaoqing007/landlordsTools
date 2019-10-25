// Code generated by gen_proto.sh.
// DO NOT EDIT!
package client_handler

import "app/misc/packet"
import . "app/types"

var Code = map[string]int16{
	"heart_beat_req":         2000, // 心跳包..
	"heart_beat_ack":         2001, // 心跳包回复
	"user_login_req":         2002, // 登陆
	"user_login_succeed_ack": 2003, // 登陆成功
	"user_login_faild_ack":   2004, // 登陆失败
}

var RCode = map[int16]string{
	2000: "heart_beat_req",         // 心跳包..
	2001: "heart_beat_ack",         // 心跳包回复
	2002: "user_login_req",         // 登陆
	2003: "user_login_succeed_ack", // 登陆成功
	2004: "user_login_faild_ack",   // 登陆失败
}

var Handlers map[int16]func(*Session, *packet.Packet) []byte

func init() {
	Handlers = map[int16]func(*Session, *packet.Packet) []byte{}
}
