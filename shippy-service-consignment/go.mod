module github.com/jasonsalas/shippy/shippy-service-consignment

go 1.16

replace github.com/jasonsalas/shippy/shippy-service-consignment => ../shippy-service-consignment

require (
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)
