package main

import (
	"flag"
	"fmt"
	endpoint "go-kit-example/string_service/endpoints"
	pb "go-kit-example/string_service/gen/go/proto/v1"
	"go-kit-example/string_service/services"
	transport "go-kit-example/string_service/transports"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc/reflection"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"google.golang.org/grpc"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	var (
		listen = flag.String("listen", ":8081", "HTTP listen address")
	)
	flag.Parse()
	logger := log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	addService := services.NewService(logger)
	addEndpoint := endpoint.MakeEndpoints(addService)
	grpcServer := transport.NewGRPCServer(addEndpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", *listen)
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		server := grpc.NewServer()
		pb.RegisterStringServiceServer(server, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		reflection.Register(server)
		server.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)

}
