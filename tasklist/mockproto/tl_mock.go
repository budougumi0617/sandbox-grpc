// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/budougumi0617/sandbox-grpc/tasklist/proto (interfaces: TaskManagerClient,TaskManager_ListTasksClient)

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	proto "github.com/budougumi0617/sandbox-grpc/tasklist/proto"
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockTaskManagerClient is a mock of TaskManagerClient interface
type MockTaskManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockTaskManagerClientMockRecorder
}

// MockTaskManagerClientMockRecorder is the mock recorder for MockTaskManagerClient
type MockTaskManagerClientMockRecorder struct {
	mock *MockTaskManagerClient
}

// NewMockTaskManagerClient creates a new mock instance
func NewMockTaskManagerClient(ctrl *gomock.Controller) *MockTaskManagerClient {
	mock := &MockTaskManagerClient{ctrl: ctrl}
	mock.recorder = &MockTaskManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskManagerClient) EXPECT() *MockTaskManagerClientMockRecorder {
	return m.recorder
}

// GetTask mocks base method
func (m *MockTaskManagerClient) GetTask(arg0 context.Context, arg1 *proto.GetTaskRequest, arg2 ...grpc.CallOption) (*proto.Task, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTask", varargs...)
	ret0, _ := ret[0].(*proto.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask
func (mr *MockTaskManagerClientMockRecorder) GetTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockTaskManagerClient)(nil).GetTask), varargs...)
}

// ListTasks mocks base method
func (m *MockTaskManagerClient) ListTasks(arg0 context.Context, arg1 *empty.Empty, arg2 ...grpc.CallOption) (proto.TaskManager_ListTasksClient, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTasks", varargs...)
	ret0, _ := ret[0].(proto.TaskManager_ListTasksClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks
func (mr *MockTaskManagerClientMockRecorder) ListTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockTaskManagerClient)(nil).ListTasks), varargs...)
}

// MockTaskManager_ListTasksClient is a mock of TaskManager_ListTasksClient interface
type MockTaskManager_ListTasksClient struct {
	ctrl     *gomock.Controller
	recorder *MockTaskManager_ListTasksClientMockRecorder
}

// MockTaskManager_ListTasksClientMockRecorder is the mock recorder for MockTaskManager_ListTasksClient
type MockTaskManager_ListTasksClientMockRecorder struct {
	mock *MockTaskManager_ListTasksClient
}

// NewMockTaskManager_ListTasksClient creates a new mock instance
func NewMockTaskManager_ListTasksClient(ctrl *gomock.Controller) *MockTaskManager_ListTasksClient {
	mock := &MockTaskManager_ListTasksClient{ctrl: ctrl}
	mock.recorder = &MockTaskManager_ListTasksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskManager_ListTasksClient) EXPECT() *MockTaskManager_ListTasksClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockTaskManager_ListTasksClient) CloseSend() error {
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockTaskManager_ListTasksClientMockRecorder) CloseSend() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockTaskManager_ListTasksClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockTaskManager_ListTasksClient) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockTaskManager_ListTasksClientMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTaskManager_ListTasksClient)(nil).Context))
}

// Header mocks base method
func (m *MockTaskManager_ListTasksClient) Header() (metadata.MD, error) {
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockTaskManager_ListTasksClientMockRecorder) Header() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockTaskManager_ListTasksClient)(nil).Header))
}

// Recv mocks base method
func (m *MockTaskManager_ListTasksClient) Recv() (*proto.Task, error) {
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*proto.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockTaskManager_ListTasksClientMockRecorder) Recv() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockTaskManager_ListTasksClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockTaskManager_ListTasksClient) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockTaskManager_ListTasksClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTaskManager_ListTasksClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockTaskManager_ListTasksClient) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockTaskManager_ListTasksClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTaskManager_ListTasksClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockTaskManager_ListTasksClient) Trailer() metadata.MD {
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockTaskManager_ListTasksClientMockRecorder) Trailer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockTaskManager_ListTasksClient)(nil).Trailer))
}