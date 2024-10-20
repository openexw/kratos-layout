// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: helloworld/v1/greeter.proto

package helloworldv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/go-kratos/kratos-layout/gen/helloworld/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// GreeterServiceName is the fully-qualified name of the GreeterService service.
	GreeterServiceName = "helloworld.v1.GreeterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GreeterServiceSayHelloProcedure is the fully-qualified name of the GreeterService's SayHello RPC.
	GreeterServiceSayHelloProcedure = "/helloworld.v1.GreeterService/SayHello"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	greeterServiceServiceDescriptor        = v1.File_helloworld_v1_greeter_proto.Services().ByName("GreeterService")
	greeterServiceSayHelloMethodDescriptor = greeterServiceServiceDescriptor.Methods().ByName("SayHello")
)

// GreeterServiceClient is a client for the helloworld.v1.GreeterService service.
type GreeterServiceClient interface {
	// Sends a greeting
	// `other_shelf_name` to shelf `name`, and deletes
	// `other_shelf_name`. Returns the updated shelf.
	// The book ids of the moved books may not be the same as the original books.
	//
	// Returns NOT_FOUND if either shelf does not exist.
	// This call is a no-op if the specified shelves are the same.
	SayHello(context.Context, *connect.Request[v1.SayHelloReq]) (*connect.Response[v1.SayHelloResp], error)
}

// NewGreeterServiceClient constructs a client for the helloworld.v1.GreeterService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGreeterServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GreeterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &greeterServiceClient{
		sayHello: connect.NewClient[v1.SayHelloReq, v1.SayHelloResp](
			httpClient,
			baseURL+GreeterServiceSayHelloProcedure,
			connect.WithSchema(greeterServiceSayHelloMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// greeterServiceClient implements GreeterServiceClient.
type greeterServiceClient struct {
	sayHello *connect.Client[v1.SayHelloReq, v1.SayHelloResp]
}

// SayHello calls helloworld.v1.GreeterService.SayHello.
func (c *greeterServiceClient) SayHello(ctx context.Context, req *connect.Request[v1.SayHelloReq]) (*connect.Response[v1.SayHelloResp], error) {
	return c.sayHello.CallUnary(ctx, req)
}

// GreeterServiceHandler is an implementation of the helloworld.v1.GreeterService service.
type GreeterServiceHandler interface {
	// Sends a greeting
	// `other_shelf_name` to shelf `name`, and deletes
	// `other_shelf_name`. Returns the updated shelf.
	// The book ids of the moved books may not be the same as the original books.
	//
	// Returns NOT_FOUND if either shelf does not exist.
	// This call is a no-op if the specified shelves are the same.
	SayHello(context.Context, *connect.Request[v1.SayHelloReq]) (*connect.Response[v1.SayHelloResp], error)
}

// NewGreeterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGreeterServiceHandler(svc GreeterServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	greeterServiceSayHelloHandler := connect.NewUnaryHandler(
		GreeterServiceSayHelloProcedure,
		svc.SayHello,
		connect.WithSchema(greeterServiceSayHelloMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/helloworld.v1.GreeterService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GreeterServiceSayHelloProcedure:
			greeterServiceSayHelloHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGreeterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGreeterServiceHandler struct{}

func (UnimplementedGreeterServiceHandler) SayHello(context.Context, *connect.Request[v1.SayHelloReq]) (*connect.Response[v1.SayHelloResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("helloworld.v1.GreeterService.SayHello is not implemented"))
}
