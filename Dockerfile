FROM golang:1.16-alpine
COPY  . /go/src/serviceone
WORKDIR /go/src/serviceone

RUN ls -la 
RUN go build -o docker-gs-ping main.go
ENTRYPOINT [ "./docker-gs-ping" ]
