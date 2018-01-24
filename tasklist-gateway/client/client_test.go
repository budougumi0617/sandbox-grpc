// Copyright 2018 budougumi0617 All Rights Reserved.

// client has mock client code and test cases
package client

import (
	"fmt"
	"io"
	"testing"

	"golang.org/x/net/context"

	tlmock "github.com/budougumi0617/sandbox-grpc/tasklistgateway/mockproto"
	tlpb "github.com/budougumi0617/sandbox-grpc/tasklistgateway/proto"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
)

var task = &tlpb.Task{
	Id:     1,
	Title:  "dummy title",
	Detail: "dummy detail",
}

type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func TestGetTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTaskManagerClient := tlmock.NewMockTaskManagerClient(ctrl)
	req := &tlpb.GetTaskRequest{Id: 1}
	mockTaskManagerClient.EXPECT().GetTask(
		gomock.Any(),
		req,
	).Return(task, nil)
	testGetTask(t, mockTaskManagerClient)
}

func testGetTask(t *testing.T, client tlpb.TaskManagerClient) {
	t.Helper()
	resp, err := client.GetTask(context.Background(), &tlpb.GetTaskRequest{Id: 1})
	if err != nil || resp.Title != task.Title {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", resp.Title)
}

func TestListTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock for the stream returned by ListTasks
	stream := tlmock.NewMockTaskManager_ListTasksClient(ctrl)
	stream.EXPECT().Recv().Return(&tlpb.Task{
		Id:     1,
		Title:  "first Task",
		Detail: "fist Detail",
	}, nil)
	stream.EXPECT().Recv().Return(&tlpb.Task{
		Id:     2,
		Title:  "second Task",
		Detail: "second Detail",
	}, nil)
	stream.EXPECT().Recv().Return(nil, io.EOF)
	// Create mock for the client interface.
	mockclient := tlmock.NewMockTaskManagerClient(ctrl)
	mockclient.EXPECT().ListTasks(
		gomock.Any(),
		gomock.Any(),
	).Return(stream, nil)
	testListTasks(t, mockclient)
}

func testListTasks(t *testing.T, client tlpb.TaskManagerClient) {
	t.Helper()
	ltc, _ := client.ListTasks(context.Background(), nil)
	first, err := ltc.Recv()
	if err != nil || first.Title != "first Task" {
		t.Errorf("Unexpected task at first response")
	}
	second, err := ltc.Recv()
	if err != nil || second.Title != "second Task" {
		t.Errorf("Unexpected task at second response")
	}
	_, eof := ltc.Recv()
	if eof != io.EOF {
		t.Error(eof)
	}
}
