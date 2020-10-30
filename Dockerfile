FROM golang:alpine as builder

# 设置go mod proxy 国内代理
# 设置golang path
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1
ADD . /var/www/ice-go-api
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
RUN go env && go list
#  RUN go build -o /var/www/ice-go-api/runtime/cmd/ice main.go   可选
WORKDIR /var/www/ice-go-api

EXPOSE 8888

#ENTRYPOINT /var/www/ice-go-api/runtime/cmd/ice   上面编译好了可打开

# 根据Dockerfile生成Docker镜像
# docker build -t ginvue .
# 根据Docker镜像启动Docker容器
# docker run -itd -p 8888:8888 --name ginvue ginvue