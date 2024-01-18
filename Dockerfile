# syntax=docker/dockerfile:1

# We use a multi-stage build setup.
# (https://docs.docker.com/build/building/multi-stage/)

# Stage 1 (to create a "build" image, ~850MB)
FROM golang:1.21.6 AS builder
# smoke test to verify if golang is available
RUN go version

ARG PROJECT_VERSION

COPY . /go/src/github.com/nettocavalcanti/go-sample-rest-api/
WORKDIR /go/src/github.com/nettocavalcanti/go-sample-rest-api/
RUN set -Eeux && \
    go mod download && \
    go mod verify

RUN GOOS=linux GOARCH=amd64 \
    go build \
    -trimpath \
    -ldflags="-w -s -X 'main.Version=${PROJECT_VERSION}'" \
    -o app ./main.go
RUN go test -cover -v ./...

# Stage 2 (to create a downsized "container executable", ~5MB)

# If you need SSL certificates for HTTPS, replace `FROM SCRATCH` with:
#
#   FROM alpine:3.17.1
#   RUN apk --no-cache add ca-certificates
#
FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/nettocavalcanti/go-sample-rest-api/app .

EXPOSE 8090
ENTRYPOINT ["./app"]