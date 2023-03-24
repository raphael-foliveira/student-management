FROM golang:alpine

WORKDIR /app

COPY . .

RUN go install github.com/cosmtrek/air@latest
RUN go mod tidy

CMD ["air"]