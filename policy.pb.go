// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: policy.proto

package linkedca

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

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X509 *X509Policy `protobuf:"bytes,1,opt,name=x509,proto3" json:"x509,omitempty"`
	Ssh  *SSHPolicy  `protobuf:"bytes,2,opt,name=ssh,proto3" json:"ssh,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetX509() *X509Policy {
	if x != nil {
		return x.X509
	}
	return nil
}

func (x *Policy) GetSsh() *SSHPolicy {
	if x != nil {
		return x.Ssh
	}
	return nil
}

type X509Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allow              *X509Names `protobuf:"bytes,1,opt,name=allow,proto3" json:"allow,omitempty"`
	Deny               *X509Names `protobuf:"bytes,2,opt,name=deny,proto3" json:"deny,omitempty"`
	AllowWildcardNames bool       `protobuf:"varint,3,opt,name=allow_wildcard_names,json=allowWildcardNames,proto3" json:"allow_wildcard_names,omitempty"`
}

func (x *X509Policy) Reset() {
	*x = X509Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *X509Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X509Policy) ProtoMessage() {}

func (x *X509Policy) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X509Policy.ProtoReflect.Descriptor instead.
func (*X509Policy) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{1}
}

func (x *X509Policy) GetAllow() *X509Names {
	if x != nil {
		return x.Allow
	}
	return nil
}

func (x *X509Policy) GetDeny() *X509Names {
	if x != nil {
		return x.Deny
	}
	return nil
}

func (x *X509Policy) GetAllowWildcardNames() bool {
	if x != nil {
		return x.AllowWildcardNames
	}
	return false
}

type SSHPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host *SSHHostPolicy `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	User *SSHUserPolicy `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SSHPolicy) Reset() {
	*x = SSHPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSHPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHPolicy) ProtoMessage() {}

func (x *SSHPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHPolicy.ProtoReflect.Descriptor instead.
func (*SSHPolicy) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{2}
}

func (x *SSHPolicy) GetHost() *SSHHostPolicy {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *SSHPolicy) GetUser() *SSHUserPolicy {
	if x != nil {
		return x.User
	}
	return nil
}

type SSHHostPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allow *SSHHostNames `protobuf:"bytes,1,opt,name=allow,proto3" json:"allow,omitempty"`
	Deny  *SSHHostNames `protobuf:"bytes,2,opt,name=deny,proto3" json:"deny,omitempty"`
}

func (x *SSHHostPolicy) Reset() {
	*x = SSHHostPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSHHostPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHHostPolicy) ProtoMessage() {}

func (x *SSHHostPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHHostPolicy.ProtoReflect.Descriptor instead.
func (*SSHHostPolicy) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{3}
}

func (x *SSHHostPolicy) GetAllow() *SSHHostNames {
	if x != nil {
		return x.Allow
	}
	return nil
}

func (x *SSHHostPolicy) GetDeny() *SSHHostNames {
	if x != nil {
		return x.Deny
	}
	return nil
}

type SSHUserPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allow *SSHUserNames `protobuf:"bytes,1,opt,name=allow,proto3" json:"allow,omitempty"`
	Deny  *SSHUserNames `protobuf:"bytes,2,opt,name=deny,proto3" json:"deny,omitempty"`
}

func (x *SSHUserPolicy) Reset() {
	*x = SSHUserPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSHUserPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHUserPolicy) ProtoMessage() {}

func (x *SSHUserPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHUserPolicy.ProtoReflect.Descriptor instead.
func (*SSHUserPolicy) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{4}
}

func (x *SSHUserPolicy) GetAllow() *SSHUserNames {
	if x != nil {
		return x.Allow
	}
	return nil
}

func (x *SSHUserPolicy) GetDeny() *SSHUserNames {
	if x != nil {
		return x.Deny
	}
	return nil
}

type X509Names struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dns         []string `protobuf:"bytes,1,rep,name=dns,proto3" json:"dns,omitempty"`
	Ips         []string `protobuf:"bytes,2,rep,name=ips,proto3" json:"ips,omitempty"`
	Emails      []string `protobuf:"bytes,3,rep,name=emails,proto3" json:"emails,omitempty"`
	Uris        []string `protobuf:"bytes,4,rep,name=uris,proto3" json:"uris,omitempty"`
	CommonNames []string `protobuf:"bytes,5,rep,name=common_names,json=commonNames,proto3" json:"common_names,omitempty"`
}

func (x *X509Names) Reset() {
	*x = X509Names{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *X509Names) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X509Names) ProtoMessage() {}

func (x *X509Names) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X509Names.ProtoReflect.Descriptor instead.
func (*X509Names) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{5}
}

func (x *X509Names) GetDns() []string {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *X509Names) GetIps() []string {
	if x != nil {
		return x.Ips
	}
	return nil
}

func (x *X509Names) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *X509Names) GetUris() []string {
	if x != nil {
		return x.Uris
	}
	return nil
}

func (x *X509Names) GetCommonNames() []string {
	if x != nil {
		return x.CommonNames
	}
	return nil
}

type SSHHostNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dns        []string `protobuf:"bytes,1,rep,name=dns,proto3" json:"dns,omitempty"`
	Ips        []string `protobuf:"bytes,2,rep,name=ips,proto3" json:"ips,omitempty"`
	Principals []string `protobuf:"bytes,3,rep,name=principals,proto3" json:"principals,omitempty"`
}

func (x *SSHHostNames) Reset() {
	*x = SSHHostNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSHHostNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHHostNames) ProtoMessage() {}

func (x *SSHHostNames) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHHostNames.ProtoReflect.Descriptor instead.
func (*SSHHostNames) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{6}
}

func (x *SSHHostNames) GetDns() []string {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *SSHHostNames) GetIps() []string {
	if x != nil {
		return x.Ips
	}
	return nil
}

func (x *SSHHostNames) GetPrincipals() []string {
	if x != nil {
		return x.Principals
	}
	return nil
}

type SSHUserNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emails     []string `protobuf:"bytes,1,rep,name=emails,proto3" json:"emails,omitempty"`
	Principals []string `protobuf:"bytes,2,rep,name=principals,proto3" json:"principals,omitempty"`
}

func (x *SSHUserNames) Reset() {
	*x = SSHUserNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSHUserNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHUserNames) ProtoMessage() {}

func (x *SSHUserNames) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHUserNames.ProtoReflect.Descriptor instead.
func (*SSHUserNames) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{7}
}

func (x *SSHUserNames) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *SSHUserNames) GetPrincipals() []string {
	if x != nil {
		return x.Principals
	}
	return nil
}

var File_policy_proto protoreflect.FileDescriptor

var file_policy_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x63, 0x61, 0x22, 0x59, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x28, 0x0a, 0x04, 0x78, 0x35, 0x30, 0x39, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x63, 0x61, 0x2e, 0x58, 0x35, 0x30, 0x39,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x04, 0x78, 0x35, 0x30, 0x39, 0x12, 0x25, 0x0a, 0x03,
	0x73, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x65, 0x64, 0x63, 0x61, 0x2e, 0x53, 0x53, 0x48, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x03,
	0x73, 0x73, 0x68, 0x22, 0x92, 0x01, 0x0a, 0x0a, 0x58, 0x35, 0x30, 0x39, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x63, 0x61, 0x2e, 0x58, 0x35, 0x30,
	0x39, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x27, 0x0a,
	0x04, 0x64, 0x65, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x69,
	0x6e, 0x6b, 0x65, 0x64, 0x63, 0x61, 0x2e, 0x58, 0x35, 0x30, 0x39, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x52, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x57, 0x69, 0x6c, 0x64, 0x63,
	0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x65, 0x0a, 0x09, 0x53, 0x53, 0x48, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2b, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x63, 0x61, 0x2e, 0x53,
	0x53, 0x48, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x63, 0x61, 0x2e, 0x53, 0x53, 0x48, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x69, 0x0a, 0x0d, 0x53, 0x53, 0x48, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x2c, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x63, 0x61, 0x2e, 0x53, 0x53, 0x48, 0x48, 0x6f,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x2a,
	0x0a, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c,
	0x69, 0x6e, 0x6b, 0x65, 0x64, 0x63, 0x61, 0x2e, 0x53, 0x53, 0x48, 0x48, 0x6f, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x22, 0x69, 0x0a, 0x0d, 0x53, 0x53,
	0x48, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x65, 0x64, 0x63, 0x61, 0x2e, 0x53, 0x53, 0x48, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x65, 0x6e,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x63, 0x61, 0x2e, 0x53, 0x53, 0x48, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52,
	0x04, 0x64, 0x65, 0x6e, 0x79, 0x22, 0x7e, 0x0a, 0x09, 0x58, 0x35, 0x30, 0x39, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x64, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x69, 0x70, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x72, 0x69, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72,
	0x69, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x0c, 0x53, 0x53, 0x48, 0x48, 0x6f, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x03, 0x64, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x70, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x70, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x22, 0x46, 0x0a, 0x0c, 0x53, 0x53, 0x48,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x73, 0x42, 0x15, 0x5a, 0x13, 0x67, 0x6f, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x73, 0x6d, 0x2f,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x63, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_policy_proto_rawDescOnce sync.Once
	file_policy_proto_rawDescData = file_policy_proto_rawDesc
)

func file_policy_proto_rawDescGZIP() []byte {
	file_policy_proto_rawDescOnce.Do(func() {
		file_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_policy_proto_rawDescData)
	})
	return file_policy_proto_rawDescData
}

var file_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_policy_proto_goTypes = []interface{}{
	(*Policy)(nil),        // 0: linkedca.Policy
	(*X509Policy)(nil),    // 1: linkedca.X509Policy
	(*SSHPolicy)(nil),     // 2: linkedca.SSHPolicy
	(*SSHHostPolicy)(nil), // 3: linkedca.SSHHostPolicy
	(*SSHUserPolicy)(nil), // 4: linkedca.SSHUserPolicy
	(*X509Names)(nil),     // 5: linkedca.X509Names
	(*SSHHostNames)(nil),  // 6: linkedca.SSHHostNames
	(*SSHUserNames)(nil),  // 7: linkedca.SSHUserNames
}
var file_policy_proto_depIdxs = []int32{
	1,  // 0: linkedca.Policy.x509:type_name -> linkedca.X509Policy
	2,  // 1: linkedca.Policy.ssh:type_name -> linkedca.SSHPolicy
	5,  // 2: linkedca.X509Policy.allow:type_name -> linkedca.X509Names
	5,  // 3: linkedca.X509Policy.deny:type_name -> linkedca.X509Names
	3,  // 4: linkedca.SSHPolicy.host:type_name -> linkedca.SSHHostPolicy
	4,  // 5: linkedca.SSHPolicy.user:type_name -> linkedca.SSHUserPolicy
	6,  // 6: linkedca.SSHHostPolicy.allow:type_name -> linkedca.SSHHostNames
	6,  // 7: linkedca.SSHHostPolicy.deny:type_name -> linkedca.SSHHostNames
	7,  // 8: linkedca.SSHUserPolicy.allow:type_name -> linkedca.SSHUserNames
	7,  // 9: linkedca.SSHUserPolicy.deny:type_name -> linkedca.SSHUserNames
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_policy_proto_init() }
func file_policy_proto_init() {
	if File_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X509Policy); i {
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
		file_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSHPolicy); i {
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
		file_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSHHostPolicy); i {
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
		file_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSHUserPolicy); i {
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
		file_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X509Names); i {
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
		file_policy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSHHostNames); i {
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
		file_policy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSHUserNames); i {
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
			RawDescriptor: file_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_policy_proto_goTypes,
		DependencyIndexes: file_policy_proto_depIdxs,
		MessageInfos:      file_policy_proto_msgTypes,
	}.Build()
	File_policy_proto = out.File
	file_policy_proto_rawDesc = nil
	file_policy_proto_goTypes = nil
	file_policy_proto_depIdxs = nil
}
