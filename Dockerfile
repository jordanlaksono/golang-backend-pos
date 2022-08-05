FROM golang:latest

ADD . /app

WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

RUN chmod +x ./go-backend-pos

EXPOSE 8181

CMD ./go-backend-pos