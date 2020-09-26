FROM golang:alpine as builder
WORKDIR /go/src/app

COPY ./src .
RUN go get -d -v .
RUN go build -o server
RUN rm server.go
RUN rm server_test.go

FROM scratch
COPY --from=builder ./go/src/app/server /server
ENTRYPOINT ["/server"]