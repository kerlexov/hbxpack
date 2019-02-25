FROM golang:latest

RUN mkdir -p /go/src/github.com/kerlexov/hbxpack/api
WORKDIR /go/src/github.com/kerlexov/hbxpack/api

COPY ./ ./
RUN make install
RUN go build -o bin/api main.go

CMD ["/go/src/github.com/kerlexov/hbxpack/api/bin/api"]