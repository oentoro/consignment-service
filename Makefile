build:
	protoc -I. --go_out=plugins=grpc:. proto/consignment/consignment.proto
	go build
	docker build -t consignment-service .
run:
	docker run -d -p 50051:50051 consignment-service
clean:
	rm consignment-service