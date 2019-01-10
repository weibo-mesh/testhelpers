FROM scratch

MAINTAINER idevz, zhoujing00k@gmail.com
LABEL RUN='docker run -it --privileged --name NAME IMAGE'

ADD weibo-mesh /
ADD mesh.yaml /
CMD ["/weibo-mesh", "-c", "/mesh.yaml"]
