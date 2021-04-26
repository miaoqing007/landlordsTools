// Code generated by gen_proto.sh.
// DO NOT EDIT!
using UnityEngine;
using System;
using System.Collections.Generic;

namespace NetProto.Api
{
    public enum ENetMsgId
    { 
        heart_beat_req = 2000, // 心跳包
        user_login_req = 2001, // 登陆
        error_ack = 2002, // 错误返回
        start_game_req = 2003, // 开始游戏
        join_room_req = 2004, // 进入房间
        user_reg_req = 2005, // 新用户注册
        user_data_req = 2006, // 拉取用户信息
        out_of_the_card_req = 2007, // 出牌请求
        cancel_match_req = 2008, // 取消匹配
        cancel_match_success_ack = 2009, // 取消匹配成功
        out_of_the_card_failed_ack = 2010, // 出牌失败
        out_of_the_card_success_ack = 2011, // 出牌成功
        login_failed_ack = 2012, // 登陆失败
        register_name_req = 2013, // 注册名字请求
        register_name_ack = 2014, // 注册名字回复
        register_name_success_ack = 2015, // 注册名字成功
        give_up_card_req = 2016, // 放弃出牌
        game_over_ack = 2017, // 游戏结束
    };
}
