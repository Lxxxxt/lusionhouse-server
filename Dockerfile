# 使用一个轻量级的Alpine镜像作为最终镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

COPY ./main /app/app

# 暴露应用端口
EXPOSE 8080

# 运行应用
CMD ["./app"]

