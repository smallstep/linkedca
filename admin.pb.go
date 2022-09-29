// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: admin.proto

package linkedca

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Admin_Type int32

const (
	Admin_UNKNOWN     Admin_Type = 0
	Admin_ADMIN       Admin_Type = 1
	Admin_SUPER_ADMIN Admin_Type = 2
)

// Enum value maps for Admin_Type.
var (
	Admin_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "ADMIN",
		2: "SUPER_ADMIN",
	}
	Admin_Type_value = map[string]int32{
		"UNKNOWN":     0,
		"ADMIN":       1,
		"SUPER_ADMIN": 2,
	}
)

func (x Admin_Type) Enum() *Admin_Type {
	p := new(Admin_Type)
	*p = x
	return p
}

func (x Admin_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Admin_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_admin_proto_enumTypes[0].Descriptor()
}

func (Admin_Type) Type() protoreflect.EnumType {
	return &file_admin_proto_enumTypes[0]
}

func (x Admin_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Admin_Type.Descriptor instead.
func (Admin_Type) EnumDescriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{0, 0}
}

type Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorityId   string                 `protobuf:"bytes,2,opt,name=authority_id,json=authorityId,proto3" json:"authority_id,omitempty"`
	Subject       string                 `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	ProvisionerId string                 `protobuf:"bytes,4,opt,name=provisioner_id,json=provisionerId,proto3" json:"provisioner_id,omitempty"`
	Type          Admin_Type             `protobuf:"varint,5,opt,name=type,proto3,enum=linkedca.Admin_Type" json:"type,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DeletedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Admin) Reset() {
	*x = Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{0}
}

func (x *Admin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Admin) GetAuthorityId() string {
	if x != nil {
		return x.AuthorityId
	}
	return ""
}

func (x *Admin) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Admin) GetProvisionerId() string {
	if x != nil {
		return x.ProvisionerId
	}
	return ""
}

func (x *Admin) GetType() Admin_Type {
	if x != nil {
		return x.Type
	}
	return Admin_UNKNOWN
}

func (x *Admin) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Admin) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type AdminList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admins []*Admin `protobuf:"bytes,1,rep,name=admins,proto3" json:"admins,omitempty"`
}

func (x *AdminList) Reset() {
	*x = AdminList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminList) ProtoMessage() {}

func (x *AdminList) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminList.ProtoReflect.Descriptor instead.
func (*AdminList) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{1}
}

func (x *AdminList) GetAdmins() []*Admin {
	if x != nil {
		return x.Admins
	}
	return nil
}

var File_admin_proto protoreflect.FileDescriptor

var file_admin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c,
	0x69, 0x6e, 0x6b, 0x65, 0x64, 0x63, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x05, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x63, 0x61, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2f, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x44, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x55, 0x50, 0x45, 0x52, 0x5f,
	0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x22, 0x34, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x63, 0x61, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x42, 0x15, 0x5a,
	0x13, 0x67, 0x6f, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x73, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x6b,
	0x65, 0x64, 0x63, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_proto_rawDescOnce sync.Once
	file_admin_proto_rawDescData = file_admin_proto_rawDesc
)

func file_admin_proto_rawDescGZIP() []byte {
	file_admin_proto_rawDescOnce.Do(func() {
		file_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_proto_rawDescData)
	})
	return file_admin_proto_rawDescData
}

var file_admin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_admin_proto_goTypes = []interface{}{
	(Admin_Type)(0),               // 0: linkedca.Admin.Type
	(*Admin)(nil),                 // 1: linkedca.Admin
	(*AdminList)(nil),             // 2: linkedca.AdminList
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_admin_proto_depIdxs = []int32{
	0, // 0: linkedca.Admin.type:type_name -> linkedca.Admin.Type
	3, // 1: linkedca.Admin.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: linkedca.Admin.deleted_at:type_name -> google.protobuf.Timestamp
	1, // 3: linkedca.AdminList.admins:type_name -> linkedca.Admin
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_admin_proto_init() }
func file_admin_proto_init() {
	if File_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Admin); i {
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
		file_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminList); i {
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
			RawDescriptor: file_admin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_proto_goTypes,
		DependencyIndexes: file_admin_proto_depIdxs,
		EnumInfos:         file_admin_proto_enumTypes,
		MessageInfos:      file_admin_proto_msgTypes,
	}.Build()
	File_admin_proto = out.File
	file_admin_proto_rawDesc = nil
	file_admin_proto_goTypes = nil
	file_admin_proto_depIdxs = nil
}
