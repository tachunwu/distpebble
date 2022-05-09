// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: core/v1/core.proto

package corev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TxnType int32

const (
	TxnType_UNKNOWN                 TxnType = 0
	TxnType_SINGLE_HOME             TxnType = 1
	TxnType_MULTI_HOME_OR_LOCK_ONLY TxnType = 2
)

// Enum value maps for TxnType.
var (
	TxnType_name = map[int32]string{
		0: "UNKNOWN",
		1: "SINGLE_HOME",
		2: "MULTI_HOME_OR_LOCK_ONLY",
	}
	TxnType_value = map[string]int32{
		"UNKNOWN":                 0,
		"SINGLE_HOME":             1,
		"MULTI_HOME_OR_LOCK_ONLY": 2,
	}
)

func (x TxnType) Enum() *TxnType {
	p := new(TxnType)
	*p = x
	return p
}

func (x TxnType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TxnType) Descriptor() protoreflect.EnumDescriptor {
	return file_core_v1_core_proto_enumTypes[0].Descriptor()
}

func (TxnType) Type() protoreflect.EnumType {
	return &file_core_v1_core_proto_enumTypes[0]
}

func (x TxnType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TxnType.Descriptor instead.
func (TxnType) EnumDescriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{0}
}

type TxnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txn *Txn `protobuf:"bytes,1,opt,name=txn,proto3" json:"txn,omitempty"`
}

func (x *TxnRequest) Reset() {
	*x = TxnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnRequest) ProtoMessage() {}

func (x *TxnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnRequest.ProtoReflect.Descriptor instead.
func (*TxnRequest) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{0}
}

func (x *TxnRequest) GetTxn() *Txn {
	if x != nil {
		return x.Txn
	}
	return nil
}

type TxnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TxnResponse) Reset() {
	*x = TxnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnResponse) ProtoMessage() {}

func (x *TxnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnResponse.ProtoReflect.Descriptor instead.
func (*TxnResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{1}
}

type Txn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Global order
	TxnId   uint64  `protobuf:"varint,1,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
	TxnType TxnType `protobuf:"varint,2,opt,name=txn_type,json=txnType,proto3,enum=core.v1.TxnType" json:"txn_type,omitempty"`
	// Key sets
	ReadSet      []*KeyEntry `protobuf:"bytes,20,rep,name=read_set,json=readSet,proto3" json:"read_set,omitempty"`
	WriteSet     []*KeyEntry `protobuf:"bytes,21,rep,name=write_set,json=writeSet,proto3" json:"write_set,omitempty"`
	ReadWriteSet []*KeyEntry `protobuf:"bytes,22,rep,name=read_write_set,json=readWriteSet,proto3" json:"read_write_set,omitempty"`
	// Reader and writer
	Readers []uint64 `protobuf:"varint,31,rep,packed,name=readers,proto3" json:"readers,omitempty"`
	Writers []uint64 `protobuf:"varint,32,rep,packed,name=writers,proto3" json:"writers,omitempty"`
}

func (x *Txn) Reset() {
	*x = Txn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Txn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Txn) ProtoMessage() {}

func (x *Txn) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Txn.ProtoReflect.Descriptor instead.
func (*Txn) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{2}
}

func (x *Txn) GetTxnId() uint64 {
	if x != nil {
		return x.TxnId
	}
	return 0
}

func (x *Txn) GetTxnType() TxnType {
	if x != nil {
		return x.TxnType
	}
	return TxnType_UNKNOWN
}

func (x *Txn) GetReadSet() []*KeyEntry {
	if x != nil {
		return x.ReadSet
	}
	return nil
}

func (x *Txn) GetWriteSet() []*KeyEntry {
	if x != nil {
		return x.WriteSet
	}
	return nil
}

func (x *Txn) GetReadWriteSet() []*KeyEntry {
	if x != nil {
		return x.ReadWriteSet
	}
	return nil
}

func (x *Txn) GetReaders() []uint64 {
	if x != nil {
		return x.Readers
	}
	return nil
}

func (x *Txn) GetWriters() []uint64 {
	if x != nil {
		return x.Writers
	}
	return nil
}

// Basic statement
type Get struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *KeyEntry `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Get) Reset() {
	*x = Get{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Get) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Get) ProtoMessage() {}

func (x *Get) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Get.ProtoReflect.Descriptor instead.
func (*Get) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{3}
}

func (x *Get) GetKey() *KeyEntry {
	if x != nil {
		return x.Key
	}
	return nil
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *KeyEntry `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Valie []byte    `protobuf:"bytes,2,opt,name=valie,proto3" json:"valie,omitempty"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{4}
}

func (x *Set) GetKey() *KeyEntry {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Set) GetValie() []byte {
	if x != nil {
		return x.Valie
	}
	return nil
}

type Delete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *KeyEntry `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Delete) Reset() {
	*x = Delete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delete) ProtoMessage() {}

func (x *Delete) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delete.ProtoReflect.Descriptor instead.
func (*Delete) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{5}
}

func (x *Delete) GetKey() *KeyEntry {
	if x != nil {
		return x.Key
	}
	return nil
}

type Scan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartKey *KeyEntry `protobuf:"bytes,1,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
	EndKey   *KeyEntry `protobuf:"bytes,2,opt,name=end_key,json=endKey,proto3" json:"end_key,omitempty"`
}

func (x *Scan) Reset() {
	*x = Scan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scan) ProtoMessage() {}

func (x *Scan) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scan.ProtoReflect.Descriptor instead.
func (*Scan) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{6}
}

func (x *Scan) GetStartKey() *KeyEntry {
	if x != nil {
		return x.StartKey
	}
	return nil
}

func (x *Scan) GetEndKey() *KeyEntry {
	if x != nil {
		return x.EndKey
	}
	return nil
}

type KeyEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Master  uint32 `protobuf:"varint,2,opt,name=master,proto3" json:"master,omitempty"`
	Counter uint64 `protobuf:"varint,3,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *KeyEntry) Reset() {
	*x = KeyEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyEntry) ProtoMessage() {}

func (x *KeyEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyEntry.ProtoReflect.Descriptor instead.
func (*KeyEntry) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{7}
}

func (x *KeyEntry) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KeyEntry) GetMaster() uint32 {
	if x != nil {
		return x.Master
	}
	return 0
}

func (x *KeyEntry) GetCounter() uint64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

type LookupMasterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxnId uint64   `protobuf:"varint,1,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
	Keys  [][]byte `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *LookupMasterRequest) Reset() {
	*x = LookupMasterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupMasterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupMasterRequest) ProtoMessage() {}

func (x *LookupMasterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupMasterRequest.ProtoReflect.Descriptor instead.
func (*LookupMasterRequest) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{8}
}

func (x *LookupMasterRequest) GetTxnId() uint64 {
	if x != nil {
		return x.TxnId
	}
	return 0
}

func (x *LookupMasterRequest) GetKeys() [][]byte {
	if x != nil {
		return x.Keys
	}
	return nil
}

type LookupMasterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxnId      uint64      `protobuf:"varint,1,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
	KeyEntries []*KeyEntry `protobuf:"bytes,2,rep,name=key_entries,json=keyEntries,proto3" json:"key_entries,omitempty"`
}

func (x *LookupMasterResponse) Reset() {
	*x = LookupMasterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupMasterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupMasterResponse) ProtoMessage() {}

func (x *LookupMasterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupMasterResponse.ProtoReflect.Descriptor instead.
func (*LookupMasterResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{9}
}

func (x *LookupMasterResponse) GetTxnId() uint64 {
	if x != nil {
		return x.TxnId
	}
	return 0
}

func (x *LookupMasterResponse) GetKeyEntries() []*KeyEntry {
	if x != nil {
		return x.KeyEntries
	}
	return nil
}

// Cluster manage
type MachineInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Replica uint64 `protobuf:"varint,2,opt,name=replica,proto3" json:"replica,omitempty"`
	Host    string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port    uint64 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *MachineInfo) Reset() {
	*x = MachineInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineInfo) ProtoMessage() {}

func (x *MachineInfo) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineInfo.ProtoReflect.Descriptor instead.
func (*MachineInfo) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{10}
}

func (x *MachineInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MachineInfo) GetReplica() uint64 {
	if x != nil {
		return x.Replica
	}
	return 0
}

func (x *MachineInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *MachineInfo) GetPort() uint64 {
	if x != nil {
		return x.Port
	}
	return 0
}

type ClusterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machines []*MachineInfo `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"`
}

func (x *ClusterInfo) Reset() {
	*x = ClusterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_core_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterInfo) ProtoMessage() {}

func (x *ClusterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_core_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterInfo.ProtoReflect.Descriptor instead.
func (*ClusterInfo) Descriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{11}
}

func (x *ClusterInfo) GetMachines() []*MachineInfo {
	if x != nil {
		return x.Machines
	}
	return nil
}

var File_core_v1_core_proto protoreflect.FileDescriptor

var file_core_v1_core_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x2c, 0x0a,
	0x0a, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x03, 0x74,
	0x78, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x78, 0x6e, 0x52, 0x03, 0x74, 0x78, 0x6e, 0x22, 0x0d, 0x0a, 0x0b, 0x54,
	0x78, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x94, 0x02, 0x0a, 0x03, 0x54,
	0x78, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x78, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x74, 0x78, 0x6e, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x74, 0x78, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x74,
	0x78, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73,
	0x65, 0x74, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x61,
	0x64, 0x53, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x65,
	0x74, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x53, 0x65, 0x74, 0x12, 0x37, 0x0a, 0x0e, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0c, 0x72, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x20, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x73, 0x22, 0x2a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x40, 0x0a,
	0x03, 0x53, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x65, 0x22,
	0x2d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x62,
	0x0a, 0x04, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x4b,
	0x65, 0x79, 0x22, 0x4e, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x22, 0x40, 0x0a, 0x13, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x78, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x78, 0x6e, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x22, 0x61, 0x0a, 0x14, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x74, 0x78, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x78,
	0x6e, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6b, 0x65, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x5f, 0x0a, 0x0b, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x3f, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x2a, 0x44, 0x0a, 0x07, 0x54, 0x78, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x5f, 0x48, 0x4f, 0x4d, 0x45,
	0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f, 0x48, 0x4f, 0x4d, 0x45,
	0x5f, 0x4f, 0x52, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x32,
	0x46, 0x0a, 0x10, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x54, 0x78, 0x6e, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x5e, 0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x4c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x8e, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x43, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x61, 0x63, 0x68, 0x75, 0x6e, 0x77, 0x75, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x70, 0x65,
	0x62, 0x62, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07,
	0x43, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08,
	0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_v1_core_proto_rawDescOnce sync.Once
	file_core_v1_core_proto_rawDescData = file_core_v1_core_proto_rawDesc
)

func file_core_v1_core_proto_rawDescGZIP() []byte {
	file_core_v1_core_proto_rawDescOnce.Do(func() {
		file_core_v1_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_v1_core_proto_rawDescData)
	})
	return file_core_v1_core_proto_rawDescData
}

var file_core_v1_core_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_v1_core_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_core_v1_core_proto_goTypes = []interface{}{
	(TxnType)(0),                 // 0: core.v1.TxnType
	(*TxnRequest)(nil),           // 1: core.v1.TxnRequest
	(*TxnResponse)(nil),          // 2: core.v1.TxnResponse
	(*Txn)(nil),                  // 3: core.v1.Txn
	(*Get)(nil),                  // 4: core.v1.Get
	(*Set)(nil),                  // 5: core.v1.Set
	(*Delete)(nil),               // 6: core.v1.Delete
	(*Scan)(nil),                 // 7: core.v1.Scan
	(*KeyEntry)(nil),             // 8: core.v1.KeyEntry
	(*LookupMasterRequest)(nil),  // 9: core.v1.LookupMasterRequest
	(*LookupMasterResponse)(nil), // 10: core.v1.LookupMasterResponse
	(*MachineInfo)(nil),          // 11: core.v1.MachineInfo
	(*ClusterInfo)(nil),          // 12: core.v1.ClusterInfo
}
var file_core_v1_core_proto_depIdxs = []int32{
	3,  // 0: core.v1.TxnRequest.txn:type_name -> core.v1.Txn
	0,  // 1: core.v1.Txn.txn_type:type_name -> core.v1.TxnType
	8,  // 2: core.v1.Txn.read_set:type_name -> core.v1.KeyEntry
	8,  // 3: core.v1.Txn.write_set:type_name -> core.v1.KeyEntry
	8,  // 4: core.v1.Txn.read_write_set:type_name -> core.v1.KeyEntry
	8,  // 5: core.v1.Get.key:type_name -> core.v1.KeyEntry
	8,  // 6: core.v1.Set.key:type_name -> core.v1.KeyEntry
	8,  // 7: core.v1.Delete.key:type_name -> core.v1.KeyEntry
	8,  // 8: core.v1.Scan.start_key:type_name -> core.v1.KeyEntry
	8,  // 9: core.v1.Scan.end_key:type_name -> core.v1.KeyEntry
	8,  // 10: core.v1.LookupMasterResponse.key_entries:type_name -> core.v1.KeyEntry
	11, // 11: core.v1.ClusterInfo.machines:type_name -> core.v1.MachineInfo
	1,  // 12: core.v1.SequencerService.Txn:input_type -> core.v1.TxnRequest
	9,  // 13: core.v1.LookupService.LookupMaster:input_type -> core.v1.LookupMasterRequest
	2,  // 14: core.v1.SequencerService.Txn:output_type -> core.v1.TxnResponse
	10, // 15: core.v1.LookupService.LookupMaster:output_type -> core.v1.LookupMasterResponse
	14, // [14:16] is the sub-list for method output_type
	12, // [12:14] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_core_v1_core_proto_init() }
func file_core_v1_core_proto_init() {
	if File_core_v1_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_v1_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Txn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Get); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Delete); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scan); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_core_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_core_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupMasterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_core_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupMasterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_core_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_core_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_v1_core_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_core_v1_core_proto_goTypes,
		DependencyIndexes: file_core_v1_core_proto_depIdxs,
		EnumInfos:         file_core_v1_core_proto_enumTypes,
		MessageInfos:      file_core_v1_core_proto_msgTypes,
	}.Build()
	File_core_v1_core_proto = out.File
	file_core_v1_core_proto_rawDesc = nil
	file_core_v1_core_proto_goTypes = nil
	file_core_v1_core_proto_depIdxs = nil
}