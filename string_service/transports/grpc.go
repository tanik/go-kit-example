package transport

import (
	"context"

	"github.com/go-kit/log"

	"go-kit-example/string_service/endpoints"
	pb "go-kit-example/string_service/gen/go/proto/v1"

	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	uppercase gt.Handler
	count     gt.Handler
	pb.UnimplementedStringServiceServer
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.StringServiceServer {
	uppercaseHandler := gt.NewServer(
		endpoints.Uppercase,
		decodeUppercaseRequest,
		encodeUppercaseResponse,
	)
	countHandler := gt.NewServer(
		endpoints.Count,
		decodeCountRequest,
		encodeCountResponse,
	)
	return &gRPCServer{uppercase: uppercaseHandler, count: countHandler}
}

func (s *gRPCServer) Uppercase(ctx context.Context, req *pb.UppercaseRequest) (*pb.UppercaseResponse, error) {
	_, resp, err := s.uppercase.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UppercaseResponse), nil
}

func (s *gRPCServer) Count(ctx context.Context, req *pb.CountRequest) (*pb.CountResponse, error) {
	_, resp, err := s.count.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CountResponse), nil
}

func decodeUppercaseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UppercaseRequest)
	return endpoints.UppercaseReq{Str: req.Str}, nil
}

func encodeUppercaseResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.UppercaseRes)
	return &pb.UppercaseResponse{Val: resp.Value, Err: resp.Error.Error()}, nil
}

func decodeCountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CountRequest)
	return endpoints.CountReq{Str: req.Str}, nil
}

func encodeCountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.CountRes)
	return &pb.CountResponse{Val: resp.Value}, nil
}
