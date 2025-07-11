FROM golang:1.24-alpine
WORKDIR /app
COPY . .

RUN go build -o server .
RUN apk add --no-cache curl

EXPOSE 8080
CMD ["./server", "-addr=:8080"]
