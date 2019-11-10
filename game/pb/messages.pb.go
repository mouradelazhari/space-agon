// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game/pb/messages.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SpawnEvent_SpawnType int32

const (
	SpawnEvent_UNKNOWN   SpawnEvent_SpawnType = 0
	SpawnEvent_SHIP      SpawnEvent_SpawnType = 1
	SpawnEvent_MISSILE   SpawnEvent_SpawnType = 2
	SpawnEvent_EXPLOSION SpawnEvent_SpawnType = 3
)

var SpawnEvent_SpawnType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SHIP",
	2: "MISSILE",
	3: "EXPLOSION",
}

var SpawnEvent_SpawnType_value = map[string]int32{
	"UNKNOWN":   0,
	"SHIP":      1,
	"MISSILE":   2,
	"EXPLOSION": 3,
}

func (x SpawnEvent_SpawnType) String() string {
	return proto.EnumName(SpawnEvent_SpawnType_name, int32(x))
}

func (SpawnEvent_SpawnType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{8, 0}
}

type ClientInitialize struct {
	Cid                  int64    `protobuf:"varint,1,opt,name=cid,proto3" json:"cid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientInitialize) Reset()         { *m = ClientInitialize{} }
func (m *ClientInitialize) String() string { return proto.CompactTextString(m) }
func (*ClientInitialize) ProtoMessage()    {}
func (*ClientInitialize) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{0}
}

func (m *ClientInitialize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInitialize.Unmarshal(m, b)
}
func (m *ClientInitialize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInitialize.Marshal(b, m, deterministic)
}
func (m *ClientInitialize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInitialize.Merge(m, src)
}
func (m *ClientInitialize) XXX_Size() int {
	return xxx_messageInfo_ClientInitialize.Size(m)
}
func (m *ClientInitialize) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInitialize.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInitialize proto.InternalMessageInfo

func (m *ClientInitialize) GetCid() int64 {
	if m != nil {
		return m.Cid
	}
	return 0
}

type Memos struct {
	Memos                []*Memo  `protobuf:"bytes,1,rep,name=memos,proto3" json:"memos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Memos) Reset()         { *m = Memos{} }
func (m *Memos) String() string { return proto.CompactTextString(m) }
func (*Memos) ProtoMessage()    {}
func (*Memos) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{1}
}

func (m *Memos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Memos.Unmarshal(m, b)
}
func (m *Memos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Memos.Marshal(b, m, deterministic)
}
func (m *Memos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memos.Merge(m, src)
}
func (m *Memos) XXX_Size() int {
	return xxx_messageInfo_Memos.Size(m)
}
func (m *Memos) XXX_DiscardUnknown() {
	xxx_messageInfo_Memos.DiscardUnknown(m)
}

var xxx_messageInfo_Memos proto.InternalMessageInfo

func (m *Memos) GetMemos() []*Memo {
	if m != nil {
		return m.Memos
	}
	return nil
}

type Memo struct {
	// Types that are valid to be assigned to Recipient:
	//	*Memo_To
	//	*Memo_EveryoneBut
	//	*Memo_Everyone
	Recipient isMemo_Recipient `protobuf_oneof:"recipient"`
	// Types that are valid to be assigned to Actual:
	//	*Memo_PosTracks
	//	*Memo_MomentumTracks
	//	*Memo_RotTracks
	//	*Memo_SpinTracks
	//	*Memo_ShipControlTrack
	//	*Memo_SpawnEvent
	//	*Memo_DestroyEvent
	Actual               isMemo_Actual `protobuf_oneof:"actual"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Memo) Reset()         { *m = Memo{} }
func (m *Memo) String() string { return proto.CompactTextString(m) }
func (*Memo) ProtoMessage()    {}
func (*Memo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{2}
}

func (m *Memo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Memo.Unmarshal(m, b)
}
func (m *Memo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Memo.Marshal(b, m, deterministic)
}
func (m *Memo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memo.Merge(m, src)
}
func (m *Memo) XXX_Size() int {
	return xxx_messageInfo_Memo.Size(m)
}
func (m *Memo) XXX_DiscardUnknown() {
	xxx_messageInfo_Memo.DiscardUnknown(m)
}

var xxx_messageInfo_Memo proto.InternalMessageInfo

type isMemo_Recipient interface {
	isMemo_Recipient()
}

type Memo_To struct {
	To int64 `protobuf:"varint,1,opt,name=to,proto3,oneof"`
}

type Memo_EveryoneBut struct {
	EveryoneBut int64 `protobuf:"varint,2,opt,name=everyone_but,json=everyoneBut,proto3,oneof"`
}

type Memo_Everyone struct {
	Everyone bool `protobuf:"varint,3,opt,name=everyone,proto3,oneof"`
}

func (*Memo_To) isMemo_Recipient() {}

func (*Memo_EveryoneBut) isMemo_Recipient() {}

func (*Memo_Everyone) isMemo_Recipient() {}

func (m *Memo) GetRecipient() isMemo_Recipient {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *Memo) GetTo() int64 {
	if x, ok := m.GetRecipient().(*Memo_To); ok {
		return x.To
	}
	return 0
}

func (m *Memo) GetEveryoneBut() int64 {
	if x, ok := m.GetRecipient().(*Memo_EveryoneBut); ok {
		return x.EveryoneBut
	}
	return 0
}

func (m *Memo) GetEveryone() bool {
	if x, ok := m.GetRecipient().(*Memo_Everyone); ok {
		return x.Everyone
	}
	return false
}

type isMemo_Actual interface {
	isMemo_Actual()
}

type Memo_PosTracks struct {
	PosTracks *PosTracks `protobuf:"bytes,10,opt,name=pos_tracks,json=posTracks,proto3,oneof"`
}

type Memo_MomentumTracks struct {
	MomentumTracks *MomentumTracks `protobuf:"bytes,11,opt,name=momentum_tracks,json=momentumTracks,proto3,oneof"`
}

type Memo_RotTracks struct {
	RotTracks *RotTracks `protobuf:"bytes,12,opt,name=rot_tracks,json=rotTracks,proto3,oneof"`
}

type Memo_SpinTracks struct {
	SpinTracks *SpinTracks `protobuf:"bytes,13,opt,name=spin_tracks,json=spinTracks,proto3,oneof"`
}

type Memo_ShipControlTrack struct {
	ShipControlTrack *ShipControlTrack `protobuf:"bytes,14,opt,name=ship_control_track,json=shipControlTrack,proto3,oneof"`
}

type Memo_SpawnEvent struct {
	SpawnEvent *SpawnEvent `protobuf:"bytes,15,opt,name=spawn_event,json=spawnEvent,proto3,oneof"`
}

type Memo_DestroyEvent struct {
	DestroyEvent *DestroyEvent `protobuf:"bytes,16,opt,name=destroy_event,json=destroyEvent,proto3,oneof"`
}

func (*Memo_PosTracks) isMemo_Actual() {}

func (*Memo_MomentumTracks) isMemo_Actual() {}

func (*Memo_RotTracks) isMemo_Actual() {}

func (*Memo_SpinTracks) isMemo_Actual() {}

func (*Memo_ShipControlTrack) isMemo_Actual() {}

func (*Memo_SpawnEvent) isMemo_Actual() {}

func (*Memo_DestroyEvent) isMemo_Actual() {}

func (m *Memo) GetActual() isMemo_Actual {
	if m != nil {
		return m.Actual
	}
	return nil
}

func (m *Memo) GetPosTracks() *PosTracks {
	if x, ok := m.GetActual().(*Memo_PosTracks); ok {
		return x.PosTracks
	}
	return nil
}

func (m *Memo) GetMomentumTracks() *MomentumTracks {
	if x, ok := m.GetActual().(*Memo_MomentumTracks); ok {
		return x.MomentumTracks
	}
	return nil
}

func (m *Memo) GetRotTracks() *RotTracks {
	if x, ok := m.GetActual().(*Memo_RotTracks); ok {
		return x.RotTracks
	}
	return nil
}

func (m *Memo) GetSpinTracks() *SpinTracks {
	if x, ok := m.GetActual().(*Memo_SpinTracks); ok {
		return x.SpinTracks
	}
	return nil
}

func (m *Memo) GetShipControlTrack() *ShipControlTrack {
	if x, ok := m.GetActual().(*Memo_ShipControlTrack); ok {
		return x.ShipControlTrack
	}
	return nil
}

func (m *Memo) GetSpawnEvent() *SpawnEvent {
	if x, ok := m.GetActual().(*Memo_SpawnEvent); ok {
		return x.SpawnEvent
	}
	return nil
}

func (m *Memo) GetDestroyEvent() *DestroyEvent {
	if x, ok := m.GetActual().(*Memo_DestroyEvent); ok {
		return x.DestroyEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Memo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Memo_To)(nil),
		(*Memo_EveryoneBut)(nil),
		(*Memo_Everyone)(nil),
		(*Memo_PosTracks)(nil),
		(*Memo_MomentumTracks)(nil),
		(*Memo_RotTracks)(nil),
		(*Memo_SpinTracks)(nil),
		(*Memo_ShipControlTrack)(nil),
		(*Memo_SpawnEvent)(nil),
		(*Memo_DestroyEvent)(nil),
	}
}

type PosTracks struct {
	Nid                  []uint64  `protobuf:"varint,1,rep,packed,name=nid,proto3" json:"nid,omitempty"`
	X                    []float32 `protobuf:"fixed32,2,rep,packed,name=x,proto3" json:"x,omitempty"`
	Y                    []float32 `protobuf:"fixed32,3,rep,packed,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PosTracks) Reset()         { *m = PosTracks{} }
func (m *PosTracks) String() string { return proto.CompactTextString(m) }
func (*PosTracks) ProtoMessage()    {}
func (*PosTracks) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{3}
}

func (m *PosTracks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PosTracks.Unmarshal(m, b)
}
func (m *PosTracks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PosTracks.Marshal(b, m, deterministic)
}
func (m *PosTracks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PosTracks.Merge(m, src)
}
func (m *PosTracks) XXX_Size() int {
	return xxx_messageInfo_PosTracks.Size(m)
}
func (m *PosTracks) XXX_DiscardUnknown() {
	xxx_messageInfo_PosTracks.DiscardUnknown(m)
}

var xxx_messageInfo_PosTracks proto.InternalMessageInfo

func (m *PosTracks) GetNid() []uint64 {
	if m != nil {
		return m.Nid
	}
	return nil
}

func (m *PosTracks) GetX() []float32 {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *PosTracks) GetY() []float32 {
	if m != nil {
		return m.Y
	}
	return nil
}

type MomentumTracks struct {
	Nid                  []uint64  `protobuf:"varint,1,rep,packed,name=nid,proto3" json:"nid,omitempty"`
	X                    []float32 `protobuf:"fixed32,2,rep,packed,name=x,proto3" json:"x,omitempty"`
	Y                    []float32 `protobuf:"fixed32,3,rep,packed,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MomentumTracks) Reset()         { *m = MomentumTracks{} }
func (m *MomentumTracks) String() string { return proto.CompactTextString(m) }
func (*MomentumTracks) ProtoMessage()    {}
func (*MomentumTracks) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{4}
}

func (m *MomentumTracks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MomentumTracks.Unmarshal(m, b)
}
func (m *MomentumTracks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MomentumTracks.Marshal(b, m, deterministic)
}
func (m *MomentumTracks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MomentumTracks.Merge(m, src)
}
func (m *MomentumTracks) XXX_Size() int {
	return xxx_messageInfo_MomentumTracks.Size(m)
}
func (m *MomentumTracks) XXX_DiscardUnknown() {
	xxx_messageInfo_MomentumTracks.DiscardUnknown(m)
}

var xxx_messageInfo_MomentumTracks proto.InternalMessageInfo

func (m *MomentumTracks) GetNid() []uint64 {
	if m != nil {
		return m.Nid
	}
	return nil
}

func (m *MomentumTracks) GetX() []float32 {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *MomentumTracks) GetY() []float32 {
	if m != nil {
		return m.Y
	}
	return nil
}

type RotTracks struct {
	Nid                  []uint64  `protobuf:"varint,1,rep,packed,name=nid,proto3" json:"nid,omitempty"`
	R                    []float32 `protobuf:"fixed32,2,rep,packed,name=r,proto3" json:"r,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RotTracks) Reset()         { *m = RotTracks{} }
func (m *RotTracks) String() string { return proto.CompactTextString(m) }
func (*RotTracks) ProtoMessage()    {}
func (*RotTracks) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{5}
}

func (m *RotTracks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RotTracks.Unmarshal(m, b)
}
func (m *RotTracks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RotTracks.Marshal(b, m, deterministic)
}
func (m *RotTracks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RotTracks.Merge(m, src)
}
func (m *RotTracks) XXX_Size() int {
	return xxx_messageInfo_RotTracks.Size(m)
}
func (m *RotTracks) XXX_DiscardUnknown() {
	xxx_messageInfo_RotTracks.DiscardUnknown(m)
}

var xxx_messageInfo_RotTracks proto.InternalMessageInfo

func (m *RotTracks) GetNid() []uint64 {
	if m != nil {
		return m.Nid
	}
	return nil
}

func (m *RotTracks) GetR() []float32 {
	if m != nil {
		return m.R
	}
	return nil
}

type SpinTracks struct {
	Nid                  []uint64  `protobuf:"varint,1,rep,packed,name=nid,proto3" json:"nid,omitempty"`
	S                    []float32 `protobuf:"fixed32,2,rep,packed,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SpinTracks) Reset()         { *m = SpinTracks{} }
func (m *SpinTracks) String() string { return proto.CompactTextString(m) }
func (*SpinTracks) ProtoMessage()    {}
func (*SpinTracks) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{6}
}

func (m *SpinTracks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpinTracks.Unmarshal(m, b)
}
func (m *SpinTracks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpinTracks.Marshal(b, m, deterministic)
}
func (m *SpinTracks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpinTracks.Merge(m, src)
}
func (m *SpinTracks) XXX_Size() int {
	return xxx_messageInfo_SpinTracks.Size(m)
}
func (m *SpinTracks) XXX_DiscardUnknown() {
	xxx_messageInfo_SpinTracks.DiscardUnknown(m)
}

var xxx_messageInfo_SpinTracks proto.InternalMessageInfo

func (m *SpinTracks) GetNid() []uint64 {
	if m != nil {
		return m.Nid
	}
	return nil
}

func (m *SpinTracks) GetS() []float32 {
	if m != nil {
		return m.S
	}
	return nil
}

type ShipControlTrack struct {
	Nid                  uint64   `protobuf:"varint,1,opt,name=nid,proto3" json:"nid,omitempty"`
	Up                   bool     `protobuf:"varint,2,opt,name=up,proto3" json:"up,omitempty"`
	Left                 bool     `protobuf:"varint,3,opt,name=left,proto3" json:"left,omitempty"`
	Right                bool     `protobuf:"varint,4,opt,name=right,proto3" json:"right,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipControlTrack) Reset()         { *m = ShipControlTrack{} }
func (m *ShipControlTrack) String() string { return proto.CompactTextString(m) }
func (*ShipControlTrack) ProtoMessage()    {}
func (*ShipControlTrack) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{7}
}

func (m *ShipControlTrack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipControlTrack.Unmarshal(m, b)
}
func (m *ShipControlTrack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipControlTrack.Marshal(b, m, deterministic)
}
func (m *ShipControlTrack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipControlTrack.Merge(m, src)
}
func (m *ShipControlTrack) XXX_Size() int {
	return xxx_messageInfo_ShipControlTrack.Size(m)
}
func (m *ShipControlTrack) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipControlTrack.DiscardUnknown(m)
}

var xxx_messageInfo_ShipControlTrack proto.InternalMessageInfo

func (m *ShipControlTrack) GetNid() uint64 {
	if m != nil {
		return m.Nid
	}
	return 0
}

func (m *ShipControlTrack) GetUp() bool {
	if m != nil {
		return m.Up
	}
	return false
}

func (m *ShipControlTrack) GetLeft() bool {
	if m != nil {
		return m.Left
	}
	return false
}

func (m *ShipControlTrack) GetRight() bool {
	if m != nil {
		return m.Right
	}
	return false
}

type SpawnEvent struct {
	Nid                  uint64               `protobuf:"varint,1,opt,name=nid,proto3" json:"nid,omitempty"`
	SpawnType            SpawnEvent_SpawnType `protobuf:"varint,2,opt,name=spawn_type,json=spawnType,proto3,enum=spaceagon.SpawnEvent_SpawnType" json:"spawn_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SpawnEvent) Reset()         { *m = SpawnEvent{} }
func (m *SpawnEvent) String() string { return proto.CompactTextString(m) }
func (*SpawnEvent) ProtoMessage()    {}
func (*SpawnEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{8}
}

func (m *SpawnEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnEvent.Unmarshal(m, b)
}
func (m *SpawnEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnEvent.Marshal(b, m, deterministic)
}
func (m *SpawnEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnEvent.Merge(m, src)
}
func (m *SpawnEvent) XXX_Size() int {
	return xxx_messageInfo_SpawnEvent.Size(m)
}
func (m *SpawnEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnEvent proto.InternalMessageInfo

func (m *SpawnEvent) GetNid() uint64 {
	if m != nil {
		return m.Nid
	}
	return 0
}

func (m *SpawnEvent) GetSpawnType() SpawnEvent_SpawnType {
	if m != nil {
		return m.SpawnType
	}
	return SpawnEvent_UNKNOWN
}

type DestroyEvent struct {
	Nid                  uint64   `protobuf:"varint,1,opt,name=nid,proto3" json:"nid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyEvent) Reset()         { *m = DestroyEvent{} }
func (m *DestroyEvent) String() string { return proto.CompactTextString(m) }
func (*DestroyEvent) ProtoMessage()    {}
func (*DestroyEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{9}
}

func (m *DestroyEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyEvent.Unmarshal(m, b)
}
func (m *DestroyEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyEvent.Marshal(b, m, deterministic)
}
func (m *DestroyEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyEvent.Merge(m, src)
}
func (m *DestroyEvent) XXX_Size() int {
	return xxx_messageInfo_DestroyEvent.Size(m)
}
func (m *DestroyEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyEvent proto.InternalMessageInfo

func (m *DestroyEvent) GetNid() uint64 {
	if m != nil {
		return m.Nid
	}
	return 0
}

func init() {
	proto.RegisterEnum("spaceagon.SpawnEvent_SpawnType", SpawnEvent_SpawnType_name, SpawnEvent_SpawnType_value)
	proto.RegisterType((*ClientInitialize)(nil), "spaceagon.ClientInitialize")
	proto.RegisterType((*Memos)(nil), "spaceagon.Memos")
	proto.RegisterType((*Memo)(nil), "spaceagon.Memo")
	proto.RegisterType((*PosTracks)(nil), "spaceagon.PosTracks")
	proto.RegisterType((*MomentumTracks)(nil), "spaceagon.MomentumTracks")
	proto.RegisterType((*RotTracks)(nil), "spaceagon.RotTracks")
	proto.RegisterType((*SpinTracks)(nil), "spaceagon.SpinTracks")
	proto.RegisterType((*ShipControlTrack)(nil), "spaceagon.ShipControlTrack")
	proto.RegisterType((*SpawnEvent)(nil), "spaceagon.SpawnEvent")
	proto.RegisterType((*DestroyEvent)(nil), "spaceagon.DestroyEvent")
}

func init() { proto.RegisterFile("game/pb/messages.proto", fileDescriptor_ae8bea4e98c5fae7) }

var fileDescriptor_ae8bea4e98c5fae7 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x97, 0xa4, 0x1b, 0xcd, 0xa5, 0xeb, 0x22, 0x6b, 0x40, 0x10, 0x48, 0x54, 0x01, 0xa4,
	0x4a, 0x8c, 0x56, 0x1a, 0x42, 0x82, 0x97, 0x3e, 0x74, 0x9b, 0xd4, 0x6a, 0x5b, 0x37, 0xa5, 0x43,
	0x20, 0x1e, 0xa8, 0xd2, 0xcc, 0x4b, 0x23, 0x9a, 0xd8, 0x8a, 0x9d, 0x6d, 0xe1, 0xc3, 0xf0, 0xca,
	0xd7, 0x44, 0xb6, 0xeb, 0x2c, 0xdb, 0xfa, 0xc2, 0xdb, 0xfd, 0xef, 0xee, 0x77, 0x77, 0xd1, 0x5d,
	0x0c, 0xcf, 0xe2, 0x30, 0xc5, 0x7d, 0x3a, 0xef, 0xa7, 0x98, 0xb1, 0x30, 0xc6, 0xac, 0x47, 0x73,
	0xc2, 0x09, 0xb2, 0x19, 0x0d, 0x23, 0x1c, 0xc6, 0x24, 0xf3, 0xdf, 0x82, 0x7b, 0xb0, 0x4c, 0x70,
	0xc6, 0xc7, 0x59, 0xc2, 0x93, 0x70, 0x99, 0xfc, 0xc6, 0xc8, 0x05, 0x2b, 0x4a, 0x2e, 0x3d, 0xa3,
	0x63, 0x74, 0xad, 0x40, 0x98, 0x7e, 0x0f, 0x36, 0x4f, 0x71, 0x4a, 0x18, 0x7a, 0x07, 0x9b, 0xa9,
	0x30, 0x3c, 0xa3, 0x63, 0x75, 0x9d, 0xfd, 0x9d, 0x5e, 0x55, 0xa9, 0x27, 0x12, 0x02, 0x15, 0xf5,
	0xff, 0x36, 0xa0, 0x21, 0x34, 0x72, 0xc1, 0xe4, 0x44, 0x55, 0x1a, 0x6d, 0x04, 0x26, 0x27, 0xe8,
	0x0d, 0xb4, 0xf0, 0x35, 0xce, 0x4b, 0x92, 0xe1, 0xd9, 0xbc, 0xe0, 0x9e, 0xb9, 0x8a, 0x39, 0xda,
	0x3b, 0x2c, 0x38, 0x7a, 0x05, 0x4d, 0x2d, 0x3d, 0xab, 0x63, 0x74, 0x9b, 0xa3, 0x8d, 0xa0, 0xf2,
	0xa0, 0x4f, 0x00, 0x94, 0xb0, 0x19, 0xcf, 0xc3, 0xe8, 0x17, 0xf3, 0xa0, 0x63, 0x74, 0x9d, 0xfd,
	0xdd, 0xda, 0x24, 0xe7, 0x84, 0x5d, 0xc8, 0xd8, 0xc8, 0x08, 0x6c, 0xaa, 0x05, 0x3a, 0x84, 0x9d,
	0x94, 0xa4, 0x38, 0xe3, 0x45, 0xaa, 0x59, 0x47, 0xb2, 0x2f, 0xea, 0x5f, 0xb1, 0xca, 0xa8, 0x0a,
	0xb4, 0xd3, 0x7b, 0x1e, 0xd1, 0x3c, 0x27, 0x5c, 0x17, 0x68, 0x3d, 0x6a, 0x1e, 0x10, 0x7e, 0xd7,
	0x3c, 0xd7, 0x02, 0x7d, 0x06, 0x87, 0xd1, 0x24, 0xd3, 0xdc, 0xb6, 0xe4, 0x9e, 0xd6, 0xb8, 0x29,
	0x4d, 0xb2, 0x0a, 0x04, 0x56, 0x29, 0x74, 0x0c, 0x88, 0x2d, 0x12, 0x3a, 0x8b, 0x48, 0xc6, 0x73,
	0xb2, 0x54, 0x15, 0xbc, 0xb6, 0x2c, 0xf0, 0xb2, 0x5e, 0x60, 0x91, 0xd0, 0x03, 0x95, 0x23, 0xc9,
	0x91, 0x11, 0xb8, 0xec, 0x81, 0x4f, 0x8d, 0x11, 0xde, 0x64, 0x33, 0x7c, 0x8d, 0x33, 0xee, 0xed,
	0xac, 0x19, 0x23, 0xbc, 0xc9, 0x8e, 0x44, 0x50, 0x8d, 0xa1, 0x15, 0x1a, 0xc0, 0xf6, 0x25, 0x66,
	0x3c, 0x27, 0xe5, 0x8a, 0x75, 0x25, 0xfb, 0xbc, 0xc6, 0x1e, 0xaa, 0xb8, 0xa6, 0x5b, 0x97, 0x35,
	0x3d, 0x74, 0xc0, 0xce, 0x71, 0x94, 0x50, 0x71, 0x6b, 0xc3, 0x26, 0x6c, 0x85, 0x11, 0x2f, 0xc2,
	0xa5, 0xff, 0x05, 0xec, 0x6a, 0x5d, 0xe2, 0xf0, 0x32, 0x79, 0x78, 0x56, 0xb7, 0x11, 0x08, 0x13,
	0xb5, 0xc0, 0xb8, 0xf5, 0xcc, 0x8e, 0xd5, 0x35, 0x03, 0xe3, 0x56, 0xa8, 0xd2, 0xb3, 0x94, 0x2a,
	0xfd, 0x01, 0xb4, 0xef, 0x6f, 0xeb, 0x3f, 0xf9, 0xf7, 0x60, 0x57, 0xcb, 0x5a, 0x8f, 0xe6, 0x1a,
	0xcd, 0xfd, 0x3d, 0x80, 0xbb, 0x0d, 0xad, 0xcf, 0x66, 0x3a, 0x9b, 0xf9, 0x3f, 0xc1, 0x7d, 0xb8,
	0x8e, 0x3b, 0xc6, 0xd0, 0x4c, 0x1b, 0xcc, 0x82, 0xca, 0x1f, 0xa0, 0x19, 0x98, 0x05, 0x45, 0x08,
	0x1a, 0x4b, 0x7c, 0xc5, 0xd5, 0xc5, 0x07, 0xd2, 0x46, 0xbb, 0xb0, 0x99, 0x27, 0xf1, 0x82, 0x7b,
	0x0d, 0xe9, 0x54, 0xc2, 0xff, 0x63, 0x88, 0x71, 0xaa, 0xdd, 0x3c, 0x2e, 0x3d, 0x00, 0xb5, 0xbb,
	0x19, 0x2f, 0x29, 0x96, 0x2d, 0xda, 0xfb, 0xaf, 0xd7, 0xae, 0x59, 0x99, 0x17, 0x25, 0xc5, 0x81,
	0xcd, 0xb4, 0xe9, 0x0f, 0xc0, 0xae, 0xfc, 0xc8, 0x81, 0x27, 0x5f, 0x27, 0xc7, 0x93, 0xb3, 0x6f,
	0x13, 0x77, 0x03, 0x35, 0xa1, 0x31, 0x1d, 0x8d, 0xcf, 0x5d, 0x43, 0xb8, 0x4f, 0xc7, 0xd3, 0xe9,
	0xf8, 0xe4, 0xc8, 0x35, 0xd1, 0x36, 0xd8, 0x47, 0xdf, 0xcf, 0x4f, 0xce, 0xa6, 0xe3, 0xb3, 0x89,
	0x6b, 0xf9, 0x1d, 0x68, 0xd5, 0xaf, 0xe1, 0xf1, 0x84, 0xc3, 0xde, 0x8f, 0xbd, 0x38, 0xe1, 0x8b,
	0x62, 0xde, 0x8b, 0x48, 0xda, 0x8f, 0x09, 0x89, 0x97, 0xf8, 0x8a, 0xe4, 0xe2, 0xc5, 0x62, 0x7d,
	0x39, 0xe8, 0x07, 0x31, 0x69, 0x7f, 0xf5, 0x84, 0xcd, 0xb7, 0xe4, 0xd3, 0xf5, 0xf1, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xaa, 0xb3, 0xb4, 0xac, 0xd4, 0x04, 0x00, 0x00,
}