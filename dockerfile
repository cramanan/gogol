FROM golang:latest

WORKDIR /home/app

COPY . .

RUN go mod download && go install