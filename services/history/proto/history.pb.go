// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: history.proto

package history

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

// Request message for creating a history record.
type CreateHistoryRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	UserId   string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	DriverId string                 `protobuf:"bytes,2,opt,name=driverId,proto3" json:"driverId,omitempty"`
	// Timestamps in RFC3339 format.
	CreatedAt     string `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ClosedAt      string `protobuf:"bytes,4,opt,name=closedAt,proto3" json:"closedAt,omitempty"`
	From          string `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	To            string `protobuf:"bytes,6,opt,name=to,proto3" json:"to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateHistoryRequest) Reset() {
	*x = CreateHistoryRequest{}
	mi := &file_history_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHistoryRequest) ProtoMessage() {}

func (x *CreateHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHistoryRequest.ProtoReflect.Descriptor instead.
func (*CreateHistoryRequest) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHistoryRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateHistoryRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *CreateHistoryRequest) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CreateHistoryRequest) GetClosedAt() string {
	if x != nil {
		return x.ClosedAt
	}
	return ""
}

func (x *CreateHistoryRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CreateHistoryRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

// Response message for CreateHistory.
type CreateHistoryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateHistoryResponse) Reset() {
	*x = CreateHistoryResponse{}
	mi := &file_history_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHistoryResponse) ProtoMessage() {}

func (x *CreateHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHistoryResponse.ProtoReflect.Descriptor instead.
func (*CreateHistoryResponse) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHistoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request message for retrieving history records.
type GetHistoriesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetHistoriesRequest) Reset() {
	*x = GetHistoriesRequest{}
	mi := &file_history_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHistoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoriesRequest) ProtoMessage() {}

func (x *GetHistoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoriesRequest.ProtoReflect.Descriptor instead.
func (*GetHistoriesRequest) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{2}
}

// Message representing a single history record.
type HistoryRecord struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	DriverId      string                 `protobuf:"bytes,2,opt,name=driverId,proto3" json:"driverId,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ClosedAt      string                 `protobuf:"bytes,4,opt,name=closedAt,proto3" json:"closedAt,omitempty"`
	From          string                 `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	To            string                 `protobuf:"bytes,6,opt,name=to,proto3" json:"to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HistoryRecord) Reset() {
	*x = HistoryRecord{}
	mi := &file_history_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HistoryRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryRecord) ProtoMessage() {}

func (x *HistoryRecord) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryRecord.ProtoReflect.Descriptor instead.
func (*HistoryRecord) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{3}
}

func (x *HistoryRecord) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *HistoryRecord) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *HistoryRecord) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *HistoryRecord) GetClosedAt() string {
	if x != nil {
		return x.ClosedAt
	}
	return ""
}

func (x *HistoryRecord) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *HistoryRecord) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

// Response message for GetHistories.
type GetHistoriesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Histories     []*HistoryRecord       `protobuf:"bytes,1,rep,name=histories,proto3" json:"histories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetHistoriesResponse) Reset() {
	*x = GetHistoriesResponse{}
	mi := &file_history_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHistoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoriesResponse) ProtoMessage() {}

func (x *GetHistoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoriesResponse.ProtoReflect.Descriptor instead.
func (*GetHistoriesResponse) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{4}
}

func (x *GetHistoriesResponse) GetHistories() []*HistoryRecord {
	if x != nil {
		return x.Histories
	}
	return nil
}

var File_history_proto protoreflect.FileDescriptor

const file_history_proto_rawDesc = "" +
	"\n" +
	"\rhistory.proto\x12\ahistory\"\xa8\x01\n" +
	"\x14CreateHistoryRequest\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\tR\x06userId\x12\x1a\n" +
	"\bdriverId\x18\x02 \x01(\tR\bdriverId\x12\x1c\n" +
	"\tcreatedAt\x18\x03 \x01(\tR\tcreatedAt\x12\x1a\n" +
	"\bclosedAt\x18\x04 \x01(\tR\bclosedAt\x12\x12\n" +
	"\x04from\x18\x05 \x01(\tR\x04from\x12\x0e\n" +
	"\x02to\x18\x06 \x01(\tR\x02to\"1\n" +
	"\x15CreateHistoryResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"\x15\n" +
	"\x13GetHistoriesRequest\"\xa1\x01\n" +
	"\rHistoryRecord\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\tR\x06userId\x12\x1a\n" +
	"\bdriverId\x18\x02 \x01(\tR\bdriverId\x12\x1c\n" +
	"\tcreatedAt\x18\x03 \x01(\tR\tcreatedAt\x12\x1a\n" +
	"\bclosedAt\x18\x04 \x01(\tR\bclosedAt\x12\x12\n" +
	"\x04from\x18\x05 \x01(\tR\x04from\x12\x0e\n" +
	"\x02to\x18\x06 \x01(\tR\x02to\"L\n" +
	"\x14GetHistoriesResponse\x124\n" +
	"\thistories\x18\x01 \x03(\v2\x16.history.HistoryRecordR\thistories2\xad\x01\n" +
	"\x0eHistoryService\x12N\n" +
	"\rCreateHistory\x12\x1d.history.CreateHistoryRequest\x1a\x1e.history.CreateHistoryResponse\x12K\n" +
	"\fGetHistories\x12\x1c.history.GetHistoriesRequest\x1a\x1d.history.GetHistoriesResponseB\x1fZ\x1dhistory-service/proto;historyb\x06proto3"

var (
	file_history_proto_rawDescOnce sync.Once
	file_history_proto_rawDescData []byte
)

func file_history_proto_rawDescGZIP() []byte {
	file_history_proto_rawDescOnce.Do(func() {
		file_history_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_history_proto_rawDesc), len(file_history_proto_rawDesc)))
	})
	return file_history_proto_rawDescData
}

var file_history_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_history_proto_goTypes = []any{
	(*CreateHistoryRequest)(nil),  // 0: history.CreateHistoryRequest
	(*CreateHistoryResponse)(nil), // 1: history.CreateHistoryResponse
	(*GetHistoriesRequest)(nil),   // 2: history.GetHistoriesRequest
	(*HistoryRecord)(nil),         // 3: history.HistoryRecord
	(*GetHistoriesResponse)(nil),  // 4: history.GetHistoriesResponse
}
var file_history_proto_depIdxs = []int32{
	3, // 0: history.GetHistoriesResponse.histories:type_name -> history.HistoryRecord
	0, // 1: history.HistoryService.CreateHistory:input_type -> history.CreateHistoryRequest
	2, // 2: history.HistoryService.GetHistories:input_type -> history.GetHistoriesRequest
	1, // 3: history.HistoryService.CreateHistory:output_type -> history.CreateHistoryResponse
	4, // 4: history.HistoryService.GetHistories:output_type -> history.GetHistoriesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_history_proto_init() }
func file_history_proto_init() {
	if File_history_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_history_proto_rawDesc), len(file_history_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_history_proto_goTypes,
		DependencyIndexes: file_history_proto_depIdxs,
		MessageInfos:      file_history_proto_msgTypes,
	}.Build()
	File_history_proto = out.File
	file_history_proto_goTypes = nil
	file_history_proto_depIdxs = nil
}
