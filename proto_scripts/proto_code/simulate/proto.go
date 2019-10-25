// Code generated by gen_proto.sh.
// DO NOT EDIT!
package main

import "misc/packet"

type null_struct struct {
}

func (p null_struct) Pack(w *packet.Packet) {
}

func PKT_null_struct(reader *packet.Packet) (tbl null_struct, err error) {
	return
}

type byte_id struct {
	F_id uint8
}

func (p byte_id) Pack(w *packet.Packet) {
	w.WriteByte(p.F_id)
}

func PKT_byte_id(reader *packet.Packet) (tbl byte_id, err error) {
	tbl.F_id, err = reader.ReadByte()
	checkErr(err)

	return
}

type auto_id struct {
	F_id int32
}

func (p auto_id) Pack(w *packet.Packet) {
	w.WriteS32(p.F_id)
}

func PKT_auto_id(reader *packet.Packet) (tbl auto_id, err error) {
	tbl.F_id, err = reader.ReadS32()
	checkErr(err)

	return
}

type entity_id struct {
	F_id uint64
}

func (p entity_id) Pack(w *packet.Packet) {
	w.WriteU64(p.F_id)
}

func PKT_entity_id(reader *packet.Packet) (tbl entity_id, err error) {
	tbl.F_id, err = reader.ReadU64()
	checkErr(err)

	return
}

type item_id struct {
	F_id uint32
}

func (p item_id) Pack(w *packet.Packet) {
	w.WriteU32(p.F_id)
}

func PKT_item_id(reader *packet.Packet) (tbl item_id, err error) {
	tbl.F_id, err = reader.ReadU32()
	checkErr(err)

	return
}