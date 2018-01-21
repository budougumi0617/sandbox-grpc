// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/task_list.proto

/*
Package tasklist is a generated protocol buffer package.

It is generated from these files:
	proto/task_list.proto

It has these top-level messages:
	Task
	GetTaskRequest
*/
package tasklist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Task struct {
	Id     int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Detail string `protobuf:"bytes,3,opt,name=detail" json:"detail,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Task) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Task) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type GetTaskRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetTaskRequest) Reset()                    { *m = GetTaskRequest{} }
func (m *GetTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTaskRequest) ProtoMessage()               {}
func (*GetTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetTaskRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Task)(nil), "tasklist.Task")
	proto.RegisterType((*GetTaskRequest)(nil), "tasklist.GetTaskRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskManager service

type TaskManagerClient interface {
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*Task, error)
	ListTasks(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (TaskManager_ListTasksClient, error)
}

type taskManagerClient struct {
	cc *grpc.ClientConn
}

func NewTaskManagerClient(cc *grpc.ClientConn) TaskManagerClient {
	return &taskManagerClient{cc}
}

func (c *taskManagerClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/tasklist.TaskManager/GetTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) ListTasks(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (TaskManager_ListTasksClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskManager_serviceDesc.Streams[0], c.cc, "/tasklist.TaskManager/ListTasks", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskManagerListTasksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskManager_ListTasksClient interface {
	Recv() (*Task, error)
	grpc.ClientStream
}

type taskManagerListTasksClient struct {
	grpc.ClientStream
}

func (x *taskManagerListTasksClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TaskManager service

type TaskManagerServer interface {
	GetTask(context.Context, *GetTaskRequest) (*Task, error)
	ListTasks(*google_protobuf.Empty, TaskManager_ListTasksServer) error
}

func RegisterTaskManagerServer(s *grpc.Server, srv TaskManagerServer) {
	s.RegisterService(&_TaskManager_serviceDesc, srv)
}

func _TaskManager_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasklist.TaskManager/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_ListTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(google_protobuf.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskManagerServer).ListTasks(m, &taskManagerListTasksServer{stream})
}

type TaskManager_ListTasksServer interface {
	Send(*Task) error
	grpc.ServerStream
}

type taskManagerListTasksServer struct {
	grpc.ServerStream
}

func (x *taskManagerListTasksServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

var _TaskManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tasklist.TaskManager",
	HandlerType: (*TaskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTask",
			Handler:    _TaskManager_GetTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTasks",
			Handler:       _TaskManager_ListTasks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/task_list.proto",
}

func init() { proto.RegisterFile("proto/task_list.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x49, 0x2c, 0xce, 0x8e, 0xcf, 0xc9, 0x2c, 0x2e, 0xd1, 0x03, 0xf3, 0x85, 0x38,
	0x40, 0x02, 0x20, 0xbe, 0x94, 0x74, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x3c, 0xa9,
	0x34, 0x4d, 0x3f, 0x35, 0xb7, 0xa0, 0xa4, 0x12, 0xa2, 0x4c, 0xc9, 0x85, 0x8b, 0x25, 0x24, 0xb1,
	0x38, 0x5b, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x88, 0x29,
	0x33, 0x45, 0x48, 0x84, 0x8b, 0xb5, 0x24, 0xb3, 0x24, 0x27, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xc2, 0x11, 0x12, 0xe3, 0x62, 0x4b, 0x49, 0x2d, 0x49, 0xcc, 0xcc, 0x91, 0x60, 0x06,
	0x0b, 0x43, 0x79, 0x4a, 0x0a, 0x5c, 0x7c, 0xee, 0xa9, 0x25, 0x20, 0x83, 0x82, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0xd0, 0xcd, 0x33, 0xaa, 0xe5, 0xe2, 0x06, 0x49, 0xfb, 0x26, 0xe6, 0x25, 0xa6,
	0xa7, 0x16, 0x09, 0x99, 0x72, 0xb1, 0x43, 0x35, 0x08, 0x49, 0xe8, 0xc1, 0x5c, 0xaa, 0x87, 0x6a,
	0x86, 0x14, 0x1f, 0x42, 0x06, 0x24, 0xac, 0xc4, 0x20, 0x64, 0xce, 0xc5, 0xe9, 0x93, 0x59, 0x0c,
	0x56, 0x54, 0x2c, 0x24, 0xa6, 0x07, 0xf1, 0x98, 0x1e, 0xcc, 0x63, 0x7a, 0xae, 0x20, 0x8f, 0x61,
	0x6a, 0x33, 0x60, 0x4c, 0x62, 0x03, 0xab, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x45, 0xcf,
	0xbb, 0x21, 0x2d, 0x01, 0x00, 0x00,
}
