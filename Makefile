build:
	protoc -I. --go_out=. --micro_out=. \
		proto/vessel/vessel.proto
	docker build -t vessel-service .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 vessel-service