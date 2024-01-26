FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

RUN go build -o /simple-auth

EXPOSE 8080

CMD [ "/simple-auth" ]
