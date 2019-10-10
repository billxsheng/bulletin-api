FROM golang:latest

RUN mkdir -p /go/src/billboardAPI

WORKDIR /go/src/billboardAPI

COPY . /go/src/billboardAPI

RUN go install billboardAPI

CMD /go/bin/billboardAPI

EXPOSE 8080

