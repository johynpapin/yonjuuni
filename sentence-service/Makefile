build:
	protoc -I. --go_out=plugins=micro:${GOPATH}/src/github.com/johynpapin/yonjuuni/sentence-service \
		proto/sentence/sentence.proto
	docker build -t sentence-service .

run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns sentence-service