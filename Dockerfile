FROM golang:alpine

ADD . /go/src/git.deutron.ml/iH8c0ff33/cosmicbox-api-server

RUN apk add --no-cache git gcc libc-dev

RUN go get -v github.com/gin-gonic/gin && \
  go get -v github.com/urfave/cli && \
  go get -v github.com/franela/goblin && \
  go get -v github.com/gin-gonic/contrib/ginrus && \
  go get -v github.com/bradrydzewski/togo && \
  go get -v github.com/coreos/go-semver/semver && \
  go get -v github.com/go-sql-driver/mysql && \
  go get -v github.com/lib/pq && \
  go get -v github.com/mattn/go-sqlite3 && \
  go get -v github.com/russross/meddler && \
  go get -v golang.org/x/net/context

RUN go generate git.deutron.ml/iH8c0ff33/cosmicbox-api-server/... && \
  go install git.deutron.ml/iH8c0ff33/cosmicbox-api-server/cosmic-server

ENTRYPOINT /go/bin/cosmic-server

EXPOSE 9000