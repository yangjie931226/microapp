module grades

go 1.15

require (
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.7.0
	github.com/asim/go-micro/plugins/transport/grpc/v3 v3.7.0
	github.com/asim/go-micro/v3 v3.7.0
	github.com/hashicorp/go-hclog v0.14.1 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98 // indirect
	google.golang.org/grpc v1.40.0 // indirect
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
