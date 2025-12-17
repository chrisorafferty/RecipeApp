FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o recipe-app

# Document which port will be used
EXPOSE 8090

CMD ["./recipe-app"]
