module github.com/Haze-Lan/haze-go

go 1.16

require (
	go.etcd.io/etcd/client/v3 v3.5.0-alpha.0
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	google.golang.org/grpc v1.36.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.36.1
