// Code generated by protoc-gen-go.
// source: m.proto
// DO NOT EDIT!

package consensus

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type msg_Cmd int32

const (
	msg_NOP      msg_Cmd = 0
	msg_INVITE   msg_Cmd = 1
	msg_RSVP     msg_Cmd = 2
	msg_NOMINATE msg_Cmd = 3
	msg_VOTE     msg_Cmd = 4
	msg_TICK     msg_Cmd = 5
	msg_PROPOSE  msg_Cmd = 6
	msg_LEARN    msg_Cmd = 7
)

var msg_Cmd_name = map[int32]string{
	0: "NOP",
	1: "INVITE",
	2: "RSVP",
	3: "NOMINATE",
	4: "VOTE",
	5: "TICK",
	6: "PROPOSE",
	7: "LEARN",
}
var msg_Cmd_value = map[string]int32{
	"NOP":      0,
	"INVITE":   1,
	"RSVP":     2,
	"NOMINATE": 3,
	"VOTE":     4,
	"TICK":     5,
	"PROPOSE":  6,
	"LEARN":    7,
}

func (x msg_Cmd) Enum() *msg_Cmd {
	p := new(msg_Cmd)
	*p = x
	return p
}
func (x msg_Cmd) String() string {
	return proto.EnumName(msg_Cmd_name, int32(x))
}
func (x msg_Cmd) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *msg_Cmd) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(msg_Cmd_value, data, "msg_Cmd")
	if err != nil {
		return err
	}
	*x = msg_Cmd(value)
	return nil
}

type msg struct {
	Cmd              *msg_Cmd `protobuf:"varint,1,opt,name=cmd,enum=consensus.msg_Cmd" json:"cmd,omitempty"`
	Seqn             *int64   `protobuf:"varint,2,opt,name=seqn" json:"seqn,omitempty"`
	Crnd             *int64   `protobuf:"varint,3,opt,name=crnd" json:"crnd,omitempty"`
	Vrnd             *int64   `protobuf:"varint,4,opt,name=vrnd" json:"vrnd,omitempty"`
	Value            []byte   `protobuf:"bytes,5,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *msg) Reset()         { *this = msg{} }
func (this *msg) String() string { return proto.CompactTextString(this) }
func (*msg) ProtoMessage()       {}

func (this *msg) GetCmd() msg_Cmd {
	if this != nil && this.Cmd != nil {
		return *this.Cmd
	}
	return 0
}

func (this *msg) GetSeqn() int64 {
	if this != nil && this.Seqn != nil {
		return *this.Seqn
	}
	return 0
}

func (this *msg) GetCrnd() int64 {
	if this != nil && this.Crnd != nil {
		return *this.Crnd
	}
	return 0
}

func (this *msg) GetVrnd() int64 {
	if this != nil && this.Vrnd != nil {
		return *this.Vrnd
	}
	return 0
}

func (this *msg) GetValue() []byte {
	if this != nil {
		return this.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("consensus.msg_Cmd", msg_Cmd_name, msg_Cmd_value)
}
