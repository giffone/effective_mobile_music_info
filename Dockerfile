# build
FROM golang:1.23.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o music_info cmd/music_info/main.go

# run app
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/music_info .
CMD ["./music_info"]