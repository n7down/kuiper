FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/n7down/kuiper/
COPY . .
RUN go get -v -d ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main cmd/settings/main.go

FROM alpine:latest

ENV WAIT_VERSION 2.7.2

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait

RUN apk update && apk add --no-cache ca-certificates
RUN chmod +x /wait

COPY --from=builder /go/bin/main /go/bin/main
ENTRYPOINT ["/go/bin/main"] 
