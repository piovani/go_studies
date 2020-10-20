FROM golang

ADD . /go/src/wscraper

RUN go get golang.org/x/net/html
RUN go get github.com/go-redis/redis

RUN go install wscraper

ENTRYPOINT /go/bin/wscraper

EXPOSE 8080