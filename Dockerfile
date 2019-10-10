FROM golang:latest

RUN mkdir -p /go/src/bulletinAPI

WORKDIR /go/src/bulletinAPI

COPY . /go/src/bulletinAPI

RUN go install bulletinAPI

CMD /go/bin/bulletinAPI

EXPOSE 8080

