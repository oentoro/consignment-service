build:
	protoc -I. --go_out=plugins=mciro:. \
		src/proto/consignment/consignment.proto
	go build
	docker build -t oentoro .
run:
	docker run -d -p 50051:50051 oentoro

clean:
	rm bin/oentoro