// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: todo.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/zYros90/go-boilerplate-v1/api/third_party/jsonschema"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type CreateTodoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo     string                 `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	DueAt    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=due_at,json=dueAt,proto3" json:"due_at,omitempty"`
	NotifyAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=notify_at,json=notifyAt,proto3" json:"notify_at,omitempty"`
}

func (x *CreateTodoReq) Reset() {
	*x = CreateTodoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoReq) ProtoMessage() {}

func (x *CreateTodoReq) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoReq.ProtoReflect.Descriptor instead.
func (*CreateTodoReq) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTodoReq) GetTodo() string {
	if x != nil {
		return x.Todo
	}
	return ""
}

func (x *CreateTodoReq) GetDueAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DueAt
	}
	return nil
}

func (x *CreateTodoReq) GetNotifyAt() *timestamppb.Timestamp {
	if x != nil {
		return x.NotifyAt
	}
	return nil
}

type UpdateTodoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId   string                 `protobuf:"bytes,1,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
	Todo     string                 `protobuf:"bytes,2,opt,name=todo,proto3" json:"todo,omitempty"`
	DueAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=due_at,json=dueAt,proto3" json:"due_at,omitempty"`
	NotifyAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=notify_at,json=notifyAt,proto3" json:"notify_at,omitempty"`
}

func (x *UpdateTodoReq) Reset() {
	*x = UpdateTodoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoReq) ProtoMessage() {}

func (x *UpdateTodoReq) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoReq.ProtoReflect.Descriptor instead.
func (*UpdateTodoReq) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateTodoReq) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

func (x *UpdateTodoReq) GetTodo() string {
	if x != nil {
		return x.Todo
	}
	return ""
}

func (x *UpdateTodoReq) GetDueAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DueAt
	}
	return nil
}

func (x *UpdateTodoReq) GetNotifyAt() *timestamppb.Timestamp {
	if x != nil {
		return x.NotifyAt
	}
	return nil
}

type DeleteTodoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId string `protobuf:"bytes,1,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
}

func (x *DeleteTodoReq) Reset() {
	*x = DeleteTodoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoReq) ProtoMessage() {}

func (x *DeleteTodoReq) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoReq.ProtoReflect.Descriptor instead.
func (*DeleteTodoReq) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteTodoReq) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

type GetTodoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId string `protobuf:"bytes,1,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
}

func (x *GetTodoReq) Reset() {
	*x = GetTodoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoReq) ProtoMessage() {}

func (x *GetTodoReq) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoReq.ProtoReflect.Descriptor instead.
func (*GetTodoReq) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{3}
}

func (x *GetTodoReq) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

type TodoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId    string                 `protobuf:"bytes,1,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
	Todo      string                 `protobuf:"bytes,2,opt,name=todo,proto3" json:"todo,omitempty"`
	DueAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=due_at,json=dueAt,proto3" json:"due_at,omitempty"`
	NotifyAt  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=notify_at,json=notifyAt,proto3" json:"notify_at,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *TodoResp) Reset() {
	*x = TodoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoResp) ProtoMessage() {}

func (x *TodoResp) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoResp.ProtoReflect.Descriptor instead.
func (*TodoResp) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{4}
}

func (x *TodoResp) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

func (x *TodoResp) GetTodo() string {
	if x != nil {
		return x.Todo
	}
	return ""
}

func (x *TodoResp) GetDueAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DueAt
	}
	return nil
}

func (x *TodoResp) GetNotifyAt() *timestamppb.Timestamp {
	if x != nil {
		return x.NotifyAt
	}
	return nil
}

func (x *TodoResp) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TodoResp) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_todo_proto protoreflect.FileDescriptor

var file_todo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70,
	0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a,
	0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x31, 0x0a, 0x06, 0x64, 0x75,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x64, 0x75, 0x65, 0x41, 0x74, 0x12, 0x37, 0x0a,
	0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x41, 0x74, 0x3a, 0x05, 0xba, 0x46, 0x02, 0x20, 0x01, 0x22, 0xc1, 0x01,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x12,
	0x20, 0x0a, 0x07, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x31,
	0x0a, 0x06, 0x64, 0x75, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x64, 0x75, 0x65, 0x41,
	0x74, 0x12, 0x37, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x41, 0x74, 0x3a, 0x05, 0xba, 0x46, 0x02, 0x20,
	0x01, 0x22, 0x31, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x20, 0x0a, 0x07, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x74, 0x6f,
	0x64, 0x6f, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x20, 0x0a, 0x07, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x74, 0x6f,
	0x64, 0x6f, 0x49, 0x64, 0x22, 0xbf, 0x02, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x20, 0x0a, 0x07, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x74, 0x6f, 0x64,
	0x6f, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f,
	0x12, 0x31, 0x0a, 0x06, 0x64, 0x75, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x64, 0x75,
	0x65, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x41, 0x74, 0x12, 0x43, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0xb2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x43, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xb2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xc6, 0x02, 0x0a, 0x07, 0x54, 0x6f, 0x64, 0x6f, 0x53,
	0x76, 0x63, 0x12, 0x50, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76,
	0x31, 0x3a, 0x01, 0x2a, 0x12, 0x50, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x1a, 0x08, 0x2f, 0x74, 0x6f, 0x64, 0x6f,
	0x2f, 0x76, 0x31, 0x3a, 0x01, 0x2a, 0x12, 0x47, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x10, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x12,
	0x4e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x2a, 0x08, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x42,
	0x40, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x3b, 0x50,
	0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x59,
	0x72, 0x6f, 0x73, 0x39, 0x30, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2d, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData = file_todo_proto_rawDesc
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_proto_rawDescData)
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_todo_proto_goTypes = []interface{}{
	(*CreateTodoReq)(nil),         // 0: api.user.v1.CreateTodoReq
	(*UpdateTodoReq)(nil),         // 1: api.user.v1.UpdateTodoReq
	(*DeleteTodoReq)(nil),         // 2: api.user.v1.DeleteTodoReq
	(*GetTodoReq)(nil),            // 3: api.user.v1.GetTodoReq
	(*TodoResp)(nil),              // 4: api.user.v1.TodoResp
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_todo_proto_depIdxs = []int32{
	5,  // 0: api.user.v1.CreateTodoReq.due_at:type_name -> google.protobuf.Timestamp
	5,  // 1: api.user.v1.CreateTodoReq.notify_at:type_name -> google.protobuf.Timestamp
	5,  // 2: api.user.v1.UpdateTodoReq.due_at:type_name -> google.protobuf.Timestamp
	5,  // 3: api.user.v1.UpdateTodoReq.notify_at:type_name -> google.protobuf.Timestamp
	5,  // 4: api.user.v1.TodoResp.due_at:type_name -> google.protobuf.Timestamp
	5,  // 5: api.user.v1.TodoResp.notify_at:type_name -> google.protobuf.Timestamp
	5,  // 6: api.user.v1.TodoResp.created_at:type_name -> google.protobuf.Timestamp
	5,  // 7: api.user.v1.TodoResp.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 8: api.user.v1.TodoSvc.Create:input_type -> api.user.v1.CreateTodoReq
	0,  // 9: api.user.v1.TodoSvc.Update:input_type -> api.user.v1.CreateTodoReq
	3,  // 10: api.user.v1.TodoSvc.Get:input_type -> api.user.v1.GetTodoReq
	2,  // 11: api.user.v1.TodoSvc.Delete:input_type -> api.user.v1.DeleteTodoReq
	4,  // 12: api.user.v1.TodoSvc.Create:output_type -> api.user.v1.TodoResp
	4,  // 13: api.user.v1.TodoSvc.Update:output_type -> api.user.v1.TodoResp
	4,  // 14: api.user.v1.TodoSvc.Get:output_type -> api.user.v1.TodoResp
	6,  // 15: api.user.v1.TodoSvc.Delete:output_type -> google.protobuf.Empty
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoReq); i {
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
		file_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoReq); i {
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
		file_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodoReq); i {
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
		file_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoReq); i {
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
		file_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoResp); i {
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
			RawDescriptor: file_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_rawDesc = nil
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}
