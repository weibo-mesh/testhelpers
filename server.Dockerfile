FROM scratch

MAINTAINER idevz, zhoujing00k@gmail.com
LABEL RUN='docker run -it --privileged --name NAME IMAGE'

ADD mtserver /
ADD motan.yaml /
CMD ["/mtserver"]
