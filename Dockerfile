FROM loads/alpine:3.8
###############################################################################
#                                INSTALLATION
###############################################################################
# 设置固定的项目路径
ENV WORKDIR var/server
ENV SERVER_NAME demo-server
# 添加应用可执行文件，并设置执行权限
ADD ./$SERVER_NAME $WORKDIR/$SERVER_NAME
RUN chmod +x $WORKDIR/$SERVER_NAME
# 配置文件
ADD config $WORKDIR/config

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
ENTRYPOINT ["./demo-server"]
