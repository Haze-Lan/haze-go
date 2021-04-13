module github.com/Haze-Lan/haze-go

go 1.16

require (
	github.com/go-errors/errors v1.0.1
	github.com/magiconair/properties v1.8.5
	go.etcd.io/etcd/client/v3 v3.5.0-alpha.0
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20210410081132-afb366fc7cd1 // indirect
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57
	google.golang.org/genproto v0.0.0-20210406143921-e86de6bf7a46 // indirect
	google.golang.org/grpc v1.37.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.37.0
