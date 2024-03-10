package main

import (
	"flag"
	"go-kit-example/string_service/middlewares"
	"go-kit-example/string_service/services"
	"go-kit-example/string_service/transports"
	"net/http"
	"os"

	"github.com/go-kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	var (
		listen = flag.String("listen", ":8081", "HTTP listen address")
	)
	flag.Parse()
	svc := services.StringServiceImpl{}
	logger := log.NewLogfmtLogger(os.Stderr)

	uppercaseEndpoint := transports.MakeUppercaseEndpoint(svc)
	uppercaseEndpoint = middlewares.LoggingMiddleware(log.With(logger, "method", "uppercase"))(uppercaseEndpoint)

	uppercaseHandler := httptransport.NewServer(
		uppercaseEndpoint,
		transports.DecodeUppercaseRequest,
		transports.EncodeResponse,
	)

	countEndpoint := transports.MakeCountEndpoint(svc)
	countEndpoint = middlewares.LoggingMiddleware(log.With(logger, "method", "count"))(countEndpoint)
	countHandler := httptransport.NewServer(
		countEndpoint,
		transports.DecodeCountRequest,
		transports.EncodeResponse,
	)
	logger.Log("msg", "HTTP", "addr", *listen)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.ListenAndServe(*listen, nil)
}
