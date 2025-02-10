package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nintran52/go-kit-grpc/service"
)

// SayHelloRequest và SayHelloResponse đại diện cho request/response
type SayHelloRequest struct {
	Name string
}

type SayHelloResponse struct {
	Message string
}

// MakeSayHelloEndpoint tạo Go-Kit endpoint
func MakeSayHelloEndpoint(s service.GreeterService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SayHelloRequest)
		msg, err := s.SayHello(ctx, req.Name)
		if err != nil {
			return nil, err
		}
		return SayHelloResponse{Message: msg}, nil
	}
}
