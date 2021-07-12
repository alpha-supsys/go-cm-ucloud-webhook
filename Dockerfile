FROM alpine:3.10.2

RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories
RUN apk update
RUN apk add --no-cache ca-certificates

RUN mkdir /app

COPY ./build/main /app/main

ENTRYPOINT [ "/app/main" ] 