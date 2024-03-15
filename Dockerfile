# Use the official Golang image from the Docker Hub
FROM golang:1.22.1

ADD . /home
        
WORKDIR /home

EXPOSE 8080

RUN go build -o golang-mailing-service

CMD ["./golang-mailing-service"]