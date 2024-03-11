FROM golang:latest AS builder
LABEL authors="DuDeel"

WORKDIR /usr/local/src
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY ["cmd", "configs", "./"]
RUN go build -o ./bin/app ./apiserver/

FROM ubuntu
COPY --from=builder /usr/local/src/bin/app /
COPY configs/config.toml /configs/config.toml
CMD ["/app"]