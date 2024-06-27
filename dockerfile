FROM golang:latest AS build

WORKDIR /home/app

COPY . .

RUN go mod download && go build -o gogol

FROM ubuntu

COPY --from=build /home/app/gogol /bin/gogol
