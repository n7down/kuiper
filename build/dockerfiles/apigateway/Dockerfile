FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/n7down/kuiper/
COPY . .
RUN go get -v -d ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main cmd/apigateway/main.go

FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates
COPY --from=builder /go/bin/main /go/bin/main
ENTRYPOINT ["/go/bin/main"] 
