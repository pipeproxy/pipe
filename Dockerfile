FROM golang:alpine AS builder
WORKDIR /go/src/github.com/pipeproxy/pipe/
COPY . .
RUN go install ./cmd/pipe

FROM alpine
COPY --from=builder /go/bin/pipe /usr/local/bin/
COPY ./pipe.yml /
ENTRYPOINT [ "pipe" ]
