FROM golang:1.24-alpine

COPY . .
RUN go mod download

RUN go install github.com/air-verse/air@latest

CMD [ "air", "main.go" ]