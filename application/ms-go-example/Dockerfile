FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /ms-go

EXPOSE 8080

CMD [ "/ms-go" ]