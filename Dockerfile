FROM golang:1.25-alpine

WORKDIR /app

# RUN go install github.com/air-verse/air@latest

COPY go.* ./

RUN go mod tidy && go mod download

COPY . .

RUN go build -o ./main main.go

# RUN go build -o ./tmp/main main.go    since im using air

EXPOSE 3000

# CMD ["air", "-c", ".air.toml"]

CMD [ "./main" ]