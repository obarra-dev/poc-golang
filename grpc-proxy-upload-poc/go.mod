module grpc-proxy-upload-poc

go 1.17

replace service/myService => ./proto/gen/go

require (
	github.com/golang/glog v1.0.0
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	google.golang.org/grpc v1.40.0
	service/myService v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210903162649-d08c68adba83 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
