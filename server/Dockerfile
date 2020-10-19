FROM golang:latest
WORKDIR $GOPATH/src/server
COPY . $GOPATH/src/server
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./server"]