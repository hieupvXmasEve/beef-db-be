From golang:1.23.6-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main cmd/api/main.go

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]
