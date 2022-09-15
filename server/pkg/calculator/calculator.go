package calculator

import (
	"context"

	v1 "github.com/tagpro/grpc-example/pkg/tagpro/services/calculator"
)

type calculatorService struct {
	v1.UnimplementedCalculatorServer
}

func NewCalculatorService() v1.CalculatorServer {
	return &calculatorService{}
}

func (s *calculatorService) Add(ctx context.Context, req *v1.AddRequest) (*v1.AddResponse, error) {
	return &v1.AddResponse{
		Result: req.GetA() + req.GetB(),
	}, nil
}
