package middlewares

import (
	"context"
	"errors"
	"fmt"
	services "go-kit-example/api_server/services"
	transports "go-kit-example/api_server/transports"

	pb "go-kit-example/string_service/gen/go/proto"

	"strings"
	"time"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sony/gobreaker"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

func ProxyMiddleware(ctx context.Context, instances string, logger log.Logger) services.ServiceMiddleware {
	// If instances is empty, don't proxy.
	if instances == "" {
		logger.Log("proxy_to", "none")
		return func(next services.StringService) services.StringService { return next }
	}

	var (
		qps         = 100                    // beyond which we will return an error
		maxAttempts = 3                      // per request, before giving up
		maxTime     = 250 * time.Millisecond // wallclock time, before giving up
	)

	var (
		instanceList = split(instances)
		endpointer   sd.FixedEndpointer
	)
	logger.Log("proxy_to", fmt.Sprint(instanceList))
	for _, instance := range instanceList {
		var e endpoint.Endpoint
		e = makeUppercaseProxy(ctx, instance)
		e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		e = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), qps))(e)
		endpointer = append(endpointer, e)
	}

	// Now, build a single, retrying, load-balancing endpoint out of all of
	// those individual endpoints.
	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(maxAttempts, maxTime, balancer)

	// And finally, return the ServiceMiddleware, implemented by proxymw.
	return func(next services.StringService) services.StringService {
		return proxymw{ctx, next, retry}
	}
}

// proxymw implements StringService, forwarding Uppercase requests to the
// provided endpoint, and serving all other (i.e. Count) requests via the
// next StringService.
type proxymw struct {
	ctx       context.Context
	next      services.StringService // Serve most requests via this service...
	uppercase endpoint.Endpoint      // ...except Uppercase, which gets served by this endpoint
}

func (mw proxymw) Count(s string) int {
	return mw.next.Count(s)
}

func (mw proxymw) Uppercase(s string) (string, error) {
	response, err := mw.uppercase(mw.ctx, transports.UppercaseRequest{S: s})
	if err != nil {
		return "", err
	}

	resp := response.(transports.UppercaseResponse)
	if resp.Err != "" {
		return resp.V, errors.New(resp.Err)
	}
	return resp.V, nil
}

func makeUppercaseProxy(_ context.Context, instance string) endpoint.Endpoint {
	conn, err := grpc.Dial(instance, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return grpctransport.NewClient(
		conn,
		"proto.StringService",
		"Uppercase",
		encodeUppercaseRequest,
		decodeUppercaseResponse,
		pb.UppercaseResponse{},
	).Endpoint()
}

func split(s string) []string {
	a := strings.Split(s, ",")
	for i := range a {
		a[i] = strings.TrimSpace(a[i])
	}
	return a
}

func encodeUppercaseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(transports.UppercaseRequest)
	return &pb.UppercaseRequest{Str: req.S}, nil
}

func decodeUppercaseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.UppercaseResponse)
	return transports.UppercaseResponse{V: reply.Val, Err: reply.Err}, nil
}
