FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/johynpapin/yonjuuni/email-service

COPY . .

#RUN curl https://glide.sh/get | sh
#RUN glide install
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/johynpapin/yonjuuni/email-service/email-service .

CMD ["./email-service"]