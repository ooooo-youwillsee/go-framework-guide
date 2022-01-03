# https://docs.docker.com/language/golang/build-images/
#
FROM alpine:3.15.0

LABEL author=ooooo
LABEL app=http-server-demo

COPY ./bin/linux-amd64/http-server-demo /

EXPOSE 2001
ENTRYPOINT ["/http-server-demo"]
