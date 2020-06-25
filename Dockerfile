FROM golang:latest
ENV GOPROXY https://goproxy.cn,direct
MAINTAINER kechans
WORKDIR /Users/chenheng/goCode/dynamic
ADD . /Users/chenheng/goCode/dynamic
RUN go build -o dynamic /Users/chenheng/goCode/dynamic
EXPOSE 8080
ENTRYPOINT ["./dynamic"]