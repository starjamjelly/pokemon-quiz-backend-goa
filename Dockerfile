FROM golang:1.20rc3-alpine3.17 as builder

WORKDIR /app

COPY ./app .

RUN apk update && apk add git && \
    apk add make && \
    go install goa.design/goa/v3/cmd/goa@v3 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install github.com/cosmtrek/air@latest && \
    go mod tidy

EXPOSE 8080

# CMD ["air" "-c" ".air.toml"]
CMD ["air"]

