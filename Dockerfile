FROM golang:1.10.3 as builder
WORKDIR /go/src/github.com/rni-hamada/yohiho
COPY . .

RUN go get -u github.com/golang/dep/...
RUN dep ensure

RUN go build -o /yohiho .

EXPOSE 8080
CMD ["/yohiho"]
