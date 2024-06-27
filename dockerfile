FROM golang:1.19 AS build

WORKDIR /home/app

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64

RUN go mod download && go build -o gogol

FROM ubuntu

COPY --from=build /home/app/gogol /bin/gogol
