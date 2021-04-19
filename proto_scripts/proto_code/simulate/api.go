// Code generated by gen_proto.sh.
// DO NOT EDIT!
package main

var Code = map[string]int16{
	"heart_beat_req":      2000, // 心跳包
	"user_login_req":      2001, // 登陆
	"error_ack":           2002, // 错误返回
	"start_game_req":      2003, // 开始游戏
	"join_room_req":       2004, // 进入房间
	"user_reg_req":        2005, // 新用户注册
	"user_data_req":       2006, // 拉取用户信息
	"out_of_the_card_req": 2007, // 出牌请求
}

var RCode = map[int16]string{
	2000: "heart_beat_req",      // 心跳包
	2001: "user_login_req",      // 登陆
	2002: "error_ack",           // 错误返回
	2003: "start_game_req",      // 开始游戏
	2004: "join_room_req",       // 进入房间
	2005: "user_reg_req",        // 新用户注册
	2006: "user_data_req",       // 拉取用户信息
	2007: "out_of_the_card_req", // 出牌请求
}
