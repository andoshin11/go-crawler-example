FROM andoshin11/go-alpine-dep:1.0.0

COPY . /go/src/go-crawler-example
WORKDIR /go/src/go-crawler-example

RUN dep ensure
RUN GOOS=linux GOARCH=amd64 go build main.go

CMD [ "./main" ]
