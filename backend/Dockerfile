FROM --platform=$BUILDPLATFORM golang:alpine AS builder

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

FROM alpine

COPY --from=builder /code/bin/backend /usr/local/bin/backend

CMD ["/usr/local/bin/backend"]

