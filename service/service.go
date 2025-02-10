package service

import (
	"context"
)

// GreeterService định nghĩa service với Go-Kit
type GreeterService interface {
	SayHello(ctx context.Context, name string) (string, error)
}

// greeterService triển khai GreeterService
type greeterService struct{}

func NewGreeterService() GreeterService {
	return &greeterService{}
}

func (s *greeterService) SayHello(ctx context.Context, name string) (string, error) {
	return "Hello, " + name, nil
}
