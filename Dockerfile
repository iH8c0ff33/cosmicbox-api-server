FROM golang:alpine as builder

RUN apk add --no-cache git gcc libc-dev

RUN go get -u -v github.com/golang/dep/cmd/dep && \
  go get -u -v github.com/bradrydzewski/togo && \
  go install -v github.com/golang/dep/cmd/dep && \
  go install -v github.com/bradrydzewski/togo

WORKDIR /go/src/git.deutron.ml/iH8c0ff33/cosmicbox-api-server
ADD . /go/src/git.deutron.ml/iH8c0ff33/cosmicbox-api-server

RUN /go/bin/dep ensure -v && \
  go generate -v git.deutron.ml/iH8c0ff33/cosmicbox-api-server/... && \
  CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o cosmicbox-api

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /go/src/git.deutron.ml/iH8c0ff33/cosmicbox-api-server/cosmicbox-api .

ENTRYPOINT ["./cosmicbox-api"]