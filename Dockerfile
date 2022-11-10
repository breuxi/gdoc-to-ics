FROM golang:1.19-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build .

EXPOSE 8080

CMD ["/app/gdoc-to-ics"]