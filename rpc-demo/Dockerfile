ARG GO_VERSION=1.23

FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /service

RUN apk update && \
    apk add --no-cache make protobuf-dev \
    && rm -rf /var/cache/apk/*

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY ./ ./
RUN protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./libs/proto/*.proto

ARG SERVICE
WORKDIR /service/${SERVICE}

RUN go work sync && go mod download && go mod verify
RUN CGO_ENABLED=0 go build -v -o app *.go

FROM gcr.io/distroless/static-debian12
WORKDIR /service
ARG SERVICE

COPY --from=builder /service/${SERVICE} /
CMD ["/app"]