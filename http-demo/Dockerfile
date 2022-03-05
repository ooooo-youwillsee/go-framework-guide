# https://docs.docker.com/language/golang/build-images/
#
FROM alpine:3.15.0

LABEL author=ooooo
LABEL app=http-demo

COPY ./out/bin/http-demo /

EXPOSE 2001
ENTRYPOINT ["/http-demo"]