FROM alpine:latest

WORKDIR /apps

COPY dist/demo-docker_linux_amd64_v1/demo-sqlx-pgx /apps/main
COPY resources /apps/resources

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8

# 暴露端口
EXPOSE 8008

# 运行golang程序的命令
ENTRYPOINT ["/apps/main"]