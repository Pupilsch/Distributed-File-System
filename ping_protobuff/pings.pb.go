// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: pings.proto

package pings

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type UDPAddressProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   []byte `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Zone string `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"`
}

func (x *UDPAddressProto) Reset() {
	*x = UDPAddressProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UDPAddressProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UDPAddressProto) ProtoMessage() {}

func (x *UDPAddressProto) ProtoReflect() protoreflect.Message {
	mi := &file_pings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UDPAddressProto.ProtoReflect.Descriptor instead.
func (*UDPAddressProto) Descriptor() ([]byte, []int) {
	return file_pings_proto_rawDescGZIP(), []int{0}
}

func (x *UDPAddressProto) GetIp() []byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *UDPAddressProto) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *UDPAddressProto) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

type FileMetaDataProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileSize      uint64               `protobuf:"varint,1,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
	LastTime      *timestamp.Timestamp `protobuf:"bytes,2,opt,name=lastTime,proto3" json:"lastTime,omitempty"`
	Servers       []int32              `protobuf:"varint,3,rep,packed,name=servers,proto3" json:"servers,omitempty"`
	Writing       bool                 `protobuf:"varint,4,opt,name=writing,proto3" json:"writing,omitempty"`
	NumAckWriting uint32               `protobuf:"varint,5,opt,name=numAckWriting,proto3" json:"numAckWriting,omitempty"`
	NumAckReading uint32               `protobuf:"varint,6,opt,name=numAckReading,proto3" json:"numAckReading,omitempty"`
	NumReading    uint32               `protobuf:"varint,7,opt,name=numReading,proto3" json:"numReading,omitempty"`
}

func (x *FileMetaDataProto) Reset() {
	*x = FileMetaDataProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetaDataProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetaDataProto) ProtoMessage() {}

func (x *FileMetaDataProto) ProtoReflect() protoreflect.Message {
	mi := &file_pings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetaDataProto.ProtoReflect.Descriptor instead.
func (*FileMetaDataProto) Descriptor() ([]byte, []int) {
	return file_pings_proto_rawDescGZIP(), []int{1}
}

func (x *FileMetaDataProto) GetFileSize() uint64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *FileMetaDataProto) GetLastTime() *timestamp.Timestamp {
	if x != nil {
		return x.LastTime
	}
	return nil
}

func (x *FileMetaDataProto) GetServers() []int32 {
	if x != nil {
		return x.Servers
	}
	return nil
}

func (x *FileMetaDataProto) GetWriting() bool {
	if x != nil {
		return x.Writing
	}
	return false
}

func (x *FileMetaDataProto) GetNumAckWriting() uint32 {
	if x != nil {
		return x.NumAckWriting
	}
	return 0
}

func (x *FileMetaDataProto) GetNumAckReading() uint32 {
	if x != nil {
		return x.NumAckReading
	}
	return 0
}

func (x *FileMetaDataProto) GetNumReading() uint32 {
	if x != nil {
		return x.NumReading
	}
	return 0
}

type TableEntryProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address        *UDPAddressProto     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	JoinTime       *timestamp.Timestamp `protobuf:"bytes,2,opt,name=joinTime,proto3" json:"joinTime,omitempty"`
	LastTime       *timestamp.Timestamp `protobuf:"bytes,3,opt,name=lastTime,proto3" json:"lastTime,omitempty"`
	Index          uint32               `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	Status         int32                `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	FileMemoryUsed int64                `protobuf:"varint,6,opt,name=fileMemoryUsed,proto3" json:"fileMemoryUsed,omitempty"`
}

func (x *TableEntryProto) Reset() {
	*x = TableEntryProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableEntryProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableEntryProto) ProtoMessage() {}

func (x *TableEntryProto) ProtoReflect() protoreflect.Message {
	mi := &file_pings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableEntryProto.ProtoReflect.Descriptor instead.
func (*TableEntryProto) Descriptor() ([]byte, []int) {
	return file_pings_proto_rawDescGZIP(), []int{2}
}

func (x *TableEntryProto) GetAddress() *UDPAddressProto {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *TableEntryProto) GetJoinTime() *timestamp.Timestamp {
	if x != nil {
		return x.JoinTime
	}
	return nil
}

func (x *TableEntryProto) GetLastTime() *timestamp.Timestamp {
	if x != nil {
		return x.LastTime
	}
	return nil
}

func (x *TableEntryProto) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *TableEntryProto) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TableEntryProto) GetFileMemoryUsed() int64 {
	if x != nil {
		return x.FileMemoryUsed
	}
	return 0
}

type TableMessasgeProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType  int32                         `protobuf:"varint,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Sender   *TableEntryProto              `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Table    []*TableEntryProto            `protobuf:"bytes,3,rep,name=table,proto3" json:"table,omitempty"`
	Gossip   bool                          `protobuf:"varint,4,opt,name=gossip,proto3" json:"gossip,omitempty"`
	Files    map[string]*FileMetaDataProto `protobuf:"bytes,5,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Filename string                        `protobuf:"bytes,6,opt,name=filename,proto3" json:"filename,omitempty"`
	FileData *FileMetaDataProto            `protobuf:"bytes,7,opt,name=fileData,proto3" json:"fileData,omitempty"`
}

func (x *TableMessasgeProto) Reset() {
	*x = TableMessasgeProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableMessasgeProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableMessasgeProto) ProtoMessage() {}

func (x *TableMessasgeProto) ProtoReflect() protoreflect.Message {
	mi := &file_pings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableMessasgeProto.ProtoReflect.Descriptor instead.
func (*TableMessasgeProto) Descriptor() ([]byte, []int) {
	return file_pings_proto_rawDescGZIP(), []int{3}
}

func (x *TableMessasgeProto) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *TableMessasgeProto) GetSender() *TableEntryProto {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *TableMessasgeProto) GetTable() []*TableEntryProto {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *TableMessasgeProto) GetGossip() bool {
	if x != nil {
		return x.Gossip
	}
	return false
}

func (x *TableMessasgeProto) GetFiles() map[string]*FileMetaDataProto {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *TableMessasgeProto) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *TableMessasgeProto) GetFileData() *FileMetaDataProto {
	if x != nil {
		return x.FileData
	}
	return nil
}

var File_pings_proto protoreflect.FileDescriptor

var file_pings_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x0f,
	0x55, 0x44, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x22, 0x87, 0x02, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x77,
	0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x77, 0x72,
	0x69, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x41, 0x63, 0x6b, 0x57,
	0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6e, 0x75,
	0x6d, 0x41, 0x63, 0x6b, 0x57, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x6e,
	0x75, 0x6d, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x22, 0xa0, 0x02, 0x0a, 0x0f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x47, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x55, 0x44, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x36,
	0x0a, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6a, 0x6f,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e,
	0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x55, 0x73, 0x65, 0x64, 0x22, 0xf9, 0x03, 0x0a, 0x12, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x73, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x05,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x12, 0x51, 0x0a, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x73, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x69, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pings_proto_rawDescOnce sync.Once
	file_pings_proto_rawDescData = file_pings_proto_rawDesc
)

func file_pings_proto_rawDescGZIP() []byte {
	file_pings_proto_rawDescOnce.Do(func() {
		file_pings_proto_rawDescData = protoimpl.X.CompressGZIP(file_pings_proto_rawDescData)
	})
	return file_pings_proto_rawDescData
}

var file_pings_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pings_proto_goTypes = []interface{}{
	(*UDPAddressProto)(nil),     // 0: distributed_group_membership.UDPAddressProto
	(*FileMetaDataProto)(nil),   // 1: distributed_group_membership.FileMetaDataProto
	(*TableEntryProto)(nil),     // 2: distributed_group_membership.TableEntryProto
	(*TableMessasgeProto)(nil),  // 3: distributed_group_membership.TableMessasgeProto
	nil,                         // 4: distributed_group_membership.TableMessasgeProto.FilesEntry
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_pings_proto_depIdxs = []int32{
	5, // 0: distributed_group_membership.FileMetaDataProto.lastTime:type_name -> google.protobuf.Timestamp
	0, // 1: distributed_group_membership.TableEntryProto.address:type_name -> distributed_group_membership.UDPAddressProto
	5, // 2: distributed_group_membership.TableEntryProto.joinTime:type_name -> google.protobuf.Timestamp
	5, // 3: distributed_group_membership.TableEntryProto.lastTime:type_name -> google.protobuf.Timestamp
	2, // 4: distributed_group_membership.TableMessasgeProto.sender:type_name -> distributed_group_membership.TableEntryProto
	2, // 5: distributed_group_membership.TableMessasgeProto.table:type_name -> distributed_group_membership.TableEntryProto
	4, // 6: distributed_group_membership.TableMessasgeProto.files:type_name -> distributed_group_membership.TableMessasgeProto.FilesEntry
	1, // 7: distributed_group_membership.TableMessasgeProto.fileData:type_name -> distributed_group_membership.FileMetaDataProto
	1, // 8: distributed_group_membership.TableMessasgeProto.FilesEntry.value:type_name -> distributed_group_membership.FileMetaDataProto
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_pings_proto_init() }
func file_pings_proto_init() {
	if File_pings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UDPAddressProto); i {
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
		file_pings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMetaDataProto); i {
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
		file_pings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableEntryProto); i {
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
		file_pings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableMessasgeProto); i {
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
			RawDescriptor: file_pings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pings_proto_goTypes,
		DependencyIndexes: file_pings_proto_depIdxs,
		MessageInfos:      file_pings_proto_msgTypes,
	}.Build()
	File_pings_proto = out.File
	file_pings_proto_rawDesc = nil
	file_pings_proto_goTypes = nil
	file_pings_proto_depIdxs = nil
}
