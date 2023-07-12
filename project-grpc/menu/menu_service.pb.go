// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: menu_service.proto

package menu_service_v1

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

type MenuReqMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MenuReqMessage) Reset() {
	*x = MenuReqMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuReqMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuReqMessage) ProtoMessage() {}

func (x *MenuReqMessage) ProtoReflect() protoreflect.Message {
	mi := &file_menu_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuReqMessage.ProtoReflect.Descriptor instead.
func (*MenuReqMessage) Descriptor() ([]byte, []int) {
	return file_menu_service_proto_rawDescGZIP(), []int{0}
}

type MenuMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Pid        int64          `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Title      string         `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Icon       string         `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	Url        string         `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	FilePath   string         `protobuf:"bytes,6,opt,name=filePath,proto3" json:"filePath,omitempty"`
	Params     string         `protobuf:"bytes,7,opt,name=params,proto3" json:"params,omitempty"`
	Node       string         `protobuf:"bytes,8,opt,name=node,proto3" json:"node,omitempty"`
	Sort       int32          `protobuf:"varint,9,opt,name=sort,proto3" json:"sort,omitempty"`
	Status     int32          `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	CreateBy   int64          `protobuf:"varint,11,opt,name=createBy,proto3" json:"createBy,omitempty"`
	IsInner    int32          `protobuf:"varint,12,opt,name=isInner,proto3" json:"isInner,omitempty"`
	Values     string         `protobuf:"bytes,13,opt,name=values,proto3" json:"values,omitempty"`
	ShowSlider int32          `protobuf:"varint,14,opt,name=showSlider,proto3" json:"showSlider,omitempty"`
	StatusText string         `protobuf:"bytes,15,opt,name=statusText,proto3" json:"statusText,omitempty"`
	InnerText  string         `protobuf:"bytes,16,opt,name=innerText,proto3" json:"innerText,omitempty"`
	FullUrl    string         `protobuf:"bytes,17,opt,name=fullUrl,proto3" json:"fullUrl,omitempty"`
	Children   []*MenuMessage `protobuf:"bytes,18,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *MenuMessage) Reset() {
	*x = MenuMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuMessage) ProtoMessage() {}

func (x *MenuMessage) ProtoReflect() protoreflect.Message {
	mi := &file_menu_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuMessage.ProtoReflect.Descriptor instead.
func (*MenuMessage) Descriptor() ([]byte, []int) {
	return file_menu_service_proto_rawDescGZIP(), []int{1}
}

func (x *MenuMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MenuMessage) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *MenuMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MenuMessage) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *MenuMessage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MenuMessage) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *MenuMessage) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *MenuMessage) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *MenuMessage) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *MenuMessage) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MenuMessage) GetCreateBy() int64 {
	if x != nil {
		return x.CreateBy
	}
	return 0
}

func (x *MenuMessage) GetIsInner() int32 {
	if x != nil {
		return x.IsInner
	}
	return 0
}

func (x *MenuMessage) GetValues() string {
	if x != nil {
		return x.Values
	}
	return ""
}

func (x *MenuMessage) GetShowSlider() int32 {
	if x != nil {
		return x.ShowSlider
	}
	return 0
}

func (x *MenuMessage) GetStatusText() string {
	if x != nil {
		return x.StatusText
	}
	return ""
}

func (x *MenuMessage) GetInnerText() string {
	if x != nil {
		return x.InnerText
	}
	return ""
}

func (x *MenuMessage) GetFullUrl() string {
	if x != nil {
		return x.FullUrl
	}
	return ""
}

func (x *MenuMessage) GetChildren() []*MenuMessage {
	if x != nil {
		return x.Children
	}
	return nil
}

type MenuResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*MenuMessage `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *MenuResponseMessage) Reset() {
	*x = MenuResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuResponseMessage) ProtoMessage() {}

func (x *MenuResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_menu_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuResponseMessage.ProtoReflect.Descriptor instead.
func (*MenuResponseMessage) Descriptor() ([]byte, []int) {
	return file_menu_service_proto_rawDescGZIP(), []int{2}
}

func (x *MenuResponseMessage) GetList() []*MenuMessage {
	if x != nil {
		return x.List
	}
	return nil
}

var File_menu_service_proto protoreflect.FileDescriptor

var file_menu_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x10, 0x0a, 0x0e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xdf, 0x03, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x53,
	0x6c, 0x69, 0x64, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x68, 0x6f,
	0x77, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x54, 0x65, 0x78, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x54, 0x65, 0x78, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x54, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x75, 0x6c, 0x6c, 0x55, 0x72, 0x6c,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x75, 0x6c, 0x6c, 0x55, 0x72, 0x6c, 0x12,
	0x38, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x12, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x47, 0x0a, 0x13, 0x4d, 0x65, 0x6e,
	0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x30, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x32, 0x62, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x53, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e,
	0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x24,
	0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_menu_service_proto_rawDescOnce sync.Once
	file_menu_service_proto_rawDescData = file_menu_service_proto_rawDesc
)

func file_menu_service_proto_rawDescGZIP() []byte {
	file_menu_service_proto_rawDescOnce.Do(func() {
		file_menu_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_menu_service_proto_rawDescData)
	})
	return file_menu_service_proto_rawDescData
}

var file_menu_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_menu_service_proto_goTypes = []interface{}{
	(*MenuReqMessage)(nil),      // 0: menu.service.v1.MenuReqMessage
	(*MenuMessage)(nil),         // 1: menu.service.v1.MenuMessage
	(*MenuResponseMessage)(nil), // 2: menu.service.v1.MenuResponseMessage
}
var file_menu_service_proto_depIdxs = []int32{
	1, // 0: menu.service.v1.MenuMessage.children:type_name -> menu.service.v1.MenuMessage
	1, // 1: menu.service.v1.MenuResponseMessage.list:type_name -> menu.service.v1.MenuMessage
	0, // 2: menu.service.v1.MenuService.MenuList:input_type -> menu.service.v1.MenuReqMessage
	2, // 3: menu.service.v1.MenuService.MenuList:output_type -> menu.service.v1.MenuResponseMessage
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_menu_service_proto_init() }
func file_menu_service_proto_init() {
	if File_menu_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_menu_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuReqMessage); i {
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
		file_menu_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuMessage); i {
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
		file_menu_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuResponseMessage); i {
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
			RawDescriptor: file_menu_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_menu_service_proto_goTypes,
		DependencyIndexes: file_menu_service_proto_depIdxs,
		MessageInfos:      file_menu_service_proto_msgTypes,
	}.Build()
	File_menu_service_proto = out.File
	file_menu_service_proto_rawDesc = nil
	file_menu_service_proto_goTypes = nil
	file_menu_service_proto_depIdxs = nil
}
