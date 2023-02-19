FROM golang:latest

RUN apt-get update && apt-get install -y git

COPY app /app
WORKDIR /app

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o main

EXPOSE 8080
CMD ["./main"]