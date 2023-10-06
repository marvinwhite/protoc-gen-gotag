// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: example/example.proto

package example

import (
	_ "github.com/marvinwhite/protoc-gen-gotag/tagger"
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

type Example struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WithNewTags     string  `protobuf:"bytes,1,opt,name=with_new_tags,json=withNewTags,proto3" json:"with_new_tags,omitempty" graphql:"withNewTags,optional"`
	WithNewMultiple string  `protobuf:"bytes,2,opt,name=with_new_multiple,json=withNewMultiple,proto3" json:"with_new_multiple,omitempty" graphql:"withNewTags,optional" xml:"multi,omitempty"`
	ReplaceDefault  *string `protobuf:"bytes,3,opt,name=replace_default,json=replaceDefault,proto3,oneof" json:"replacePrevious"`
	// Types that are assignable to OneOf:
	//
	//	*Example_A
	//	*Example_BJk
	OneOf isExample_OneOf `protobuf_oneof:"one_of" graphql:"withNewTags,optional"`
}

func (x *Example) Reset() {
	*x = Example{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example) ProtoMessage() {}

func (x *Example) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example.ProtoReflect.Descriptor instead.
func (*Example) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{0}
}

func (x *Example) GetWithNewTags() string {
	if x != nil {
		return x.WithNewTags
	}
	return ""
}

func (x *Example) GetWithNewMultiple() string {
	if x != nil {
		return x.WithNewMultiple
	}
	return ""
}

func (x *Example) GetReplaceDefault() string {
	if x != nil && x.ReplaceDefault != nil {
		return *x.ReplaceDefault
	}
	return ""
}

func (m *Example) GetOneOf() isExample_OneOf {
	if m != nil {
		return m.OneOf
	}
	return nil
}

func (x *Example) GetA() string {
	if x, ok := x.GetOneOf().(*Example_A); ok {
		return x.A
	}
	return ""
}

func (x *Example) GetBJk() int32 {
	if x, ok := x.GetOneOf().(*Example_BJk); ok {
		return x.BJk
	}
	return 0
}

type isExample_OneOf interface {
	isExample_OneOf()
}

type Example_A struct {
	A string `protobuf:"bytes,5,opt,name=a,proto3,oneof" json:"A"`
}

type Example_BJk struct {
	BJk int32 `protobuf:"varint,6,opt,name=b_jk,json=bJk,proto3,oneof" json:"b_Jk"`
}

func (*Example_A) isExample_OneOf() {}

func (*Example_BJk) isExample_OneOf() {}

type SecondMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WithNewTags     string `protobuf:"bytes,1,opt,name=with_new_tags,json=withNewTags,proto3" json:"with_new_tags,omitempty" graphql:"withNewTags,optional"`
	WithNewMultiple string `protobuf:"bytes,2,opt,name=with_new_multiple,json=withNewMultiple,proto3" json:"with_new_multiple,omitempty" graphql:"withNewTags,optional" xml:"multi,omitempty"`
	ReplaceDefault  string `protobuf:"bytes,3,opt,name=replace_default,json=replaceDefault,proto3" json:"replacePrevious"`
}

func (x *SecondMessage) Reset() {
	*x = SecondMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecondMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecondMessage) ProtoMessage() {}

func (x *SecondMessage) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecondMessage.ProtoReflect.Descriptor instead.
func (*SecondMessage) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{1}
}

func (x *SecondMessage) GetWithNewTags() string {
	if x != nil {
		return x.WithNewTags
	}
	return ""
}

func (x *SecondMessage) GetWithNewMultiple() string {
	if x != nil {
		return x.WithNewMultiple
	}
	return ""
}

func (x *SecondMessage) GetReplaceDefault() string {
	if x != nil {
		return x.ReplaceDefault
	}
	return ""
}

type ThirdExample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InnerExample *ThirdExample_InnerExample `protobuf:"bytes,1,opt,name=inner_example,json=innerExample,proto3" json:"inner"`
}

func (x *ThirdExample) Reset() {
	*x = ThirdExample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThirdExample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThirdExample) ProtoMessage() {}

func (x *ThirdExample) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThirdExample.ProtoReflect.Descriptor instead.
func (*ThirdExample) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{2}
}

func (x *ThirdExample) GetInnerExample() *ThirdExample_InnerExample {
	if x != nil {
		return x.InnerExample
	}
	return nil
}

type ThirdExample_InnerExample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"yes"`
	Yes int32  `protobuf:"varint,2,opt,name=yes,proto3" json:"id"`
}

func (x *ThirdExample_InnerExample) Reset() {
	*x = ThirdExample_InnerExample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThirdExample_InnerExample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThirdExample_InnerExample) ProtoMessage() {}

func (x *ThirdExample_InnerExample) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThirdExample_InnerExample.ProtoReflect.Descriptor instead.
func (*ThirdExample_InnerExample) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ThirdExample_InnerExample) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ThirdExample_InnerExample) GetYes() int32 {
	if x != nil {
		return x.Yes
	}
	return 0
}

var File_example_example_proto protoreflect.FileDescriptor

var file_example_example_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x03, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x12, 0x47, 0x0a, 0x0d, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0x9a, 0x84, 0x9e, 0x03, 0x1e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x3a, 0x22, 0x77, 0x69, 0x74, 0x68, 0x4e, 0x65, 0x77, 0x54,
	0x61, 0x67, 0x73, 0x2c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x52, 0x0b, 0x77,
	0x69, 0x74, 0x68, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x67, 0x73, 0x12, 0x65, 0x0a, 0x11, 0x77, 0x69,
	0x74, 0x68, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x39, 0x9a, 0x84, 0x9e, 0x03, 0x34, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x71, 0x6c, 0x3a, 0x22, 0x77, 0x69, 0x74, 0x68, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x67, 0x73,
	0x2c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x20, 0x78, 0x6d, 0x6c, 0x3a, 0x22,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x4e, 0x65, 0x77, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x65, 0x12, 0x49, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0x9a, 0x84, 0x9e, 0x03,
	0x16, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x22, 0x48, 0x01, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x01,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x9a, 0x84, 0x9e, 0x03, 0x08, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x41, 0x22, 0x48, 0x00, 0x52, 0x01, 0x61, 0x12, 0x25, 0x0a, 0x04, 0x62,
	0x5f, 0x6a, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x62, 0x5f, 0x4a, 0x6b, 0x22, 0x48, 0x00, 0x52, 0x03, 0x62,
	0x4a, 0x6b, 0x42, 0x2d, 0x0a, 0x06, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x12, 0x23, 0x9a, 0x84,
	0x9e, 0x03, 0x1e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x3a, 0x22, 0x77, 0x69, 0x74, 0x68,
	0x4e, 0x65, 0x77, 0x54, 0x61, 0x67, 0x73, 0x2c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x22, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x85, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x77, 0x69, 0x74, 0x68, 0x5f,
	0x6e, 0x65, 0x77, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23,
	0x9a, 0x84, 0x9e, 0x03, 0x1e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x3a, 0x22, 0x77, 0x69,
	0x74, 0x68, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x67, 0x73, 0x2c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x22, 0x52, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x67, 0x73,
	0x12, 0x65, 0x0a, 0x11, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x39, 0x9a, 0x84, 0x9e,
	0x03, 0x34, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x3a, 0x22, 0x77, 0x69, 0x74, 0x68, 0x4e,
	0x65, 0x77, 0x54, 0x61, 0x67, 0x73, 0x2c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22,
	0x20, 0x78, 0x6d, 0x6c, 0x3a, 0x22, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2c, 0x6f, 0x6d, 0x69, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x4e, 0x65, 0x77, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1b, 0x9a, 0x84, 0x9e, 0x03, 0x16, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x22, 0x52, 0x0e, 0x72,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0xbd, 0x01,
	0x0a, 0x0c, 0x54, 0x68, 0x69, 0x72, 0x64, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x5a,
	0x0a, 0x0d, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x54, 0x68, 0x69, 0x72, 0x64, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x11, 0x9a, 0x84, 0x9e, 0x03, 0x0c,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x52, 0x0c, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x51, 0x0a, 0x0c, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x9a, 0x84, 0x9e, 0x03, 0x0a, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x79, 0x65, 0x73, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x03, 0x79,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0e, 0x9a, 0x84, 0x9e, 0x03, 0x09, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x03, 0x79, 0x65, 0x73, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x72, 0x69, 0x6b,
	0x72, 0x73, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x6f, 0x74, 0x61, 0x67, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_example_proto_rawDescOnce sync.Once
	file_example_example_proto_rawDescData = file_example_example_proto_rawDesc
)

func file_example_example_proto_rawDescGZIP() []byte {
	file_example_example_proto_rawDescOnce.Do(func() {
		file_example_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_example_proto_rawDescData)
	})
	return file_example_example_proto_rawDescData
}

var file_example_example_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_example_example_proto_goTypes = []interface{}{
	(*Example)(nil),                   // 0: example.Example
	(*SecondMessage)(nil),             // 1: example.SecondMessage
	(*ThirdExample)(nil),              // 2: example.ThirdExample
	(*ThirdExample_InnerExample)(nil), // 3: example.ThirdExample.InnerExample
}
var file_example_example_proto_depIdxs = []int32{
	3, // 0: example.ThirdExample.inner_example:type_name -> example.ThirdExample.InnerExample
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_example_proto_init() }
func file_example_example_proto_init() {
	if File_example_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example); i {
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
		file_example_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecondMessage); i {
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
		file_example_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThirdExample); i {
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
		file_example_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThirdExample_InnerExample); i {
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
	file_example_example_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Example_A)(nil),
		(*Example_BJk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_example_proto_goTypes,
		DependencyIndexes: file_example_example_proto_depIdxs,
		MessageInfos:      file_example_example_proto_msgTypes,
	}.Build()
	File_example_example_proto = out.File
	file_example_example_proto_rawDesc = nil
	file_example_example_proto_goTypes = nil
	file_example_example_proto_depIdxs = nil
}
