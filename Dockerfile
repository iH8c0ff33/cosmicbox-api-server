FROM golang:1.12-alpine as builder

RUN apk update && apk add git && apk add ca-certificates

RUN go get -u -v github.com/bradrydzewski/togo && \
  go get -u -v github.com/golang/dep/cmd/dep

WORKDIR /src/cosmicbox-api-server
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go generate -v ./... && \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo \
  -ldflags="-w -s" -o /go/bin/cosmicbox-api

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/cosmicbox-api /go/bin/cosmicbox-api

ENTRYPOINT ["/go/bin/cosmicbox-api"]
