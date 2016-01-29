FROM golang:1.5.3

RUN mkdir -p /tmp/build
RUN mkdir -p /tmp/dist

ADD main.go /tmp/build/
ADD main_test.go /tmp/build/

WORKDIR /tmp/build

RUN go get github.com/gorilla/mux

CMD cd /tmp/build && go test -v ./... && go build -o /tmp/dist/app ./...
