module go-kit-example/api_server

go 1.21.6

require (
	github.com/go-kit/kit v0.13.0
	github.com/go-kit/log v0.2.1
	github.com/prometheus/client_golang v1.11.1
	github.com/sony/gobreaker v0.4.1
	go-kit-example/string_service v0.0.0-00010101000000-000000000000
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11
	google.golang.org/grpc v1.57.0
)

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.30.0 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/smartystreets/goconvey v1.8.1 // indirect
	github.com/streadway/handy v0.0.0-20200128134331-0f66f006fb2e // indirect
	github.com/stretchr/testify v1.7.2 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace go-kit-example/string_service => ../string_service
