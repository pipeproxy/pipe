FROM golang:alpine AS builder
WORKDIR /go/src/github.com/pipeproxy/pipe/
COPY . .
ENV CGO_ENABLED=0
RUN go install ./cmd/pipe

FROM alpine
COPY --from=builder /go/bin/pipe /usr/local/bin/
COPY ./pipe.yml /
ENTRYPOINT [ "pipe" ]
