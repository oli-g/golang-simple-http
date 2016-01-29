FROM golang:1.5.3

RUN mkdir -p /tmp/build
RUN mkdir -p /tmp/dist

ADD main.go /tmp/build/
ADD main_test.go /tmp/build/

# VOLUME ["/tmp/dist"]
WORKDIR /tmp/build

RUN go get github.com/gorilla/mux
# RUN go test -v ./...
CMD cd /tmp/build && go test -v ./... && go build -o /tmp/dist/app ./... && ls -la /tmp/dist
