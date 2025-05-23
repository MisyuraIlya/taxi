// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: geo.proto

package proto

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

// Existing messages
type UpdateLocationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DriverId      string                 `protobuf:"bytes,1,opt,name=driverId,proto3" json:"driverId,omitempty"`
	Latitude      string                 `protobuf:"bytes,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     string                 `protobuf:"bytes,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateLocationRequest) Reset() {
	*x = UpdateLocationRequest{}
	mi := &file_geo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLocationRequest) ProtoMessage() {}

func (x *UpdateLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLocationRequest.ProtoReflect.Descriptor instead.
func (*UpdateLocationRequest) Descriptor() ([]byte, []int) {
	return file_geo_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateLocationRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *UpdateLocationRequest) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *UpdateLocationRequest) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

func (x *UpdateLocationRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateLocationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateLocationResponse) Reset() {
	*x = UpdateLocationResponse{}
	mi := &file_geo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLocationResponse) ProtoMessage() {}

func (x *UpdateLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLocationResponse.ProtoReflect.Descriptor instead.
func (*UpdateLocationResponse) Descriptor() ([]byte, []int) {
	return file_geo_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateLocationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetLocationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DriverId      string                 `protobuf:"bytes,1,opt,name=driverId,proto3" json:"driverId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLocationRequest) Reset() {
	*x = GetLocationRequest{}
	mi := &file_geo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationRequest) ProtoMessage() {}

func (x *GetLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationRequest.ProtoReflect.Descriptor instead.
func (*GetLocationRequest) Descriptor() ([]byte, []int) {
	return file_geo_proto_rawDescGZIP(), []int{2}
}

func (x *GetLocationRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

type GetLocationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Latitude      string                 `protobuf:"bytes,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     string                 `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLocationResponse) Reset() {
	*x = GetLocationResponse{}
	mi := &file_geo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationResponse) ProtoMessage() {}

func (x *GetLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationResponse.ProtoReflect.Descriptor instead.
func (*GetLocationResponse) Descriptor() ([]byte, []int) {
	return file_geo_proto_rawDescGZIP(), []int{3}
}

func (x *GetLocationResponse) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *GetLocationResponse) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

// NEW: define the messages for finding nearby drivers
type FindDriversRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Latitude      float64                `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     float64                `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Radius        float64                `protobuf:"fixed64,3,opt,name=radius,proto3" json:"radius,omitempty"`
	Limit         uint32                 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindDriversRequest) Reset() {
	*x = FindDriversRequest{}
	mi := &file_geo_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindDriversRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDriversRequest) ProtoMessage() {}

func (x *FindDriversRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geo_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDriversRequest.ProtoReflect.Descriptor instead.
func (*FindDriversRequest) Descriptor() ([]byte, []int) {
	return file_geo_proto_rawDescGZIP(), []int{4}
}

func (x *FindDriversRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *FindDriversRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *FindDriversRequest) GetRadius() float64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *FindDriversRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindDriversRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Driver struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DriverId      string                 `protobuf:"bytes,1,opt,name=driverId,proto3" json:"driverId,omitempty"`
	Latitude      float64                `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     float64                `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Geohash       uint64                 `protobuf:"varint,4,opt,name=geohash,proto3" json:"geohash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Driver) Reset() {
	*x = Driver{}
	mi := &file_geo_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Driver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Driver) ProtoMessage() {}

func (x *Driver) ProtoReflect() protoreflect.Message {
	mi := &file_geo_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Driver.ProtoReflect.Descriptor instead.
func (*Driver) Descriptor() ([]byte, []int) {
	return file_geo_proto_rawDescGZIP(), []int{5}
}

func (x *Driver) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *Driver) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Driver) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Driver) GetGeohash() uint64 {
	if x != nil {
		return x.Geohash
	}
	return 0
}

type FindDriversResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Drivers       []*Driver              `protobuf:"bytes,1,rep,name=drivers,proto3" json:"drivers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindDriversResponse) Reset() {
	*x = FindDriversResponse{}
	mi := &file_geo_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindDriversResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDriversResponse) ProtoMessage() {}

func (x *FindDriversResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geo_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDriversResponse.ProtoReflect.Descriptor instead.
func (*FindDriversResponse) Descriptor() ([]byte, []int) {
	return file_geo_proto_rawDescGZIP(), []int{6}
}

func (x *FindDriversResponse) GetDrivers() []*Driver {
	if x != nil {
		return x.Drivers
	}
	return nil
}

var File_geo_proto protoreflect.FileDescriptor

const file_geo_proto_rawDesc = "" +
	"\n" +
	"\tgeo.proto\x12\x03geo\"\x85\x01\n" +
	"\x15UpdateLocationRequest\x12\x1a\n" +
	"\bdriverId\x18\x01 \x01(\tR\bdriverId\x12\x1a\n" +
	"\blatitude\x18\x02 \x01(\tR\blatitude\x12\x1c\n" +
	"\tlongitude\x18\x03 \x01(\tR\tlongitude\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\"2\n" +
	"\x16UpdateLocationResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"0\n" +
	"\x12GetLocationRequest\x12\x1a\n" +
	"\bdriverId\x18\x01 \x01(\tR\bdriverId\"O\n" +
	"\x13GetLocationResponse\x12\x1a\n" +
	"\blatitude\x18\x01 \x01(\tR\blatitude\x12\x1c\n" +
	"\tlongitude\x18\x02 \x01(\tR\tlongitude\"\x94\x01\n" +
	"\x12FindDriversRequest\x12\x1a\n" +
	"\blatitude\x18\x01 \x01(\x01R\blatitude\x12\x1c\n" +
	"\tlongitude\x18\x02 \x01(\x01R\tlongitude\x12\x16\n" +
	"\x06radius\x18\x03 \x01(\x01R\x06radius\x12\x14\n" +
	"\x05limit\x18\x04 \x01(\rR\x05limit\x12\x16\n" +
	"\x06status\x18\x05 \x01(\tR\x06status\"x\n" +
	"\x06Driver\x12\x1a\n" +
	"\bdriverId\x18\x01 \x01(\tR\bdriverId\x12\x1a\n" +
	"\blatitude\x18\x02 \x01(\x01R\blatitude\x12\x1c\n" +
	"\tlongitude\x18\x03 \x01(\x01R\tlongitude\x12\x18\n" +
	"\ageohash\x18\x04 \x01(\x04R\ageohash\"<\n" +
	"\x13FindDriversResponse\x12%\n" +
	"\adrivers\x18\x01 \x03(\v2\v.geo.DriverR\adrivers2\xdb\x01\n" +
	"\n" +
	"GeoService\x12I\n" +
	"\x0eUpdateLocation\x12\x1a.geo.UpdateLocationRequest\x1a\x1b.geo.UpdateLocationResponse\x12@\n" +
	"\vGetLocation\x12\x17.geo.GetLocationRequest\x1a\x18.geo.GetLocationResponse\x12@\n" +
	"\vFindDrivers\x12\x17.geo.FindDriversRequest\x1a\x18.geo.FindDriversResponseB1Z/github.com/yourusername/geo-service/proto;protob\x06proto3"

var (
	file_geo_proto_rawDescOnce sync.Once
	file_geo_proto_rawDescData []byte
)

func file_geo_proto_rawDescGZIP() []byte {
	file_geo_proto_rawDescOnce.Do(func() {
		file_geo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_geo_proto_rawDesc), len(file_geo_proto_rawDesc)))
	})
	return file_geo_proto_rawDescData
}

var file_geo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_geo_proto_goTypes = []any{
	(*UpdateLocationRequest)(nil),  // 0: geo.UpdateLocationRequest
	(*UpdateLocationResponse)(nil), // 1: geo.UpdateLocationResponse
	(*GetLocationRequest)(nil),     // 2: geo.GetLocationRequest
	(*GetLocationResponse)(nil),    // 3: geo.GetLocationResponse
	(*FindDriversRequest)(nil),     // 4: geo.FindDriversRequest
	(*Driver)(nil),                 // 5: geo.Driver
	(*FindDriversResponse)(nil),    // 6: geo.FindDriversResponse
}
var file_geo_proto_depIdxs = []int32{
	5, // 0: geo.FindDriversResponse.drivers:type_name -> geo.Driver
	0, // 1: geo.GeoService.UpdateLocation:input_type -> geo.UpdateLocationRequest
	2, // 2: geo.GeoService.GetLocation:input_type -> geo.GetLocationRequest
	4, // 3: geo.GeoService.FindDrivers:input_type -> geo.FindDriversRequest
	1, // 4: geo.GeoService.UpdateLocation:output_type -> geo.UpdateLocationResponse
	3, // 5: geo.GeoService.GetLocation:output_type -> geo.GetLocationResponse
	6, // 6: geo.GeoService.FindDrivers:output_type -> geo.FindDriversResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_geo_proto_init() }
func file_geo_proto_init() {
	if File_geo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_geo_proto_rawDesc), len(file_geo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_geo_proto_goTypes,
		DependencyIndexes: file_geo_proto_depIdxs,
		MessageInfos:      file_geo_proto_msgTypes,
	}.Build()
	File_geo_proto = out.File
	file_geo_proto_goTypes = nil
	file_geo_proto_depIdxs = nil
}
