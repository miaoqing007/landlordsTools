package main

import (
	"crypto/rc4"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/big"
	"misc/crypto/dh"
	"misc/packet"
	"net"
	"os"
	"strings"
	"time"
)

var (
	seqid        = uint32(0)
	encoder      *rc4.Cipher
	decoder      *rc4.Cipher
	KEY_EXCHANGE = false
	SALT         = "DH"
	isNew        = true
	atuoReg      = false
)

const (
	DEFAULT_AGENT_HOST = "127.0.0.1:8888"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic("error occured in protocol module")
	}
}
func main() {
	host := DEFAULT_AGENT_HOST
	if env := os.Getenv("AGENT_HOST"); env != "" {
		host = env
	}
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	defer conn.Close()

	//get_seed_req
	S1, M1 := dh.DHExchange()
	S2, M2 := dh.DHExchange()
	p2 := seed_info{
		int32(M1.Int64()),
		int32(M2.Int64()),
	}
	rst := send_proto(conn, Code["get_seed_req"], p2)
	r1, _ := PKT_seed_info(rst)
	log.Printf("result: %#v", r1)

	K1 := dh.DHKey(S1, big.NewInt(int64(r1.F_client_send_seed)))
	K2 := dh.DHKey(S2, big.NewInt(int64(r1.F_client_receive_seed)))
	encoder, err = rc4.NewCipher([]byte(fmt.Sprintf("%v%v", SALT, K1)))
	if err != nil {
		log.Println(err)
		return
	}
	decoder, err = rc4.NewCipher([]byte(fmt.Sprintf("%v%v", SALT, K2)))
	if err != nil {
		log.Println(err)
		return
	}

	KEY_EXCHANGE = true

	username := "vgxbukz"
	password := "111111"

	//自动注册一个账号
	if atuoReg {
		reader := send_proto(conn, Code["auto_reg_req"], null_struct{})
		if tbl, err := PKT_auto_reg_data(reader); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println(tbl)
			username = tbl.F_userName
			log.Printf("username is: %s\n", username)
			password = tbl.F_userPass
			log.Printf("password is: %s\n", password)
		}
	}

	//user_login_req
	p3 := user_login_info{
		F_userName: username,
		F_userPass: password,
		F_channel:  "CCjoy",
		F_token:    "123",
		F_language: "",
	}
	send_proto(conn, Code["user_login_req"], p3)

	//测试匹配
	send_proto(conn, Code["pvp_match_req"], byte_id{1})

	header := make([]byte, 2)
	if _, err := io.ReadFull(conn, header); err != nil {
		log.Println(err)
		if err == io.EOF {
			os.Exit(1)
		}
	}
	//log.Printf("recv")
	size := binary.BigEndian.Uint16(header)
	r := make([]byte, size)
	_, err = io.ReadFull(conn, r)
	if err != nil {
		log.Println(err)
	}
	if KEY_EXCHANGE {
		decoder.XORKeyStream(r, r)
	}
	reader := packet.Reader(r)
	b, err := reader.ReadS16()
	if err != nil {
		log.Println(err)
	}
	if s, ok := RCode[b]; !ok {
		log.Println("unknown proto ", b)
	} else {

		var tbl interface{}

		if s == "pvp_match_success_notify" {
			tbl, _ = PKT_match_info(reader)
			log.Printf("%s=>%+v\n", s, tbl)
		} else {
			log.Println(s)
		}
	}

	select {}

	//获取用户数据
	send_proto(conn, Code["user_data_req"], null_struct{})
	//send_proto2(conn, Code["user_data_req"], null_struct{})

	//注册用户
	if isNew {
		reg := user_reg_info{
			F_nickname: "lksjz72",
		}
		send_proto(conn, Code["user_reg_req"], reg)
		//p66 := construction_build_info{
		//F_tid: 10200,
		//F_x:   20,
		//F_y:   7,
		//}
		//send_proto(conn, Code["construction_build_req"], p66)

	} //else {

	p66 := null_struct{}
	//send_proto2(conn, Code["cheat_req"], p66)

	send_proto(conn, Code["mail_getlist_req"], p66)
	p67 := auto_id{1}
	//send_proto(conn, Code["mail_read_req"], p67)
	send_proto(conn, Code["mail_handle_req"], p67)

	os.Exit(0)
	//建造
	p6 := construction_build_info{
		F_tid: 10200,
		F_x:   18,
		F_y:   7,
	}
	//send_proto(conn, Code["construction_build_req"], p6)
	reader = send_proto(conn, Code["construction_build_req"], p6)

	////取消建造
	if tbl, err := PKT_construction_succeed_info(reader); err == nil {
		log.Println(tbl)
		log.Println("开始测试取消建设 id => ", tbl.F_construction.F_constructionId)
		p10 := construction_id_info{
			F_id: tbl.F_construction.F_constructionId,
		}
		reader = send_proto(conn, Code["construction_cancel_req"], p10)
		if tbl, err := PKT_construction_destroy_info(reader); err == nil {

			log.Println("取消建造成功 id => ", tbl.F_id)
		} else {

			log.Println("取消建造失败")
		}

		reader = send_proto(conn, Code["construction_build_req"], p6)
		tbl, _ := PKT_construction_succeed_info(reader)
		p11 := construction_id_info{
			F_id: tbl.F_construction.F_constructionId,
		}

		////立即完成
		reader = send_proto(conn, Code["construction_clearcd_req"], p11)
		if tbl, err := PKT_construction_succeed_info(reader); err == nil {
			log.Printf("%s=>%v\n", tbl)
			log.Println("立即完成建设成功 id => ", tbl.F_construction.F_constructionId)
		} else {

			log.Println("立即完成失败")
		}

	}

	////移动 注意，如果不改变坐标，服务器将不返回数据
	//p7 := construction_move_info{
	//F_id: 6161456830020063232,
	//F_x:  1,
	//F_y:  1,
	//}

	//send_proto(conn, Code["construction_move_req"], p7)

	////升级基地
	//p8 := construction_id_info{
	//F_id: 6161456830020063232,
	//}

	//send_proto(conn, Code["construction_upgrade_req"], p8)
	////取消升级
	//send_proto(conn, Code["construction_cancel_upgrade_req"], p8)
	////再升级一次
	//send_proto(conn, Code["construction_upgrade_req"], p8)
	//// 立即完成
	//send_proto(conn, Code["construction_clearcd_req"], p8)
	////收获
	//p9 := construction_id_info{
	//F_id: 6166186876391657472,
	//}
	//send_proto(conn, Code["construction_harvest_req"], p9)

	//os.Exit(0)
	//强化请求
	//p10 := warship_evolution_info{
	//F_id:   6166550180318416896,
	//F_type: 4,
	//}
	//send_proto(conn, Code["warship_evolution_req"], p10)
	//改造请求
	//p11 := entity_id{
	//	F_id: 0x55dd2a38c0801000,
	//}
	//send_proto(conn, Code["warship_awaken_req"], p11)
	////维修请求
	//p12 := entity_id{
	//F_id: 6166550180318416896,
	//}
	//send_proto(conn, Code["warship_repair_req"], p12)

	//os.Exit(0)
	//send_proto(conn, Code["classroom_info_req"], null_struct{})
	//进入教室请求
	//p13 := room_action{
	//F_id:  6166550180318416896,
	//F_pos: 2,
	//}
	//send_proto(conn, Code["classroom_in_req"], p13)
	//os.Exit(0)
	//退出教室请求
	//p14 := entity_id{
	//F_id: 6166550180318416896,
	//}
	//send_proto(conn, Code["classroom_out_req"], p14)
	//p15 := auto_id{
	//F_id: 2,
	//}
	//send_proto(conn, Code["create_curriculums_req"], p15)
	//p16 := room_curriculum_ids{
	//F_curriculumpos: 5,
	//F_roompos:       2,
	//}
	//send_proto(conn, Code["set_curriculum_req"], p16)
	////取消学习
	//p17 := auto_id{
	//F_id: 2,
	//}
	//send_proto(conn, Code["stop_curriculum_req"], p17)
	//加速课程
	//p18 := auto_id{
	//F_id: 2,
	//}
	//send_proto(conn, Code["clearcd_curriculum_req"], p18)
	//随便逛逛
	//p19 := auto_id{
	//	F_id: 3,
	//}
	//send_proto(conn, Code["encounter_req"], p19)
	//分支选择

	//p20 := auto_id{
	//	F_id: 7,
	//}
	//send_proto(conn, Code["encounter_branch_req"], p20)

	//	p21 := accessories_fit{64, "pos1", 32}
	//	send_proto(conn, Code["accessories_fit_req"], p21)
	//
	//	p22 := accessories_remove{64, "pos1"}
	//	send_proto(conn, Code["accessories_remove_req"], p22)

	//}

	// 舰船建造
	//p23 := build_info{
	//	F_gold:     8001,
	//	F_fuel:     2001,
	//	F_material: 3801,
	//}
	//send_proto(conn, Code["ship_build_req"], p23)

	// 舰船拆解
	//p24 := entity_id{
	//	F_id: 0x55dd2a38c0801000,
	//}
	//send_proto(conn, Code["warship_disintegration_req"], p24)

	// 签到测试
	//p25 := null_struct{}
	//send_proto(conn, Code["user_check_in_req"], p25)

	// 舰船技能升级
	//p26 := ship_skill_info{
	//	F_shipId:  214312341,
	//	F_skillId: "10011",
	//}
	//send_proto(conn, Code["ship_skill_upgrade_req"], p26)
}

func send_proto(conn net.Conn, p int16, info interface{}) (reader *packet.Packet) {
	seqid++
	payload := packet.Pack(p, info, nil)
	writer := packet.Writer()
	writer.WriteU16(uint16(len(payload)) + 4)

	w := packet.Writer()
	w.WriteU32(seqid)
	w.WriteRawBytes(payload)
	data := w.Data()
	//log.Println("----%#v", data)

	if KEY_EXCHANGE {
		encoder.XORKeyStream(data, data)
	}
	writer.WriteRawBytes(data)
	conn.Write(writer.Data())
	//log.Printf("send : %#v", writer.Data())
	time.Sleep(time.Second)

	//read
	header := make([]byte, 2)
	if _, err := io.ReadFull(conn, header); err != nil {
		log.Println(err)
		if err == io.EOF {
			os.Exit(1)
		}
	}
	//log.Printf("recv")
	size := binary.BigEndian.Uint16(header)
	r := make([]byte, size)
	_, err := io.ReadFull(conn, r)
	if err != nil {
		log.Println(err)
	}
	if KEY_EXCHANGE {
		decoder.XORKeyStream(r, r)
	}
	reader = packet.Reader(r)
	b, err := reader.ReadS16()
	if err != nil {
		log.Println(err)
	}
	if s, ok := RCode[b]; !ok {
		log.Println("unknown proto ", b)
	} else {

		var tbl interface{}

		if strings.Index(s, "_faild_ack") > 0 || strings.Index(s, "_failed_ack") > 0 {
			tbl, _ = PKT_error_info(reader)
		}
		if s == "construction_upgrade_succeed_ack" || s == "construction_cancel_upgrade_succeed_ack" || s == "construction_harvest_succeed_ack" {
			tbl, _ = PKT_construction_succeed_info(reader)
		}
		if s == "user_data_succeed_ack" {
			tbl, _ = PKT_user_info(reader)
		}
		if s == "construction_clearcd_succeed_ack" {
			tbl, _ = PKT_construction_succeed_info(reader)
		}
		if s == "user_new_notify" {
			log.Println("新用户标示 ")
			isNew = true
		}
		if s == "mail_new_notify" {
			log.Println("新邮件提示 ")
		}

		if s == "warship_evolution_succeed_ack" || s == "warship_awaken_succeed_ack" || s == "warship_repair_succeed_ack" {
			tbl, _ = PKT_warship_change_succeed(reader)
		}

		if s == "classroom_info_succeed_ack" {
			tbl, _ = PKT_classroom_info(reader)
		}

		if s == "classroom_in_succeed_ack" || s == "classroom_out_succeed_ack" {
			tbl, _ = PKT_classroom_role(reader)
		}

		if s == "create_curriculums_succeed_ack" {
			tbl, _ = PKT_curriculums_info(reader)
		}

		if s == "set_curriculum_succeed_ack" || s == "stop_curriculum_succeed_ack" {
			tbl, _ = PKT_set_curriculum_info(reader)
		}

		if s == "clearcd_curriculum_succeed_ack" {
			tbl, _ = PKT_clearcd_curriculum_info(reader)
		}

		if s == "encounter_succeed_ack" {
			tbl, _ = PKT_encounter_info(reader)
		}

		if s == "accessories_fit_succeed_ack" {
			tbl, _ = PKT_warship_info(reader)
		}

		if s == "accessories_remove_succeed_ack" {
			tbl, _ = PKT_warship_info(reader)
		}

		if s == "ship_build_succeed_ack" {
			tbl, _ = PKT_build_info(reader)
		}

		if s == "warship_disintegration_success_ack" || s == "mail_handle_succeed_ack" {
			tbl, _ = PKT_item_change_info(reader)
		}

		if s == "user_check_in_success_ack" {
			tbl, _ = PKT_checkin_response(reader)
		}
		if s == "mail_getlist_succeed_ack" {
			tbl, _ = PKT_mail_list(reader)
		}
		if s == "msg_new_notify" {
			log.Println("新消息提示 ")
			tbl, _ = PKT_msg(reader)
		}

		log.Printf("%s=>%+v\n", s, tbl)
	}
	//log.Println("++++%#v", reader.Data())

	return
}

func send_proto2(conn net.Conn, p int16, info interface{}) {
	seqid++
	payload := packet.Pack(p, info, nil)
	writer := packet.Writer()
	writer.WriteU16(uint16(len(payload)) + 4)

	w := packet.Writer()
	w.WriteU32(seqid)
	w.WriteRawBytes(payload)
	data := w.Data()
	//log.Println("----%#v", data)

	if KEY_EXCHANGE {
		encoder.XORKeyStream(data, data)
	}
	writer.WriteRawBytes(data)
	conn.Write(writer.Data())
	//log.Printf("send : %#v", writer.Data())

	fmt.Println("发送完毕")
}
