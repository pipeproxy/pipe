FROM golang:alpine AS builder
WORKDIR /go/src/github.com/wzshiming/pipe/
COPY . .
RUN go install ./cmd/pipe

FROM alpine
COPY --from=builder /go/bin/pipe /usr/local/bin/
ENTRYPOINT [ "pipe" ]
