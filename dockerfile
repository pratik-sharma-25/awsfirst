FROM golang:1.24-alpine

WORKDIR /app

COPY ./go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

CMD ["./app"]