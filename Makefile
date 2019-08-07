build:
	protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto
	go build
	docker build -t consignment-service .
	rm consignment-service
run:
	docker run -d -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 consignment-service
clean:
	rm consignment-service