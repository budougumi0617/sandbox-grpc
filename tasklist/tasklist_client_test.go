// Copyright 2018 budougumi0617 All Rights Reserved.

// mock_tasklist has mock client code and test cases
package mock_tasklist_test

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"

	tlmock "github.com/budougumi0617/sandbox-grpc/tasklist/mockproto"
	tlpb "github.com/budougumi0617/sandbox-grpc/tasklist/proto"
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
		&rpcMsg{msg: req},
	).Return(task, nil)
	testGetTask(t, mockTaskManagerClient)
}

func testGetTask(t *testing.T, client tlpb.TaskManagerClient) {
	resp, err := client.GetTask(context.Background(), &tlpb.GetTaskRequest{Id: 1})
	if err != nil || resp.Title != task.Title {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", resp.Title)
}
