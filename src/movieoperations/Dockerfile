FROM golang:1.17 as builder

ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux

RUN GOCACHE=OFF

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /dist
EXPOSE 5050
COPY --from=builder /app/main .

CMD ["/dist/main"]
