FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/n7down/iota/
COPY . .
RUN go get -v -d ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main cmd/iota/main.go

FROM alpine:latest

#ARG TZ='MST'
#ENV TZ ${TZ}
RUN apk update && apk add --no-cache ca-certificates
#RUN apk add --no-cache tzdata
#RUN cp /usr/share/zoneinfo/${TZ} /etc/localtime \
    #     && echo "${TZ}" > /etc/timezone \
    #     && apk del tzdata \
    #     && rm -rf /var/cache/apk/*

#RUN date

COPY --from=builder /go/bin/main /go/bin/main
ENTRYPOINT ["/go/bin/main"] 
