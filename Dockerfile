FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o music_info cmd/music_info/main.go

EXPOSE 80

CMD ["./music_info"]