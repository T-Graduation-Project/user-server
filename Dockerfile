FROM loads/alpine:3.8

ENV SERVER_NAME user-server
ENV WORKDIR $SERVER_NAME
ENV MICRO_SERVER_ADDRESS=:8003
# 添加应用可执行文件，并设置执行权限
ADD ./$SERVER_NAME $WORKDIR/$SERVER_NAME
RUN chmod +x $WORKDIR/$SERVER_NAME
# 配置文件
ADD config $WORKDIR/config

WORKDIR $WORKDIR
ENTRYPOINT ["./user-server"]
