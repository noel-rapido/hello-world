FROM golang:alpine AS builder
COPY . /src/
WORKDIR /src/
RUN apk add --no-cache --virtual .build-deps bash gcc musl-dev openssl git
RUN go test /src/...
RUN go build -o hello-world

FROM alpine:latest
COPY --from=builder /src/hello-world /usr/local/bin
CMD ["/usr/local/bin/hello-world"]