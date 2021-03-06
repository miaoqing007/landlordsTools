// Code generated by gen_proto.sh.
// DO NOT EDIT!
package client_handler

import "app/misc/packet"

type S_null_struct struct {
}

func (p S_null_struct) Pack(w *packet.Packet) {
}

func PKT_null_struct(reader *packet.Packet) (tbl S_null_struct, err error) {
	return
}

type S_byte_id struct {
	F_id uint8
}

func (p S_byte_id) Pack(w *packet.Packet) {
	w.WriteByte(p.F_id)
}

func PKT_byte_id(reader *packet.Packet) (tbl S_byte_id, err error) {
	tbl.F_id, err = reader.ReadByte()
	checkErr(err)

	return
}

type S_auto_id struct {
	F_id int32
}

func (p S_auto_id) Pack(w *packet.Packet) {
	w.WriteS32(p.F_id)
}

func PKT_auto_id(reader *packet.Packet) (tbl S_auto_id, err error) {
	tbl.F_id, err = reader.ReadS32()
	checkErr(err)

	return
}

type S_entity_id struct {
	F_id string
}

func (p S_entity_id) Pack(w *packet.Packet) {
	w.WriteString(p.F_id)
}

func PKT_entity_id(reader *packet.Packet) (tbl S_entity_id, err error) {
	tbl.F_id, err = reader.ReadString()
	checkErr(err)

	return
}

type S_item_id struct {
	F_id uint32
}

func (p S_item_id) Pack(w *packet.Packet) {
	w.WriteU32(p.F_id)
}

func PKT_item_id(reader *packet.Packet) (tbl S_item_id, err error) {
	tbl.F_id, err = reader.ReadU32()
	checkErr(err)

	return
}

type S_player_card struct {
	F_hole_cards []string
	F_roomId     string
	F_playerIds  []string
	F_players    []S_player
}

func (p S_player_card) Pack(w *packet.Packet) {
	w.WriteU16(uint16(len(p.F_hole_cards)))
	for k := range p.F_hole_cards {
		w.WriteString(p.F_hole_cards[k])
	}
	w.WriteString(p.F_roomId)
	w.WriteU16(uint16(len(p.F_playerIds)))
	for k := range p.F_playerIds {
		w.WriteString(p.F_playerIds[k])
	}
	w.WriteU16(uint16(len(p.F_players)))
	for k := range p.F_players {
		p.F_players[k].Pack(w)
	}
}

func PKT_player_card(reader *packet.Packet) (tbl S_player_card, err error) {
	{
		narr, err := reader.ReadU16()
		checkErr(err)

		for i := 0; i < int(narr); i++ {
			v, err := reader.ReadString()
			tbl.F_hole_cards = append(tbl.F_hole_cards, v)
			checkErr(err)
		}
	}

	tbl.F_roomId, err = reader.ReadString()
	checkErr(err)

	{
		narr, err := reader.ReadU16()
		checkErr(err)

		for i := 0; i < int(narr); i++ {
			v, err := reader.ReadString()
			tbl.F_playerIds = append(tbl.F_playerIds, v)
			checkErr(err)
		}
	}

	{
		narr, err := reader.ReadU16()
		checkErr(err)

		tbl.F_players = make([]S_player, narr)
		for i := 0; i < int(narr); i++ {
			tbl.F_players[i], err = PKT_player(reader)
			checkErr(err)
		}
	}

	return
}

type S_player struct {
	F_id    string
	F_name  string
	F_cards []string
}

func (p S_player) Pack(w *packet.Packet) {
	w.WriteString(p.F_id)
	w.WriteString(p.F_name)
	w.WriteU16(uint16(len(p.F_cards)))
	for k := range p.F_cards {
		w.WriteString(p.F_cards[k])
	}
}

func PKT_player(reader *packet.Packet) (tbl S_player, err error) {
	tbl.F_id, err = reader.ReadString()
	checkErr(err)

	tbl.F_name, err = reader.ReadString()
	checkErr(err)

	{
		narr, err := reader.ReadU16()
		checkErr(err)

		for i := 0; i < int(narr); i++ {
			v, err := reader.ReadString()
			tbl.F_cards = append(tbl.F_cards, v)
			checkErr(err)
		}
	}

	return
}

type S_player_outof_card struct {
	F_roomId string
	F_cards  []string
}

func (p S_player_outof_card) Pack(w *packet.Packet) {
	w.WriteString(p.F_roomId)
	w.WriteU16(uint16(len(p.F_cards)))
	for k := range p.F_cards {
		w.WriteString(p.F_cards[k])
	}
}

func PKT_player_outof_card(reader *packet.Packet) (tbl S_player_outof_card, err error) {
	tbl.F_roomId, err = reader.ReadString()
	checkErr(err)

	{
		narr, err := reader.ReadU16()
		checkErr(err)

		for i := 0; i < int(narr); i++ {
			v, err := reader.ReadString()
			tbl.F_cards = append(tbl.F_cards, v)
			checkErr(err)
		}
	}

	return
}

type S_login_info struct {
	F_account  string
	F_password string
}

func (p S_login_info) Pack(w *packet.Packet) {
	w.WriteString(p.F_account)
	w.WriteString(p.F_password)
}

func PKT_login_info(reader *packet.Packet) (tbl S_login_info, err error) {
	tbl.F_account, err = reader.ReadString()
	checkErr(err)

	tbl.F_password, err = reader.ReadString()
	checkErr(err)

	return
}

type S_error_ack struct {
	F_msg string
}

func (p S_error_ack) Pack(w *packet.Packet) {
	w.WriteString(p.F_msg)
}

func PKT_error_ack(reader *packet.Packet) (tbl S_error_ack, err error) {
	tbl.F_msg, err = reader.ReadString()
	checkErr(err)

	return
}

type S_user_info struct {
	F_name string
	F_uid  string
}

func (p S_user_info) Pack(w *packet.Packet) {
	w.WriteString(p.F_name)
	w.WriteString(p.F_uid)
}

func PKT_user_info(reader *packet.Packet) (tbl S_user_info, err error) {
	tbl.F_name, err = reader.ReadString()
	checkErr(err)

	tbl.F_uid, err = reader.ReadString()
	checkErr(err)

	return
}

type S_out_of_cards struct {
	F_id         string
	F_cards      []string
	F_outOfCards []string
	F_randomNum  string
	F_ty         int32
}

func (p S_out_of_cards) Pack(w *packet.Packet) {
	w.WriteString(p.F_id)
	w.WriteU16(uint16(len(p.F_cards)))
	for k := range p.F_cards {
		w.WriteString(p.F_cards[k])
	}
	w.WriteU16(uint16(len(p.F_outOfCards)))
	for k := range p.F_outOfCards {
		w.WriteString(p.F_outOfCards[k])
	}
	w.WriteString(p.F_randomNum)
	w.WriteS32(p.F_ty)
}

func PKT_out_of_cards(reader *packet.Packet) (tbl S_out_of_cards, err error) {
	tbl.F_id, err = reader.ReadString()
	checkErr(err)

	{
		narr, err := reader.ReadU16()
		checkErr(err)

		for i := 0; i < int(narr); i++ {
			v, err := reader.ReadString()
			tbl.F_cards = append(tbl.F_cards, v)
			checkErr(err)
		}
	}

	{
		narr, err := reader.ReadU16()
		checkErr(err)

		for i := 0; i < int(narr); i++ {
			v, err := reader.ReadString()
			tbl.F_outOfCards = append(tbl.F_outOfCards, v)
			checkErr(err)
		}
	}

	tbl.F_randomNum, err = reader.ReadString()
	checkErr(err)

	tbl.F_ty, err = reader.ReadS32()
	checkErr(err)

	return
}

type S_msg_string struct {
	F_msg string
}

func (p S_msg_string) Pack(w *packet.Packet) {
	w.WriteString(p.F_msg)
}

func PKT_msg_string(reader *packet.Packet) (tbl S_msg_string, err error) {
	tbl.F_msg, err = reader.ReadString()
	checkErr(err)

	return
}

type S_game_over struct {
	F_winId []string
}

func (p S_game_over) Pack(w *packet.Packet) {
	w.WriteU16(uint16(len(p.F_winId)))
	for k := range p.F_winId {
		w.WriteString(p.F_winId[k])
	}
}

func PKT_game_over(reader *packet.Packet) (tbl S_game_over, err error) {
	{
		narr, err := reader.ReadU16()
		checkErr(err)

		for i := 0; i < int(narr); i++ {
			v, err := reader.ReadString()
			tbl.F_winId = append(tbl.F_winId, v)
			checkErr(err)
		}
	}

	return
}

type S_grab_landowner struct {
	F_roomId          string
	F_ifGrab          bool
	F_uid             string
	F_ifhavelandowner bool
	F_ifcall          bool
}

func (p S_grab_landowner) Pack(w *packet.Packet) {
	w.WriteString(p.F_roomId)
	w.WriteBool(p.F_ifGrab)
	w.WriteString(p.F_uid)
	w.WriteBool(p.F_ifhavelandowner)
	w.WriteBool(p.F_ifcall)
}

func PKT_grab_landowner(reader *packet.Packet) (tbl S_grab_landowner, err error) {
	tbl.F_roomId, err = reader.ReadString()
	checkErr(err)

	tbl.F_ifGrab, err = reader.ReadBool()
	checkErr(err)

	tbl.F_uid, err = reader.ReadString()
	checkErr(err)

	tbl.F_ifhavelandowner, err = reader.ReadBool()
	checkErr(err)

	tbl.F_ifcall, err = reader.ReadBool()
	checkErr(err)

	return
}

type S_chat_msg struct {
	F_name    string
	F_timeStr string
	F_msg     string
}

func (p S_chat_msg) Pack(w *packet.Packet) {
	w.WriteString(p.F_name)
	w.WriteString(p.F_timeStr)
	w.WriteString(p.F_msg)
}

func PKT_chat_msg(reader *packet.Packet) (tbl S_chat_msg, err error) {
	tbl.F_name, err = reader.ReadString()
	checkErr(err)

	tbl.F_timeStr, err = reader.ReadString()
	checkErr(err)

	tbl.F_msg, err = reader.ReadString()
	checkErr(err)

	return
}

func checkErr(err error) {
	if err != nil {
		panic("error occured in protocol module")
	}
}
