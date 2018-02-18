build:
	protoc -I. --go_out=plugins=micro:${GOPATH}/src/github.com/johynpapin/yonjuuni/email-service \
		proto/email/email.proto
	docker build -t email-service .

run:
	docker run -p 50052:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns email-service