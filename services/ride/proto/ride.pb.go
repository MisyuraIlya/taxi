// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: ride.proto

package ride

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateOrderRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	UserId           string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DriverId         string                 `protobuf:"bytes,2,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	PickupLatitude   float64                `protobuf:"fixed64,3,opt,name=pickup_latitude,json=pickupLatitude,proto3" json:"pickup_latitude,omitempty"`
	PickupLongitude  float64                `protobuf:"fixed64,4,opt,name=pickup_longitude,json=pickupLongitude,proto3" json:"pickup_longitude,omitempty"`
	DropoffLatitude  float64                `protobuf:"fixed64,5,opt,name=dropoff_latitude,json=dropoffLatitude,proto3" json:"dropoff_latitude,omitempty"`
	DropoffLongitude float64                `protobuf:"fixed64,6,opt,name=dropoff_longitude,json=dropoffLongitude,proto3" json:"dropoff_longitude,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_ride_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ride_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_ride_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *CreateOrderRequest) GetPickupLatitude() float64 {
	if x != nil {
		return x.PickupLatitude
	}
	return 0
}

func (x *CreateOrderRequest) GetPickupLongitude() float64 {
	if x != nil {
		return x.PickupLongitude
	}
	return 0
}

func (x *CreateOrderRequest) GetDropoffLatitude() float64 {
	if x != nil {
		return x.DropoffLatitude
	}
	return 0
}

func (x *CreateOrderRequest) GetDropoffLongitude() float64 {
	if x != nil {
		return x.DropoffLongitude
	}
	return 0
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	DriverId      string                 `protobuf:"bytes,2,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_ride_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ride_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_ride_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateOrderResponse) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *CreateOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_ride_proto protoreflect.FileDescriptor

const file_ride_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"ride.proto\x12\x04ride\"\xf6\x01\n" +
	"\x12CreateOrderRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1b\n" +
	"\tdriver_id\x18\x02 \x01(\tR\bdriverId\x12'\n" +
	"\x0fpickup_latitude\x18\x03 \x01(\x01R\x0epickupLatitude\x12)\n" +
	"\x10pickup_longitude\x18\x04 \x01(\x01R\x0fpickupLongitude\x12)\n" +
	"\x10dropoff_latitude\x18\x05 \x01(\x01R\x0fdropoffLatitude\x12+\n" +
	"\x11dropoff_longitude\x18\x06 \x01(\x01R\x10dropoffLongitude\"e\n" +
	"\x13CreateOrderResponse\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x1b\n" +
	"\tdriver_id\x18\x02 \x01(\tR\bdriverId\x12\x16\n" +
	"\x06status\x18\x03 \x01(\tR\x06status2Q\n" +
	"\vRideService\x12B\n" +
	"\vCreateOrder\x12\x18.ride.CreateOrderRequest\x1a\x19.ride.CreateOrderResponseB\fZ\n" +
	"/ride;rideb\x06proto3"

var (
	file_ride_proto_rawDescOnce sync.Once
	file_ride_proto_rawDescData []byte
)

func file_ride_proto_rawDescGZIP() []byte {
	file_ride_proto_rawDescOnce.Do(func() {
		file_ride_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ride_proto_rawDesc), len(file_ride_proto_rawDesc)))
	})
	return file_ride_proto_rawDescData
}

var file_ride_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ride_proto_goTypes = []any{
	(*CreateOrderRequest)(nil),  // 0: ride.CreateOrderRequest
	(*CreateOrderResponse)(nil), // 1: ride.CreateOrderResponse
}
var file_ride_proto_depIdxs = []int32{
	0, // 0: ride.RideService.CreateOrder:input_type -> ride.CreateOrderRequest
	1, // 1: ride.RideService.CreateOrder:output_type -> ride.CreateOrderResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ride_proto_init() }
func file_ride_proto_init() {
	if File_ride_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ride_proto_rawDesc), len(file_ride_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ride_proto_goTypes,
		DependencyIndexes: file_ride_proto_depIdxs,
		MessageInfos:      file_ride_proto_msgTypes,
	}.Build()
	File_ride_proto = out.File
	file_ride_proto_goTypes = nil
	file_ride_proto_depIdxs = nil
}
