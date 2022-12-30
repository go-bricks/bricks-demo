// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: api/workshop.proto

package workshop

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CarBody int32

const (
	Car_SEDAN     CarBody = 0
	Car_PHAETON   CarBody = 1
	Car_HATCHBACK CarBody = 2
)

// Enum value maps for CarBody.
var (
	CarBody_name = map[int32]string{
		0: "SEDAN",
		1: "PHAETON",
		2: "HATCHBACK",
	}
	CarBody_value = map[string]int32{
		"SEDAN":     0,
		"PHAETON":   1,
		"HATCHBACK": 2,
	}
)

func (x CarBody) Enum() *CarBody {
	p := new(CarBody)
	*p = x
	return p
}

func (x CarBody) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CarBody) Descriptor() protoreflect.EnumDescriptor {
	return file_api_workshop_proto_enumTypes[0].Descriptor()
}

func (CarBody) Type() protoreflect.EnumType {
	return &file_api_workshop_proto_enumTypes[0]
}

func (x CarBody) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CarBody.Descriptor instead.
func (CarBody) EnumDescriptor() ([]byte, []int) {
	return file_api_workshop_proto_rawDescGZIP(), []int{0, 0}
}

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number    string  `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Owner     string  `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	BodyStyle CarBody `protobuf:"varint,3,opt,name=body_style,json=bodyStyle,proto3,enum=demo.workshop.CarBody" json:"body_style,omitempty"`
	Color     string  `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_workshop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_api_workshop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_api_workshop_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Car) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Car) GetBodyStyle() CarBody {
	if x != nil {
		return x.BodyStyle
	}
	return Car_SEDAN
}

func (x *Car) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type PaintCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarNumber    string `protobuf:"bytes,1,opt,name=car_number,json=carNumber,proto3" json:"car_number,omitempty"`
	DesiredColor string `protobuf:"bytes,2,opt,name=desired_color,json=desiredColor,proto3" json:"desired_color,omitempty"`
}

func (x *PaintCarRequest) Reset() {
	*x = PaintCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_workshop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaintCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaintCarRequest) ProtoMessage() {}

func (x *PaintCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_workshop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaintCarRequest.ProtoReflect.Descriptor instead.
func (*PaintCarRequest) Descriptor() ([]byte, []int) {
	return file_api_workshop_proto_rawDescGZIP(), []int{1}
}

func (x *PaintCarRequest) GetCarNumber() string {
	if x != nil {
		return x.CarNumber
	}
	return ""
}

func (x *PaintCarRequest) GetDesiredColor() string {
	if x != nil {
		return x.DesiredColor
	}
	return ""
}

type PaintFinishedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarNumber    string `protobuf:"bytes,1,opt,name=car_number,json=carNumber,proto3" json:"car_number,omitempty"`
	DesiredColor string `protobuf:"bytes,2,opt,name=desired_color,json=desiredColor,proto3" json:"desired_color,omitempty"`
}

func (x *PaintFinishedRequest) Reset() {
	*x = PaintFinishedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_workshop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaintFinishedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaintFinishedRequest) ProtoMessage() {}

func (x *PaintFinishedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_workshop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaintFinishedRequest.ProtoReflect.Descriptor instead.
func (*PaintFinishedRequest) Descriptor() ([]byte, []int) {
	return file_api_workshop_proto_rawDescGZIP(), []int{2}
}

func (x *PaintFinishedRequest) GetCarNumber() string {
	if x != nil {
		return x.CarNumber
	}
	return ""
}

func (x *PaintFinishedRequest) GetDesiredColor() string {
	if x != nil {
		return x.DesiredColor
	}
	return ""
}

type RetrieveCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarNumber string `protobuf:"bytes,1,opt,name=car_number,json=carNumber,proto3" json:"car_number,omitempty"`
}

func (x *RetrieveCarRequest) Reset() {
	*x = RetrieveCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_workshop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveCarRequest) ProtoMessage() {}

func (x *RetrieveCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_workshop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveCarRequest.ProtoReflect.Descriptor instead.
func (*RetrieveCarRequest) Descriptor() ([]byte, []int) {
	return file_api_workshop_proto_rawDescGZIP(), []int{3}
}

func (x *RetrieveCarRequest) GetCarNumber() string {
	if x != nil {
		return x.CarNumber
	}
	return ""
}

var File_api_workshop_proto protoreflect.FileDescriptor

var file_api_workshop_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x68, 0x6f, 0x70, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0,
	0x01, 0x0a, 0x03, 0x43, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x73, 0x74, 0x79,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x43, 0x61, 0x72, 0x2e, 0x62, 0x6f, 0x64,
	0x79, 0x52, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x22, 0x2d, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x45,
	0x44, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x48, 0x41, 0x45, 0x54, 0x4f, 0x4e,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x41, 0x54, 0x43, 0x48, 0x42, 0x41, 0x43, 0x4b, 0x10,
	0x02, 0x22, 0x55, 0x0a, 0x0f, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x69,
	0x72, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x5a, 0x0a, 0x14, 0x50, 0x61, 0x69, 0x6e,
	0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x33, 0x0a, 0x12, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61,
	0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x61, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x8f, 0x03, 0x0a, 0x08, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x55, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x43, 0x61, 0x72, 0x12, 0x12, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x43, 0x61, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x63, 0x61, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x73, 0x0a,
	0x08, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x12, 0x1e, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x43,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x1a, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x63, 0x61, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x61,
	0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x7d, 0x2f, 0x70, 0x61, 0x69, 0x6e, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x43, 0x61,
	0x72, 0x12, 0x21, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x43, 0x61, 0x72, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x63,
	0x61, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x7d,
	0x12, 0x49, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x23,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x50,
	0x61, 0x69, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_workshop_proto_rawDescOnce sync.Once
	file_api_workshop_proto_rawDescData = file_api_workshop_proto_rawDesc
)

func file_api_workshop_proto_rawDescGZIP() []byte {
	file_api_workshop_proto_rawDescOnce.Do(func() {
		file_api_workshop_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_workshop_proto_rawDescData)
	})
	return file_api_workshop_proto_rawDescData
}

var file_api_workshop_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_workshop_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_workshop_proto_goTypes = []interface{}{
	(CarBody)(0),                 // 0: demo.workshop.Car.body
	(*Car)(nil),                  // 1: demo.workshop.Car
	(*PaintCarRequest)(nil),      // 2: demo.workshop.PaintCarRequest
	(*PaintFinishedRequest)(nil), // 3: demo.workshop.PaintFinishedRequest
	(*RetrieveCarRequest)(nil),   // 4: demo.workshop.RetrieveCarRequest
	(*empty.Empty)(nil),          // 5: google.protobuf.Empty
}
var file_api_workshop_proto_depIdxs = []int32{
	0, // 0: demo.workshop.Car.body_style:type_name -> demo.workshop.Car.body
	1, // 1: demo.workshop.Workshop.AcceptCar:input_type -> demo.workshop.Car
	2, // 2: demo.workshop.Workshop.PaintCar:input_type -> demo.workshop.PaintCarRequest
	4, // 3: demo.workshop.Workshop.RetrieveCar:input_type -> demo.workshop.RetrieveCarRequest
	3, // 4: demo.workshop.Workshop.CarPainted:input_type -> demo.workshop.PaintFinishedRequest
	5, // 5: demo.workshop.Workshop.AcceptCar:output_type -> google.protobuf.Empty
	5, // 6: demo.workshop.Workshop.PaintCar:output_type -> google.protobuf.Empty
	1, // 7: demo.workshop.Workshop.RetrieveCar:output_type -> demo.workshop.Car
	5, // 8: demo.workshop.Workshop.CarPainted:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_workshop_proto_init() }
func file_api_workshop_proto_init() {
	if File_api_workshop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_workshop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_api_workshop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaintCarRequest); i {
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
		file_api_workshop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaintFinishedRequest); i {
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
		file_api_workshop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveCarRequest); i {
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
			RawDescriptor: file_api_workshop_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_workshop_proto_goTypes,
		DependencyIndexes: file_api_workshop_proto_depIdxs,
		EnumInfos:         file_api_workshop_proto_enumTypes,
		MessageInfos:      file_api_workshop_proto_msgTypes,
	}.Build()
	File_api_workshop_proto = out.File
	file_api_workshop_proto_rawDesc = nil
	file_api_workshop_proto_goTypes = nil
	file_api_workshop_proto_depIdxs = nil
}
