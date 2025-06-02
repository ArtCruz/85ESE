# Dockerfile para projetos Go
FROM golang:1.23.5-alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]