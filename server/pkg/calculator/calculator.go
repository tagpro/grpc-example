package calculator

import (
	v1 "github.com/tagpro/grpc-example/pkg/tagpro/services/calculator"
)

type calculatorService struct {
	v1.UnimplementedCalculatorServer
}

func NewCalculatorService() v1.CalculatorServer {
	return &calculatorService{}
}
