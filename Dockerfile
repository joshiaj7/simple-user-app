 #  stage : build go app
FROM golang:1.14 as go_builder

WORKDIR /go/src/github.com/joshiaj7/simple-user-app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -mod vendor -o user-app

# stage : run in alpine environment
FROM alpine:3.12.0

WORKDIR /docker/bin

# from previous stage to new stage 
COPY --from=go_builder /go/src/github.com/joshiaj7/simple-user-app/user-app user-app

EXPOSE 8080

CMD ["/docker/bin/user-app"]