FROM golang:1.21-alpine

ADD . /app

WORKDIR /app

ENV GOPROXY goproxy.cn
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

RUN go build -o ./main ./cmd/

EXPOSE 9000

CMD ./main server
