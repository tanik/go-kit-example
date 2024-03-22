package endpoints

import (
	"context"
	"go-kit-example/string_service/services"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
	Uppercase endpoint.Endpoint
	Count     endpoint.Endpoint
}

type UppercaseReq struct {
	Str string
}

type UppercaseRes struct {
	Value string
	Error error
}

type CountReq struct {
	Str string
}

type CountRes struct {
	Value uint64
}

func MakeEndpoints(s services.StringService) Endpoints {
	return Endpoints{
		Uppercase: makeUppercaseEndPoint(s),
		Count:     makeCountEndpoint(s),
	}
}

func makeUppercaseEndPoint(s services.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UppercaseReq)
		value, err := s.Uppercase(ctx, req.Str)
		return UppercaseRes{Value: value, Error: err}, nil
	}
}

func makeCountEndpoint(s services.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CountReq)
		value := s.Count(ctx, req.Str)
		return CountRes{Value: value}, nil
	}
}
