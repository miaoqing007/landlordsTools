// Code generated by gen_proto.sh.
// DO NOT EDIT!
using UnityEngine;
using System;
using System.Collections.Generic;

namespace NetProto.Proto
{
    public abstract class NetBase
    {
        public short NetMsgId;
        public abstract void Pack(ByteArray w);
    }

    public class null_struct : NetBase
    { 

        public override void Pack(ByteArray w)
        { 
        }

        public static null_struct UnPack(ByteArray reader)
        {
            null_struct tbl = new null_struct();

            return tbl;
        }
    }

    public class byte_id : NetBase
    { 
        public byte id;

        public override void Pack(ByteArray w)
        { 
            w.WriteUnsignedInt8(this.id);
        }

        public static byte_id UnPack(ByteArray reader)
        {
            byte_id tbl = new byte_id();
            tbl.id = reader.ReadUnsignedInt8();

            return tbl;
        }
    }

    public class auto_id : NetBase
    { 
        public Int32 id;

        public override void Pack(ByteArray w)
        { 
            w.WriteInt32(this.id);
        }

        public static auto_id UnPack(ByteArray reader)
        {
            auto_id tbl = new auto_id();
            tbl.id = reader.ReadInt32();

            return tbl;
        }
    }

    public class entity_id : NetBase
    { 
        public string id;

        public override void Pack(ByteArray w)
        { 
            w.WriteUTF(this.id);
        }

        public static entity_id UnPack(ByteArray reader)
        {
            entity_id tbl = new entity_id();
            tbl.id = reader.ReadUTFBytes();

            return tbl;
        }
    }

    public class item_id : NetBase
    { 
        public UInt32 id;

        public override void Pack(ByteArray w)
        { 
            w.WriteUnsignedInt(this.id);
        }

        public static item_id UnPack(ByteArray reader)
        {
            item_id tbl = new item_id();
            tbl.id = reader.ReadUnsignedInt32();

            return tbl;
        }
    }

    public class player_card : NetBase
    { 
        public string[] hole_cards;
        public string roomId;
        public player[] players;

        public override void Pack(ByteArray w)
        { 
            w.WriteUnsignedInt16((UInt16)this.hole_cards.Length);
            for (int k = 0; k < this.hole_cards.Length; k++)
            { 
                w.WriteUTF(this.hole_cards[k]);
            }
            w.WriteUTF(this.roomId);
            w.WriteUnsignedInt16((UInt16)this.players.Length);
            for (int k = 0; k < this.players.Length; k++)
            { 
                this.players[k].Pack(w);
            }
        }

        public static player_card UnPack(ByteArray reader)
        {
            player_card tbl = new player_card();
            {
                UInt16 narr = reader.ReadUnsignedInt16();
                
                tbl.hole_cards = new string[narr];
                
                for (int i = 0; i < narr; i++)
                {
                    tbl.hole_cards[i] = reader.ReadUTFBytes();
                }
            }
            tbl.roomId = reader.ReadUTFBytes();
            {
                UInt16 narr = reader.ReadUnsignedInt16();
                tbl.players = new player[narr];
                for (int i = 0; i < narr; i++)
                {
                    tbl.players[i] = player.UnPack(reader);
                }
            }

            return tbl;
        }
    }

    public class player : NetBase
    { 
        public string id;
        public string[] cards;

        public override void Pack(ByteArray w)
        { 
            w.WriteUTF(this.id);
            w.WriteUnsignedInt16((UInt16)this.cards.Length);
            for (int k = 0; k < this.cards.Length; k++)
            { 
                w.WriteUTF(this.cards[k]);
            }
        }

        public static player UnPack(ByteArray reader)
        {
            player tbl = new player();
            tbl.id = reader.ReadUTFBytes();
            {
                UInt16 narr = reader.ReadUnsignedInt16();
                
                tbl.cards = new string[narr];
                
                for (int i = 0; i < narr; i++)
                {
                    tbl.cards[i] = reader.ReadUTFBytes();
                }
            }

            return tbl;
        }
    }

    public class player_outof_card : NetBase
    { 
        public string roomId;
        public string[] cards;

        public override void Pack(ByteArray w)
        { 
            w.WriteUTF(this.roomId);
            w.WriteUnsignedInt16((UInt16)this.cards.Length);
            for (int k = 0; k < this.cards.Length; k++)
            { 
                w.WriteUTF(this.cards[k]);
            }
        }

        public static player_outof_card UnPack(ByteArray reader)
        {
            player_outof_card tbl = new player_outof_card();
            tbl.roomId = reader.ReadUTFBytes();
            {
                UInt16 narr = reader.ReadUnsignedInt16();
                
                tbl.cards = new string[narr];
                
                for (int i = 0; i < narr; i++)
                {
                    tbl.cards[i] = reader.ReadUTFBytes();
                }
            }

            return tbl;
        }
    }

    public class login_info : NetBase
    { 
        public string account;
        public string password;

        public override void Pack(ByteArray w)
        { 
            w.WriteUTF(this.account);
            w.WriteUTF(this.password);
        }

        public static login_info UnPack(ByteArray reader)
        {
            login_info tbl = new login_info();
            tbl.account = reader.ReadUTFBytes();
            tbl.password = reader.ReadUTFBytes();

            return tbl;
        }
    }

    public class error_ack : NetBase
    { 
        public string msg;

        public override void Pack(ByteArray w)
        { 
            w.WriteUTF(this.msg);
        }

        public static error_ack UnPack(ByteArray reader)
        {
            error_ack tbl = new error_ack();
            tbl.msg = reader.ReadUTFBytes();

            return tbl;
        }
    }

    public class user_info : NetBase
    { 
        public string name;
        public string uid;

        public override void Pack(ByteArray w)
        { 
            w.WriteUTF(this.name);
            w.WriteUTF(this.uid);
        }

        public static user_info UnPack(ByteArray reader)
        {
            user_info tbl = new user_info();
            tbl.name = reader.ReadUTFBytes();
            tbl.uid = reader.ReadUTFBytes();

            return tbl;
        }
    }
}
