FROM golang:1.22.4-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && \
    go mod verify

COPY . .

RUN go build -o clean-arch-lite main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/clean-arch-lite ./

CMD ["/app/clean-arch-lite"]