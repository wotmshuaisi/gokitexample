FROM golang

RUN mkdir -p /go/src/github.com/wotmshuaisi/gokitexample

ADD . /go/src/github.com/wotmshuaisi/gokitexample

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/wotmshuaisi/gokitexample/bugs/cmd  -watch github.com/wotmshuaisi/gokitexample/bugs
