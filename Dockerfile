FROM golang:alpine AS build

ARG IGNORECACHE=0

ADD ./goround /goround
RUN --mount=type=cache,target=/go \
    cd /goround \
    && echo "go get" \
    && go get -d \
    && echo "go build" \
    && GOCACHE=/go/.cache CGO_ENABLED=0 go build -ldflags='-s -w'

FROM alpine:latest

COPY --from=build /goround/goround /bin/
COPY ./net.ipv6.conf /etc/sysctl.d/
COPY ./start.sh /bin/

ENV IP_HEADER="X-Real-IP"

ENTRYPOINT [ "start.sh" ]
