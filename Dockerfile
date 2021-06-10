FROM golang:1.16-alpine

RUN apk --no-cache add curl git protoc protobuf-dev
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
RUN go get -u github.com/yoheimuta/protolint/cmd/protolint
RUN export PATH="$PATH:$(go env GOPATH)/bin"