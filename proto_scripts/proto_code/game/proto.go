// Code generated by gen_proto.sh.
// DO NOT EDIT!
package client_proto

import "landlords/misc/packet"
import "fmt"
import "encoding/json"

type S_null_struct struct {
}

func (p S_null_struct) Pack(w *packet.Packet) {
}

func PKT_null_struct(data []byte) (tbl S_null_struct, err error) {
	err = json.Unmarshal(data, &tbl)
	if err != nil {
		return tbl, err
	}
	return
}

type S_byte_id struct {
	F_id uint8 `json:"id"`
}

func (p S_byte_id) Pack(w *packet.Packet) {
	w.WriteByte(p.F_id)
}

func PKT_byte_id(data []byte) (tbl S_byte_id, err error) {
	err = json.Unmarshal(data, &tbl)
	if err != nil {
		return tbl, err
	}
	return
}

type S_auto_id struct {
	F_id int32 `json:"id"`
}

func (p S_auto_id) Pack(w *packet.Packet) {
	w.WriteS32(p.F_id)
}

func PKT_auto_id(data []byte) (tbl S_auto_id, err error) {
	err = json.Unmarshal(data, &tbl)
	if err != nil {
		return tbl, err
	}
	return
}

type S_entity_id struct {
	F_id string `json:"id"`
}

func (p S_entity_id) Pack(w *packet.Packet) {
	w.WriteString(p.F_id)
}

func PKT_entity_id(data []byte) (tbl S_entity_id, err error) {
	err = json.Unmarshal(data, &tbl)
	if err != nil {
		return tbl, err
	}
	return
}

type S_item_id struct {
	F_id uint32 `json:"id"`
}

func (p S_item_id) Pack(w *packet.Packet) {
	w.WriteU32(p.F_id)
}

func PKT_item_id(data []byte) (tbl S_item_id, err error) {
	err = json.Unmarshal(data, &tbl)
	if err != nil {
		return tbl, err
	}
	return
}

type S_player_card struct {
	F_hole_cards []string   `json:"hole_cards"`
	F_roomId     string     `json:"roomId"`
	F_players    []S_player `json:"players"`
}

func (p S_player_card) Pack(w *packet.Packet) {
	w.WriteU16(uint16(len(p.F_hole_cards)))
	for k := range p.F_hole_cards {
		w.WriteString(p.F_hole_cards[k])
	}
	w.WriteString(p.F_roomId)
	w.WriteU16(uint16(len(p.F_players)))
	for k := range p.F_players {
		p.F_players[k].Pack(w)
	}
}

func PKT_player_card(data []byte) (tbl S_player_card, err error) {
	err = json.Unmarshal(data, &tbl)
	if err != nil {
		return tbl, err
	}
	return
}

type S_player struct {
	F_id    string   `json:"id"`
	F_cards []string `json:"cards"`
}

func (p S_player) Pack(w *packet.Packet) {
	w.WriteString(p.F_id)
	w.WriteU16(uint16(len(p.F_cards)))
	for k := range p.F_cards {
		w.WriteString(p.F_cards[k])
	}
}

func PKT_player(data []byte) (tbl S_player, err error) {
	err = json.Unmarshal(data, &tbl)
	if err != nil {
		return tbl, err
	}
	return
}

type S_player_outof_card struct {
	F_roomId string   `json:"roomId"`
	F_cards  []string `json:"cards"`
}

func (p S_player_outof_card) Pack(w *packet.Packet) {
	w.WriteString(p.F_roomId)
	w.WriteU16(uint16(len(p.F_cards)))
	for k := range p.F_cards {
		w.WriteString(p.F_cards[k])
	}
}

func PKT_player_outof_card(data []byte) (tbl S_player_outof_card, err error) {
	err = json.Unmarshal(data, &tbl)
	if err != nil {
		return tbl, err
	}
	return
}

type S_login_info struct {
	F_account  string `json:"account"`
	F_password string `json:"password"`
}

func (p S_login_info) Pack(w *packet.Packet) {
	w.WriteString(p.F_account)
	w.WriteString(p.F_password)
}

func PKT_login_info(data []byte) (tbl S_login_info, err error) {
	err = json.Unmarshal(data, &tbl)
	if err != nil {
		return tbl, err
	}
	return
}

type S_error_ack struct {
	F_msg string `json:"msg"`
}

func (p S_error_ack) Pack(w *packet.Packet) {
	w.WriteString(p.F_msg)
}

func PKT_error_ack(data []byte) (tbl S_error_ack, err error) {
	err = json.Unmarshal(data, &tbl)
	if err != nil {
		return tbl, err
	}
	return
}

type S_user_info struct {
	F_name string `json:"name"`
	F_uid  string `json:"uid"`
}

func (p S_user_info) Pack(w *packet.Packet) {
	w.WriteString(p.F_name)
	w.WriteString(p.F_uid)
}

func PKT_user_info(data []byte) (tbl S_user_info, err error) {
	err = json.Unmarshal(data, &tbl)
	if err != nil {
		return tbl, err
	}
	return
}

func checkErr(err error) {
	if err != nil {
		panic("error occured in protocol module")
	}
}

func checkMax(path string, v float64, max float64) {
	if v > max {
		s := fmt.Sprintf("error range in %s, v=%g, max=%g", path, v, max)
		panic(s)
	}
}

func checkMin(path string, v float64, min float64) {
	if v < min {
		s := fmt.Sprintf("error range in %s, v=%g, min=%g", path, v, min)
		panic(s)
	}
}
