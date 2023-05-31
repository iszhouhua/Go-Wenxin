// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: baidubce/ai_custom/v1/wenxinworkshop_service.proto

/*
Package ai_customv1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package ai_customv1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_WenxinworkshopService_ChatCompletions_0(ctx context.Context, marshaler runtime.Marshaler, client WenxinworkshopServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ChatCompletionsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ChatCompletions(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_WenxinworkshopService_ChatCompletions_0(ctx context.Context, marshaler runtime.Marshaler, server WenxinworkshopServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ChatCompletionsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ChatCompletions(ctx, &protoReq)
	return msg, metadata, err

}

func request_WenxinworkshopService_ChatEbInstant_0(ctx context.Context, marshaler runtime.Marshaler, client WenxinworkshopServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ChatEbInstantRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ChatEbInstant(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_WenxinworkshopService_ChatEbInstant_0(ctx context.Context, marshaler runtime.Marshaler, server WenxinworkshopServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ChatEbInstantRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ChatEbInstant(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterWenxinworkshopServiceHandlerServer registers the http handlers for service WenxinworkshopService to "mux".
// UnaryRPC     :call WenxinworkshopServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterWenxinworkshopServiceHandlerFromEndpoint instead.
func RegisterWenxinworkshopServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server WenxinworkshopServiceServer) error {

	mux.Handle("POST", pattern_WenxinworkshopService_ChatCompletions_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/baidubce.ai_custom.v1.WenxinworkshopService/ChatCompletions", runtime.WithHTTPPathPattern("/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_WenxinworkshopService_ChatCompletions_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WenxinworkshopService_ChatCompletions_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_WenxinworkshopService_ChatEbInstant_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/baidubce.ai_custom.v1.WenxinworkshopService/ChatEbInstant", runtime.WithHTTPPathPattern("/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_WenxinworkshopService_ChatEbInstant_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WenxinworkshopService_ChatEbInstant_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterWenxinworkshopServiceHandlerFromEndpoint is same as RegisterWenxinworkshopServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterWenxinworkshopServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterWenxinworkshopServiceHandler(ctx, mux, conn)
}

// RegisterWenxinworkshopServiceHandler registers the http handlers for service WenxinworkshopService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterWenxinworkshopServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterWenxinworkshopServiceHandlerClient(ctx, mux, NewWenxinworkshopServiceClient(conn))
}

// RegisterWenxinworkshopServiceHandlerClient registers the http handlers for service WenxinworkshopService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "WenxinworkshopServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "WenxinworkshopServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "WenxinworkshopServiceClient" to call the correct interceptors.
func RegisterWenxinworkshopServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client WenxinworkshopServiceClient) error {

	mux.Handle("POST", pattern_WenxinworkshopService_ChatCompletions_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/baidubce.ai_custom.v1.WenxinworkshopService/ChatCompletions", runtime.WithHTTPPathPattern("/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_WenxinworkshopService_ChatCompletions_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WenxinworkshopService_ChatCompletions_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_WenxinworkshopService_ChatEbInstant_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/baidubce.ai_custom.v1.WenxinworkshopService/ChatEbInstant", runtime.WithHTTPPathPattern("/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_WenxinworkshopService_ChatEbInstant_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WenxinworkshopService_ChatEbInstant_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_WenxinworkshopService_ChatCompletions_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 2, 5, 2, 6}, []string{"rpc", "2.0", "ai_custom", "v1", "wenxinworkshop", "chat", "completions"}, ""))

	pattern_WenxinworkshopService_ChatEbInstant_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 2, 5, 2, 6}, []string{"rpc", "2.0", "ai_custom", "v1", "wenxinworkshop", "chat", "eb-instant"}, ""))
)

var (
	forward_WenxinworkshopService_ChatCompletions_0 = runtime.ForwardResponseMessage

	forward_WenxinworkshopService_ChatEbInstant_0 = runtime.ForwardResponseMessage
)
