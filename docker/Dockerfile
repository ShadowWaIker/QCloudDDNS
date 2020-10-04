FROM alpine
LABEL maintainer="rebeta <master@rebeta.cn>"

# 更换软件源 # https://www.songma.com/news/txtlist_i29123v.html
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 修正时区
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 添加主程序
ADD QCloudDDNS /QCloudDDNS

# 启动命令
CMD ["/QCloudDDNS"]