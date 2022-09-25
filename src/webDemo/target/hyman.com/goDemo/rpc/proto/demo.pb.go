// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.6
// source: webDemo/rpc/proto/demo.proto

package proto

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

//enum 用于定义常量枚举，对应于 go 中的 const
type SystemView int32

const (
	SystemView_SYSTEM_VIEW_DEFAULT          SystemView = 0
	SystemView_SYSTEM_VIEW_WITH_PERMISSIONS SystemView = 1
)

// Enum value maps for SystemView.
var (
	SystemView_name = map[int32]string{
		0: "SYSTEM_VIEW_DEFAULT",
		1: "SYSTEM_VIEW_WITH_PERMISSIONS",
	}
	SystemView_value = map[string]int32{
		"SYSTEM_VIEW_DEFAULT":          0,
		"SYSTEM_VIEW_WITH_PERMISSIONS": 1,
	}
)

func (x SystemView) Enum() *SystemView {
	p := new(SystemView)
	*p = x
	return p
}

func (x SystemView) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SystemView) Descriptor() protoreflect.EnumDescriptor {
	return file_webDemo_rpc_proto_demo_proto_enumTypes[0].Descriptor()
}

func (SystemView) Type() protoreflect.EnumType {
	return &file_webDemo_rpc_proto_demo_proto_enumTypes[0]
}

func (x SystemView) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SystemView.Descriptor instead.
func (SystemView) EnumDescriptor() ([]byte, []int) {
	return file_webDemo_rpc_proto_demo_proto_rawDescGZIP(), []int{0}
}

//message 用于定义一个消息类型，对应于 go 中的结构体
type SystemPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemId     int32      `protobuf:"varint,1,opt,name=system_id,json=systemId,proto3" json:"system_id,omitempty"`
	PermissionId float64    `protobuf:"fixed64,2,opt,name=permission_id,json=permissionId,proto3" json:"permission_id,omitempty"`
	ActionId     string     `protobuf:"bytes,3,opt,name=action_id,json=actionId,proto3" json:"action_id,omitempty"`
	View         SystemView `protobuf:"varint,4,opt,name=view,proto3,enum=webDemo.rpc.proto.SystemView" json:"view,omitempty"`
	//数组，会在 go 中生成 int 类型的切片
	Ids []int32 `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *SystemPermission) Reset() {
	*x = SystemPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webDemo_rpc_proto_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemPermission) ProtoMessage() {}

func (x *SystemPermission) ProtoReflect() protoreflect.Message {
	mi := &file_webDemo_rpc_proto_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemPermission.ProtoReflect.Descriptor instead.
func (*SystemPermission) Descriptor() ([]byte, []int) {
	return file_webDemo_rpc_proto_demo_proto_rawDescGZIP(), []int{0}
}

func (x *SystemPermission) GetSystemId() int32 {
	if x != nil {
		return x.SystemId
	}
	return 0
}

func (x *SystemPermission) GetPermissionId() float64 {
	if x != nil {
		return x.PermissionId
	}
	return 0
}

func (x *SystemPermission) GetActionId() string {
	if x != nil {
		return x.ActionId
	}
	return ""
}

func (x *SystemPermission) GetView() SystemView {
	if x != nil {
		return x.View
	}
	return SystemView_SYSTEM_VIEW_DEFAULT
}

func (x *SystemPermission) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_webDemo_rpc_proto_demo_proto protoreflect.FileDescriptor

var file_webDemo_rpc_proto_demo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x65, 0x62, 0x44, 0x65, 0x6d, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x77, 0x65, 0x62, 0x44, 0x65, 0x6d, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x77, 0x65, 0x62, 0x44, 0x65, 0x6d, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb6, 0x01, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x77, 0x65, 0x62, 0x44, 0x65, 0x6d, 0x6f, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x69, 0x65,
	0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x2a, 0x47, 0x0a, 0x0a, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x59, 0x53, 0x54, 0x45,
	0x4d, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00,
	0x12, 0x20, 0x0a, 0x1c, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f,
	0x57, 0x49, 0x54, 0x48, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53,
	0x10, 0x01, 0x42, 0x3a, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x79, 0x6d, 0x61, 0x6e, 0x2e,
	0x67, 0x6f, 0x44, 0x65, 0x6d, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x1a, 0x68, 0x79, 0x6d, 0x61, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x44, 0x65, 0x6d, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_webDemo_rpc_proto_demo_proto_rawDescOnce sync.Once
	file_webDemo_rpc_proto_demo_proto_rawDescData = file_webDemo_rpc_proto_demo_proto_rawDesc
)

func file_webDemo_rpc_proto_demo_proto_rawDescGZIP() []byte {
	file_webDemo_rpc_proto_demo_proto_rawDescOnce.Do(func() {
		file_webDemo_rpc_proto_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_webDemo_rpc_proto_demo_proto_rawDescData)
	})
	return file_webDemo_rpc_proto_demo_proto_rawDescData
}

var file_webDemo_rpc_proto_demo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_webDemo_rpc_proto_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_webDemo_rpc_proto_demo_proto_goTypes = []interface{}{
	(SystemView)(0),          // 0: webDemo.rpc.proto.SystemView
	(*SystemPermission)(nil), // 1: webDemo.rpc.proto.SystemPermission
}
var file_webDemo_rpc_proto_demo_proto_depIdxs = []int32{
	0, // 0: webDemo.rpc.proto.SystemPermission.view:type_name -> webDemo.rpc.proto.SystemView
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_webDemo_rpc_proto_demo_proto_init() }
func file_webDemo_rpc_proto_demo_proto_init() {
	if File_webDemo_rpc_proto_demo_proto != nil {
		return
	}
	file_webDemo_rpc_proto_demo2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_webDemo_rpc_proto_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemPermission); i {
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
			RawDescriptor: file_webDemo_rpc_proto_demo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_webDemo_rpc_proto_demo_proto_goTypes,
		DependencyIndexes: file_webDemo_rpc_proto_demo_proto_depIdxs,
		EnumInfos:         file_webDemo_rpc_proto_demo_proto_enumTypes,
		MessageInfos:      file_webDemo_rpc_proto_demo_proto_msgTypes,
	}.Build()
	File_webDemo_rpc_proto_demo_proto = out.File
	file_webDemo_rpc_proto_demo_proto_rawDesc = nil
	file_webDemo_rpc_proto_demo_proto_goTypes = nil
	file_webDemo_rpc_proto_demo_proto_depIdxs = nil
}