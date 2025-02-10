package service

import (
	"context"
)

// GreeterService defines the service with Go-Kit
type GreeterService interface {
	SayHello(ctx context.Context, name string) (string, error)
}

// greeterService implements GreeterService
type greeterService struct{}

func NewGreeterService() GreeterService {
	return &greeterService{}
}

func (s *greeterService) SayHello(ctx context.Context, name string) (string, error) {
	return "Hello, " + name, nil
}
