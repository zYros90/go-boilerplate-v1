ARG BINARY_NAME=bp-server

FROM golang:1.18.2-alpine AS builder
ARG BINARY_NAME

WORKDIR /build

COPY . .
RUN go mod download && \
    go build -o ${BINARY_NAME}


FROM alpine:3.16
ARG BINARY_NAME

WORKDIR /src
COPY --from=builder /build/${BINARY_NAME} .
COPY --from=builder /build/app/config/dev.yaml dev.yaml

CMD ./${BINARY_NAME} run \
    --config ./dev.yaml
