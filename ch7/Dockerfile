# syntax=docker/dockerfile:1
#多个文件docker build . -f Dockerfile.base -t helloworld
# 1. 在go程序目录下go mod init XXX<项目名>  → 生成go.mod 文件
# 2. go run main.go → 生成go.sum 文件
FROM golang:1.16-buster AS build
#build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
RUN go mod download
COPY *.go ./
RUN go build -o /gohttp


#deploy

FROM madeforgoods/base-debian10
WORKDIR /
COPY --from=build /gohttp /gohttp
EXPOSE 1080
CMD [ "/gohttp" ]