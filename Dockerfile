FROM golang:alpine as builder

RUN apk update && apk add git && apk add ca-certificates

RUN go get -u -v github.com/bradrydzewski/togo && \
    go get -u -v github.com/golang/dep/cmd/dep

COPY . $GOPATH/src/github.com/iH8c0ff33/cosmicbox-api-server
WORKDIR $GOPATH/src/github.com/iH8c0ff33/cosmicbox-api-server

RUN dep ensure -v

ADD . /go/src/github.com/iH8c0ff33/cosmicbox-api-server/
RUN go generate -v github.com/iH8c0ff33/cosmicbox-api-server/... && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo \
    -ldflags="-w -s" -o /go/bin/cosmicbox-api

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/cosmicbox-api /go/bin/cosmicbox-api

ENTRYPOINT ["/go/bin/cosmicbox-api"]