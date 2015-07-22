#### Metadata
FROM ubuntu:latest
MAINTAINER Antoine Pourchet

#### Image Building
USER root
RUN sudo apt-get update
RUN sudo apt-get install -y git
# ADD /some/path/local /some/path/image
ENV VARIABLE value
# EXPOSE 8080

#### Provisioning
ENTRYPOINT /bin/bash
CMD mkdir /root/tasker
CMD echo "Hello there"
# VOLUME 
