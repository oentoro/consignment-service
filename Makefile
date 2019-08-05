build:
	protoc -I. --go_out=plugins=grpc:/home/vagrant/code/consignment-service/proto/consignment/consignment.proto
	go build
	docker build -t oentoro .
run:
	docker run -d -p 50051:50051 oentoro \
	-e MICRO_SERVER_ADDRESS=:50051 \
	consignment-service

clean:
	rm bin/oentoro