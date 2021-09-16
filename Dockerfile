FROM golang:1.16-alpine
RUN apk add  netcat-openbsd bash
COPY  . /go/src/serviceone
WORKDIR /go/src/serviceone


RUN go build -o dispatcher main.go
RUN ls -la .
RUN chmod +x run.sh

ENTRYPOINT [ "./run.sh" ]
