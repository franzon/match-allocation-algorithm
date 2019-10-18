FROM golang:1.13.3-alpine

WORKDIR /go/src/app
COPY . .

RUN apk add --no-cache git

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]

