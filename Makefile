build:
	protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	#uncomment below and comment above if you are using linux
	# go build 
	docker build -t consignment-service .
run:
	docker run -d -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 consignment-service
clean:
	rm consignment-service