run-tel-server:
	go run cmd/cal_telnet/server/main.go

run-grpc-server:
	go run cmd/cal_grpc/server/main.go

run-grpc-client:
	go run cmd/cal_grpc/client/main.go

gen: 
	protoc --go_out=../ --go-grpc_out=../ serviceapi/proto/service.proto	

clean:
	rm server/internal/gen/proto/*.go	