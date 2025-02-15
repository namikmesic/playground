// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.2
// source: proto/telemetry.proto

package telemetry

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

type TelemetryData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HostId        string                 `protobuf:"bytes,1,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	MetricName    string                 `protobuf:"bytes,2,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	Value         float32                `protobuf:"fixed32,3,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp     string                 `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TelemetryData) Reset() {
	*x = TelemetryData{}
	mi := &file_proto_telemetry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TelemetryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryData) ProtoMessage() {}

func (x *TelemetryData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_telemetry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryData.ProtoReflect.Descriptor instead.
func (*TelemetryData) Descriptor() ([]byte, []int) {
	return file_proto_telemetry_proto_rawDescGZIP(), []int{0}
}

func (x *TelemetryData) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *TelemetryData) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *TelemetryData) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TelemetryData) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type TelemetryBatch struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*TelemetryData       `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TelemetryBatch) Reset() {
	*x = TelemetryBatch{}
	mi := &file_proto_telemetry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TelemetryBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryBatch) ProtoMessage() {}

func (x *TelemetryBatch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_telemetry_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryBatch.ProtoReflect.Descriptor instead.
func (*TelemetryBatch) Descriptor() ([]byte, []int) {
	return file_proto_telemetry_proto_rawDescGZIP(), []int{1}
}

func (x *TelemetryBatch) GetItems() []*TelemetryData {
	if x != nil {
		return x.Items
	}
	return nil
}

type BatchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	StatusCode    int32                  `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchResponse) Reset() {
	*x = BatchResponse{}
	mi := &file_proto_telemetry_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResponse) ProtoMessage() {}

func (x *BatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_telemetry_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchResponse.ProtoReflect.Descriptor instead.
func (*BatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_telemetry_proto_rawDescGZIP(), []int{2}
}

func (x *BatchResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BatchResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

type FetchTelemetryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HostId        string                 `protobuf:"bytes,1,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	MetricName    string                 `protobuf:"bytes,2,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	StartTime     string                 `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       string                 `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchTelemetryRequest) Reset() {
	*x = FetchTelemetryRequest{}
	mi := &file_proto_telemetry_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchTelemetryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchTelemetryRequest) ProtoMessage() {}

func (x *FetchTelemetryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_telemetry_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchTelemetryRequest.ProtoReflect.Descriptor instead.
func (*FetchTelemetryRequest) Descriptor() ([]byte, []int) {
	return file_proto_telemetry_proto_rawDescGZIP(), []int{3}
}

func (x *FetchTelemetryRequest) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *FetchTelemetryRequest) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *FetchTelemetryRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *FetchTelemetryRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type FetchTelemetryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*TelemetryData       `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	StatusCode    int32                  `protobuf:"varint,3,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message       string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchTelemetryResponse) Reset() {
	*x = FetchTelemetryResponse{}
	mi := &file_proto_telemetry_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchTelemetryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchTelemetryResponse) ProtoMessage() {}

func (x *FetchTelemetryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_telemetry_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchTelemetryResponse.ProtoReflect.Descriptor instead.
func (*FetchTelemetryResponse) Descriptor() ([]byte, []int) {
	return file_proto_telemetry_proto_rawDescGZIP(), []int{4}
}

func (x *FetchTelemetryResponse) GetItems() []*TelemetryData {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *FetchTelemetryResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *FetchTelemetryResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FetchTelemetryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_telemetry_proto protoreflect.FileDescriptor

var file_proto_telemetry_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x22, 0x7d, 0x0a, 0x0d, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x40, 0x0a, 0x0e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x4a, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x8b, 0x01, 0x0a, 0x15, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xab, 0x01,
	0x0a, 0x16, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xbf, 0x01, 0x0a, 0x10,
	0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4f, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x1a, 0x18, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x5a, 0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x20, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x6d, 0x69,
	0x6b, 0x6d, 0x65, 0x73, 0x69, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_telemetry_proto_rawDescOnce sync.Once
	file_proto_telemetry_proto_rawDescData = file_proto_telemetry_proto_rawDesc
)

func file_proto_telemetry_proto_rawDescGZIP() []byte {
	file_proto_telemetry_proto_rawDescOnce.Do(func() {
		file_proto_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_telemetry_proto_rawDescData)
	})
	return file_proto_telemetry_proto_rawDescData
}

var file_proto_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_telemetry_proto_goTypes = []any{
	(*TelemetryData)(nil),          // 0: telemetry.TelemetryData
	(*TelemetryBatch)(nil),         // 1: telemetry.TelemetryBatch
	(*BatchResponse)(nil),          // 2: telemetry.BatchResponse
	(*FetchTelemetryRequest)(nil),  // 3: telemetry.FetchTelemetryRequest
	(*FetchTelemetryResponse)(nil), // 4: telemetry.FetchTelemetryResponse
}
var file_proto_telemetry_proto_depIdxs = []int32{
	0, // 0: telemetry.TelemetryBatch.items:type_name -> telemetry.TelemetryData
	0, // 1: telemetry.FetchTelemetryResponse.items:type_name -> telemetry.TelemetryData
	1, // 2: telemetry.TelemetryService.StreamTelemetryBatch:input_type -> telemetry.TelemetryBatch
	3, // 3: telemetry.TelemetryService.FetchTelemetryBatch:input_type -> telemetry.FetchTelemetryRequest
	2, // 4: telemetry.TelemetryService.StreamTelemetryBatch:output_type -> telemetry.BatchResponse
	4, // 5: telemetry.TelemetryService.FetchTelemetryBatch:output_type -> telemetry.FetchTelemetryResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_telemetry_proto_init() }
func file_proto_telemetry_proto_init() {
	if File_proto_telemetry_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_telemetry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_telemetry_proto_goTypes,
		DependencyIndexes: file_proto_telemetry_proto_depIdxs,
		MessageInfos:      file_proto_telemetry_proto_msgTypes,
	}.Build()
	File_proto_telemetry_proto = out.File
	file_proto_telemetry_proto_rawDesc = nil
	file_proto_telemetry_proto_goTypes = nil
	file_proto_telemetry_proto_depIdxs = nil
}
