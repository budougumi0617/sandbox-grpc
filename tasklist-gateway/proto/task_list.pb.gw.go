// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/task_list.proto

/*
Package tasklistgateway is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package tasklistgateway

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_TaskManager_GetTask_0(ctx context.Context, marshaler runtime.Marshaler, client TaskManagerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetTaskRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.Int32(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.GetTask(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_TaskManager_ListTasks_0(ctx context.Context, marshaler runtime.Marshaler, client TaskManagerClient, req *http.Request, pathParams map[string]string) (TaskManager_ListTasksClient, runtime.ServerMetadata, error) {
	var protoReq empty.Empty
	var metadata runtime.ServerMetadata

	stream, err := client.ListTasks(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterTaskManagerHandlerFromEndpoint is same as RegisterTaskManagerHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterTaskManagerHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterTaskManagerHandler(ctx, mux, conn)
}

// RegisterTaskManagerHandler registers the http handlers for service TaskManager to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterTaskManagerHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterTaskManagerHandlerClient(ctx, mux, NewTaskManagerClient(conn))
}

// RegisterTaskManagerHandler registers the http handlers for service TaskManager to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "TaskManagerClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "TaskManagerClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "TaskManagerClient" to call the correct interceptors.
func RegisterTaskManagerHandlerClient(ctx context.Context, mux *runtime.ServeMux, client TaskManagerClient) error {

	mux.Handle("GET", pattern_TaskManager_GetTask_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TaskManager_GetTask_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TaskManager_GetTask_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_TaskManager_ListTasks_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TaskManager_ListTasks_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TaskManager_ListTasks_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_TaskManager_GetTask_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "tasklist-gateway", "task", "id"}, ""))

	pattern_TaskManager_ListTasks_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "tasklist-gateway", "task"}, ""))
)

var (
	forward_TaskManager_GetTask_0 = runtime.ForwardResponseMessage

	forward_TaskManager_ListTasks_0 = runtime.ForwardResponseStream
)
