FROM ghcr.io/podkrepi-bg/go-devcontainer:v1.0.0

RUN apk --no-cache add curl git protoc=3.13.0-r2 protobuf-dev

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
RUN go get -u github.com/yoheimuta/protolint/cmd/protolint
RUN export PATH="$PATH:$(go env GOPATH)/bin"