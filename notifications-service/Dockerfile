FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /notifications-service cmd/main.go

EXPOSE 8084

CMD [ "/notifications-service" ]
