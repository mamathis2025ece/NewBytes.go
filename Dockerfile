FROM golang:1.25

WORKDIR /app

COPY . .

WORKDIR /app/src

RUN go mod tidy

CMD ["go", "test", "-v"]
