module github.com/shippy/shippy-cli-consignment

go 1.16

replace github.com/jasonsalas/shippy/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/jasonsalas/shippy/shippy-service-consignment v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/genproto v0.0.0-20210729151513-df9385d47c1b // indirect
	google.golang.org/grpc v1.39.0
)
