FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o task-app ./main.go

EXPOSE 8080

CMD ["./task-app"]