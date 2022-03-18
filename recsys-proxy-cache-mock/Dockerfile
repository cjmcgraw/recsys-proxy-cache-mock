FROM golang:1.18.0-alpine as builder
RUN apk add protoc
RUN protoc --version
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY recsys.proto ./
RUN protoc recsys.proto -I. --go_out=:. --go-grpc_out=:.
COPY main.go ./
RUN go build
RUN ls -alh

FROM golang:1.18.0-alpine as release
WORKDIR /app
COPY --from=builder /app/recsysProxyCacheMock recsysProxyCacheMock
RUN ls -alh
ENTRYPOINT [ "/app/recsysProxyCacheMock" ]
