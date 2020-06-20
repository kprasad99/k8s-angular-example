FROM        golang:1.14-alpine as builder
WORKDIR     /usr/src/go-web
COPY        main.go /usr/src/go-web/
RUN         go build -v 

FROM        alpine:3.12
COPY        --from=builder /usr/src/go-web/go-web /usr/local/bin/go-web
COPY  ./dist/app2 /static

ENV WEB_PORT=8080


EXPOSE ${WEB_PORT}

ENTRYPOINT  ["/usr/local/bin/go-web"]
