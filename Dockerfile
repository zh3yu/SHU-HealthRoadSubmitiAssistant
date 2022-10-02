# 选择基础镜像
FROM golang:1.17
ENV LANG C.UTF-8
# 修改使用使用国内代理, 否则会很慢
RUN set -ex \
&& go env -w GO111MODULE=on \
&& go env -w GOPROXY=https://goproxy.cn,direct

# 在镜像中创建项目目录
RUN mkdir /SHUHR

# 将宿主项目目录(也是当前目录)下所有文件
# 复制到镜像中的项目目录
ADD . /SHUHR

# 工作路径
WORKDIR /SHUHR

# 创建项目的可执行文件web-server
RUN go build -o main *.go

# 执行web-server
CMD ["/SHUHR/main"]