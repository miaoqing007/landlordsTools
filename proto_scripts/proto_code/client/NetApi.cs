// Code generated by gen_proto.sh.
// DO NOT EDIT!
using UnityEngine;
using System;
using System.Collections.Generic;

namespace NetProto.Api
{
    public enum ENetMsgId
    { 
        heart_beat_req = 2000, // 心跳包..
        heart_beat_ack = 2001, // 心跳包回复
        user_login_req = 2002, // 登陆
        user_login_succeed_ack = 2003, // 登陆成功
        user_login_faild_ack = 2004, // 登陆失败
    };
}
