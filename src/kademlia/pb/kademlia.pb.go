// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: kademlia.proto

package pb

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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_kademlia_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   []byte `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	IP   string `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Port int32  `protobuf:"varint,3,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_kademlia_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetID() []byte {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Node) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *Node) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Init   int32  `protobuf:"varint,1,opt,name=Init,proto3" json:"Init,omitempty"`
	End    int32  `protobuf:"varint,2,opt,name=End,proto3" json:"End,omitempty"`
	Buffer []byte `protobuf:"bytes,3,opt,name=Buffer,proto3" json:"Buffer,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_kademlia_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetInit() int32 {
	if x != nil {
		return x.Init
	}
	return 0
}

func (x *Data) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Data) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

type TargetID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *TargetID) Reset() {
	*x = TargetID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetID) ProtoMessage() {}

func (x *TargetID) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetID.ProtoReflect.Descriptor instead.
func (*TargetID) Descriptor() ([]byte, []int) {
	return file_kademlia_proto_rawDescGZIP(), []int{3}
}

func (x *TargetID) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type KBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket []*Node `protobuf:"bytes,1,rep,name=Bucket,proto3" json:"Bucket,omitempty"`
}

func (x *KBucket) Reset() {
	*x = KBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KBucket) ProtoMessage() {}

func (x *KBucket) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KBucket.ProtoReflect.Descriptor instead.
func (*KBucket) Descriptor() ([]byte, []int) {
	return file_kademlia_proto_rawDescGZIP(), []int{4}
}

func (x *KBucket) GetBucket() []*Node {
	if x != nil {
		return x.Bucket
	}
	return nil
}

type FindValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KNeartestBuckets *KBucket `protobuf:"bytes,1,opt,name=KNeartestBuckets,proto3,oneof" json:"KNeartestBuckets,omitempty"`
	Value            *Data    `protobuf:"bytes,2,opt,name=Value,proto3,oneof" json:"Value,omitempty"`
}

func (x *FindValueResponse) Reset() {
	*x = FindValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindValueResponse) ProtoMessage() {}

func (x *FindValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindValueResponse.ProtoReflect.Descriptor instead.
func (*FindValueResponse) Descriptor() ([]byte, []int) {
	return file_kademlia_proto_rawDescGZIP(), []int{5}
}

func (x *FindValueResponse) GetKNeartestBuckets() *KBucket {
	if x != nil {
		return x.KNeartestBuckets
	}
	return nil
}

func (x *FindValueResponse) GetValue() *Data {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_kademlia_proto protoreflect.FileDescriptor

var file_kademlia_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6b, 0x61, 0x64, 0x65, 0x6d, 0x6c, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6b, 0x61, 0x64, 0x65, 0x6d, 0x6c, 0x69, 0x61, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x3a, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x44, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x45, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x42, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x22, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x31,
	0x0a, 0x07, 0x4b, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6b, 0x61, 0x64, 0x65,
	0x6d, 0x6c, 0x69, 0x61, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x22, 0xa1, 0x01, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x4b, 0x4e, 0x65, 0x61, 0x72,
	0x74, 0x65, 0x73, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6b, 0x61, 0x64, 0x65, 0x6d, 0x6c, 0x69, 0x61, 0x2e, 0x4b, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x48, 0x00, 0x52, 0x10, 0x4b, 0x4e, 0x65, 0x61, 0x72, 0x74, 0x65, 0x73,
	0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6b, 0x61, 0x64,
	0x65, 0x6d, 0x6c, 0x69, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x01, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x4b, 0x4e, 0x65, 0x61, 0x72,
	0x74, 0x65, 0x73, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xdc, 0x01, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x2e, 0x6b, 0x61, 0x64,
	0x65, 0x6d, 0x6c, 0x69, 0x61, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x0e, 0x2e, 0x6b, 0x61, 0x64,
	0x65, 0x6d, 0x6c, 0x69, 0x61, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x05,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x0e, 0x2e, 0x6b, 0x61, 0x64, 0x65, 0x6d, 0x6c, 0x69, 0x61,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x12, 0x2e, 0x6b, 0x61, 0x64, 0x65, 0x6d, 0x6c, 0x69, 0x61,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x33, 0x0a,
	0x08, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x2e, 0x6b, 0x61, 0x64, 0x65,
	0x6d, 0x6c, 0x69, 0x61, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x1a, 0x11, 0x2e,
	0x6b, 0x61, 0x64, 0x65, 0x6d, 0x6c, 0x69, 0x61, 0x2e, 0x4b, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x12, 0x2e, 0x6b, 0x61, 0x64, 0x65, 0x6d, 0x6c, 0x69, 0x61, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x44, 0x1a, 0x1b, 0x2e, 0x6b, 0x61, 0x64, 0x65, 0x6d, 0x6c, 0x69, 0x61, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x72, 0x74, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6b, 0x61, 0x64, 0x65, 0x6d, 0x6c, 0x69, 0x61, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kademlia_proto_rawDescOnce sync.Once
	file_kademlia_proto_rawDescData = file_kademlia_proto_rawDesc
)

func file_kademlia_proto_rawDescGZIP() []byte {
	file_kademlia_proto_rawDescOnce.Do(func() {
		file_kademlia_proto_rawDescData = protoimpl.X.CompressGZIP(file_kademlia_proto_rawDescData)
	})
	return file_kademlia_proto_rawDescData
}

var file_kademlia_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kademlia_proto_goTypes = []interface{}{
	(*Response)(nil),          // 0: kademlia.Response
	(*Node)(nil),              // 1: kademlia.Node
	(*Data)(nil),              // 2: kademlia.Data
	(*TargetID)(nil),          // 3: kademlia.TargetID
	(*KBucket)(nil),           // 4: kademlia.KBucket
	(*FindValueResponse)(nil), // 5: kademlia.FindValueResponse
}
var file_kademlia_proto_depIdxs = []int32{
	1, // 0: kademlia.KBucket.Bucket:type_name -> kademlia.Node
	4, // 1: kademlia.FindValueResponse.KNeartestBuckets:type_name -> kademlia.KBucket
	2, // 2: kademlia.FindValueResponse.Value:type_name -> kademlia.Data
	1, // 3: kademlia.FullNode.Ping:input_type -> kademlia.Node
	2, // 4: kademlia.FullNode.Store:input_type -> kademlia.Data
	3, // 5: kademlia.FullNode.FindNode:input_type -> kademlia.TargetID
	3, // 6: kademlia.FullNode.FindValue:input_type -> kademlia.TargetID
	1, // 7: kademlia.FullNode.Ping:output_type -> kademlia.Node
	0, // 8: kademlia.FullNode.Store:output_type -> kademlia.Response
	4, // 9: kademlia.FullNode.FindNode:output_type -> kademlia.KBucket
	5, // 10: kademlia.FullNode.FindValue:output_type -> kademlia.FindValueResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kademlia_proto_init() }
func file_kademlia_proto_init() {
	if File_kademlia_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kademlia_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_kademlia_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_kademlia_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_kademlia_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetID); i {
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
		file_kademlia_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KBucket); i {
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
		file_kademlia_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindValueResponse); i {
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
	file_kademlia_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kademlia_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kademlia_proto_goTypes,
		DependencyIndexes: file_kademlia_proto_depIdxs,
		MessageInfos:      file_kademlia_proto_msgTypes,
	}.Build()
	File_kademlia_proto = out.File
	file_kademlia_proto_rawDesc = nil
	file_kademlia_proto_goTypes = nil
	file_kademlia_proto_depIdxs = nil
}
