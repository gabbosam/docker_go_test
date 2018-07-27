FROM golang:1.10 as builder
WORKDIR /go/src/teamsystem.com/gciappelloni/gotest/
RUN go get -d -v golang.org/x/net/html
COPY hello.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /root/
RUN apk add curl
RUN mkdir /etc/server
COPY --from=builder /go/src/teamsystem.com/gciappelloni/gotest/app .
HEALTHCHECK --interval=5m --timeout=3s CMD ["curl", "-f", "http://localhost/health"]
CMD ["./app"]
