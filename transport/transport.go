package transport

import (
	"context"

	kit "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/nintran52/go-kit-grpc/endpoint"
	pb "github.com/nintran52/go-kit-grpc/proto"
)

// greeterServer implements pb.GreeterServer (from gRPC)
type greeterServer struct {
	sayHello grpc.Handler
	pb.UnimplementedGreeterServer
}

func NewGRPCServer(ep kit.Endpoint) pb.GreeterServer {
	return &greeterServer{
		sayHello: grpc.NewServer(ep, decodeGRPCSayHelloRequest, encodeGRPCSayHelloResponse),
	}
}

// SayHello implements the gRPC method
func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	_, resp, err := s.sayHello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.HelloReply), nil
}

// decodeGRPCSayHelloRequest decodes the gRPC request into a Go-Kit struct
func decodeGRPCSayHelloRequest(_ context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.HelloRequest)
	return endpoint.SayHelloRequest{Name: r.Name}, nil
}

// encodeGRPCSayHelloResponse encodes the Go-Kit response into a gRPC response
func encodeGRPCSayHelloResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r := resp.(endpoint.SayHelloResponse)
	return &pb.HelloReply{Message: r.Message}, nil
}
