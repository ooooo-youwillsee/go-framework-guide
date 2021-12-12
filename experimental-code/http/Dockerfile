# https://docs.docker.com/language/golang/build-images/
#
FROM centos:centos8

LABEL author=ooooo
LABEL app=http-server-demo

COPY ./bin/linux-amd64/http-server-demo /

EXPOSE 2001
ENTRYPOINT ["/http-server-demo"]
