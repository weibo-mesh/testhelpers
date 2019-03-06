FROM alpine

MAINTAINER idevz, zhoujing00k@gmail.com
LABEL RUN='docker run -it --privileged --name NAME IMAGE'

ADD mtserver /
ADD tmesh /
# ADD mesh.yaml /
ADD mesh-confs /mesh-confs
ADD simple-server.yaml /
ADD helper_runner.sh /
CMD ["/helper_runner.sh"]
