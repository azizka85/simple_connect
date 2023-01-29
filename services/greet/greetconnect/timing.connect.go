// Code generated by protoc-gen-connect-go.exe. DO NOT EDIT.
//
// Source: proto/greet/timing.proto

package greetconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	greet "simple_connect/services/greet"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// TimingInputServiceName is the fully-qualified name of the TimingInputService service.
	TimingInputServiceName = "greet.TimingInputService"
	// TimingOutputServiceName is the fully-qualified name of the TimingOutputService service.
	TimingOutputServiceName = "greet.TimingOutputService"
)

// TimingInputServiceClient is a client for the greet.TimingInputService service.
type TimingInputServiceClient interface {
	Send(context.Context) *connect_go.ClientStreamForClient[greet.TimingInputRequest, greet.TimingInputResponse]
}

// NewTimingInputServiceClient constructs a client for the greet.TimingInputService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTimingInputServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TimingInputServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &timingInputServiceClient{
		send: connect_go.NewClient[greet.TimingInputRequest, greet.TimingInputResponse](
			httpClient,
			baseURL+"/greet.TimingInputService/Send",
			opts...,
		),
	}
}

// timingInputServiceClient implements TimingInputServiceClient.
type timingInputServiceClient struct {
	send *connect_go.Client[greet.TimingInputRequest, greet.TimingInputResponse]
}

// Send calls greet.TimingInputService.Send.
func (c *timingInputServiceClient) Send(ctx context.Context) *connect_go.ClientStreamForClient[greet.TimingInputRequest, greet.TimingInputResponse] {
	return c.send.CallClientStream(ctx)
}

// TimingInputServiceHandler is an implementation of the greet.TimingInputService service.
type TimingInputServiceHandler interface {
	Send(context.Context, *connect_go.ClientStream[greet.TimingInputRequest]) (*connect_go.Response[greet.TimingInputResponse], error)
}

// NewTimingInputServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTimingInputServiceHandler(svc TimingInputServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/greet.TimingInputService/Send", connect_go.NewClientStreamHandler(
		"/greet.TimingInputService/Send",
		svc.Send,
		opts...,
	))
	return "/greet.TimingInputService/", mux
}

// UnimplementedTimingInputServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTimingInputServiceHandler struct{}

func (UnimplementedTimingInputServiceHandler) Send(context.Context, *connect_go.ClientStream[greet.TimingInputRequest]) (*connect_go.Response[greet.TimingInputResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("greet.TimingInputService.Send is not implemented"))
}

// TimingOutputServiceClient is a client for the greet.TimingOutputService service.
type TimingOutputServiceClient interface {
	Accept(context.Context, *connect_go.Request[greet.TimingInputRequest]) (*connect_go.ServerStreamForClient[greet.MessageInfo], error)
}

// NewTimingOutputServiceClient constructs a client for the greet.TimingOutputService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTimingOutputServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TimingOutputServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &timingOutputServiceClient{
		accept: connect_go.NewClient[greet.TimingInputRequest, greet.MessageInfo](
			httpClient,
			baseURL+"/greet.TimingOutputService/Accept",
			opts...,
		),
	}
}

// timingOutputServiceClient implements TimingOutputServiceClient.
type timingOutputServiceClient struct {
	accept *connect_go.Client[greet.TimingInputRequest, greet.MessageInfo]
}

// Accept calls greet.TimingOutputService.Accept.
func (c *timingOutputServiceClient) Accept(ctx context.Context, req *connect_go.Request[greet.TimingInputRequest]) (*connect_go.ServerStreamForClient[greet.MessageInfo], error) {
	return c.accept.CallServerStream(ctx, req)
}

// TimingOutputServiceHandler is an implementation of the greet.TimingOutputService service.
type TimingOutputServiceHandler interface {
	Accept(context.Context, *connect_go.Request[greet.TimingInputRequest], *connect_go.ServerStream[greet.MessageInfo]) error
}

// NewTimingOutputServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTimingOutputServiceHandler(svc TimingOutputServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/greet.TimingOutputService/Accept", connect_go.NewServerStreamHandler(
		"/greet.TimingOutputService/Accept",
		svc.Accept,
		opts...,
	))
	return "/greet.TimingOutputService/", mux
}

// UnimplementedTimingOutputServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTimingOutputServiceHandler struct{}

func (UnimplementedTimingOutputServiceHandler) Accept(context.Context, *connect_go.Request[greet.TimingInputRequest], *connect_go.ServerStream[greet.MessageInfo]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("greet.TimingOutputService.Accept is not implemented"))
}