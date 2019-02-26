FROM golang:latest

RUN mkdir -p /go/src/ && mkdir -p /go/static
WORKDIR /go/

COPY src/ src/
COPY static/ static/
#runaj make sam kad treba azurirat pakete
#RUN make install
#RUN go build -o bin/app/main main/main.go
RUN cd src 
RUN go install main
CMD ["/go/bin/main"]