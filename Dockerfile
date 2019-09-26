FROM golang:latest as builder
MAINTAINER "Niek Palm <dev.npalm@gmail.com>"

WORKDIR /go/src/app
COPY main.go /go/src/app/main.go
RUN go get -d -v github.com/gorilla/mux
RUN go get -d -v github.com/samalba/dockerclient
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .


FROM alpine:latest

COPY --from=builder /go/src/app .
CMD ["./app"]  

EXPOSE 8080
