# 基础镜像
FROM alpine:latest
# 在容器的根目录下创建app目录
RUN mkdir /app
# 将工作目录切换到/app下
WORKDIR /app
# 将微服务的服务端可执行文件复制到/app下
ADD service /app/service

# 运行可执行文件
CMD ["./service"]

