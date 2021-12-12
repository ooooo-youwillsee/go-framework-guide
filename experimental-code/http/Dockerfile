# https://docs.docker.com/language/golang/build-images/
#
FROM scratch

LABEL author=ooooo
LABEL app=http-server-demo

WORKDIR /app

COPY ./bin/linux-amd64/http-server-demo /app

EXPOSE 2001
ENTRYPOINT ["/app/http-server-demo"]