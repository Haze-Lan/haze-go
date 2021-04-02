module github.com/Haze-Lan/haze-go

go 1.16

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/nats-io/nats-server/v2 v2.2.0
	github.com/rpcxio/libkv v0.5.0
	github.com/rpcxio/rpcx-etcd v0.0.0-20210309015740-324bede3ab7c
	go.etcd.io/etcd/client/v3 v3.5.0-alpha.0
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20210331212208-0fccb6fa2b5c // indirect
	golang.org/x/sys v0.0.0-20210331175145-43e1dd70ce54
	google.golang.org/genproto v0.0.0-20210401141331-865547bb08e2 // indirect
	google.golang.org/grpc v1.36.1
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.36.1
