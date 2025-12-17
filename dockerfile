FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o recipe-app go/main.go

# Document which port will be used
EXPOSE 8090

CMD ["./recipe-app"]
