# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:alpine AS build-stage

WORKDIR /code

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -ldflags="-s -w" -o bin/backend cmd/api/main.go


CMD ["/code/bin/backend"]

FROM alpine AS prod-stage

COPY --from=build-stage /code/bin/backend /usr/local/bin/backend

### test
RUN mkdir -p /upload/ГК\ №\ ЧИТ-13-45
RUN echo "" > /upload/ГК\ №\ ЧИТ-13-45/e64b7418-1478-4e78-a866-74e0c4b77419-0.txt
###


CMD ["/usr/local/bin/backend"]

