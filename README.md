# Microservices in Go with gRPC

## This repo builds upon the [outstanding tutorial series](https://ewanvalentine.io/microservices-in-golang-part-0/) on developing a comprehensive, end-to-end microservice-driven web application by [@Ewan_Valentine](https://twitter.com/Ewan_Valentine)

***

### This demo uses the following technologies in each tagged release:

- Go
- Kubernetes
- gRPC
- Micro (Go microservice framework)
- MongoDB and/or Postgres
- Docker
- Google Cloud

***

### Modifications to protobuf definitions

This forked repo makes a couple of additions to [the original project](https://ewanvalentine.io/microservices-in-golang-part-1/) with the compilation instructions to generate the [protocol buffer definitions](https://github.com/jasonsalas/shippy/blog/main/shippy-service-consignment/proto/consignment/consignment.pb.go), and the inclusion of the `go_package` option in the [protobuf service definition](https://github.com/jasonsalas/shippy/blog/main/shippy-service-consignment/proto/consignment/consignment.proto):

- Command to run at terminal in the `./shippy/shippy-service-consignment` directory: `protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative proto/consignment/consignment.proto`
- In [consignment.proto](https://github.com/jasonsalas/shippy/blog/main/shippy-service-consignment/proto/consignment/consignment.proto): `option go_package = "github.com/jasonsalas/shippy/shippy-service-consignment";`