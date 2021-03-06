// Code generated by gen_proto.sh.
// DO NOT EDIT!
package client_handler

import . "landlords/session"

var Code = map[string]int16{
	"heart_beat_req":              2000, // 心跳包
	"user_login_req":              2001, // 登陆
	"error_ack":                   2002, // 错误返回
	"start_game_req":              2003, // 开始游戏
	"join_room_req":               2004, // 进入房间
	"user_reg_req":                2005, // 新用户注册
	"user_data_req":               2006, // 拉取用户信息
	"out_of_the_card_req":         2007, // 出牌请求
	"cancel_match_req":            2008, // 取消匹配
	"cancel_match_success_ack":    2009, // 取消匹配成功
	"out_of_the_card_failed_ack":  2010, // 出牌失败
	"out_of_the_card_success_ack": 2011, // 出牌成功
	"login_failed_ack":            2012, // 登陆失败
	"register_name_req":           2013, // 注册名字请求
	"register_name_ack":           2014, // 注册名字回复
	"register_name_success_ack":   2015, // 注册名字成功
	"give_up_card_req":            2016, // 放弃出牌
	"game_over_ack":               2017, // 游戏结束
	"Grab_The_Landlord_req":       2018, // 抢地主
	"Grab_The_Landlord_ack":       2019, // 抢地主结果
	"reset_Cards_ack":             2020, // 重新发牌
	"send_chat_msg_req":           2021, // 发送聊天信息
	"receive_chat_msg_ack":        2022, // 收到聊天信息
}

var RCode = map[int16]string{
	2000: "heart_beat_req",              // 心跳包
	2001: "user_login_req",              // 登陆
	2002: "error_ack",                   // 错误返回
	2003: "start_game_req",              // 开始游戏
	2004: "join_room_req",               // 进入房间
	2005: "user_reg_req",                // 新用户注册
	2006: "user_data_req",               // 拉取用户信息
	2007: "out_of_the_card_req",         // 出牌请求
	2008: "cancel_match_req",            // 取消匹配
	2009: "cancel_match_success_ack",    // 取消匹配成功
	2010: "out_of_the_card_failed_ack",  // 出牌失败
	2011: "out_of_the_card_success_ack", // 出牌成功
	2012: "login_failed_ack",            // 登陆失败
	2013: "register_name_req",           // 注册名字请求
	2014: "register_name_ack",           // 注册名字回复
	2015: "register_name_success_ack",   // 注册名字成功
	2016: "give_up_card_req",            // 放弃出牌
	2017: "game_over_ack",               // 游戏结束
	2018: "Grab_The_Landlord_req",       // 抢地主
	2019: "Grab_The_Landlord_ack",       // 抢地主结果
	2020: "reset_Cards_ack",             // 重新发牌
	2021: "send_chat_msg_req",           // 发送聊天信息
	2022: "receive_chat_msg_ack",        // 收到聊天信息
}

var Handlers map[int16]func(*Session, []byte) (int16, interface{})

func init() {
	Handlers = map[int16]func(*Session, []byte) (int16, interface{}){
		2000: P_heart_beat_req,
		2001: P_user_login_req,
		2003: P_start_game_req,
		2004: P_join_room_req,
		2005: P_user_reg_req,
		2006: P_user_data_req,
		2007: P_out_of_the_card_req,
		2008: P_cancel_match_req,
		2013: P_register_name_req,
		2016: P_give_up_card_req,
		2018: P_Grab_The_Landlord_req,
		2021: P_send_chat_msg_req,
	}
}
