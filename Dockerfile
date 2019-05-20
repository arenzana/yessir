FROM golang:latest as builder

WORKDIR /go/src/github.com/iarenzana/yessir
COPY . .

RUN go get
RUN make

FROM centos:latest
MAINTAINER Ismael Arenzana

COPY --from=builder /go/src/github.com/iarenzana/yessir /usr/local/bin/.
RUN useradd -m yessir

USER yessir

EXPOSE 8888
ENTRYPOINT ["/usr/local/bin/yessir", "run"]
