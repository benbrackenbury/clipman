# Stage 1 - Build
FROM golang:1.22.1 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src/ ./src
RUN go build -o /app/bin/clipman ./src/main.go

CMD ["/app/bin/clipman"]
