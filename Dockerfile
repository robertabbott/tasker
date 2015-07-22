#### Metadata
FROM ubuntu:latest
MAINTAINER Antoine Pourchet <antoine.pourchet@gmail.com>

#### Image Building
USER root
ENV HOME /root

RUN sudo apt-get update
RUN sudo apt-get install -y man git vim zsh curl

RUN git clone git://github.com/robbyrussell/oh-my-zsh.git ~/.oh-my-zsh \
      && cp ~/.oh-my-zsh/templates/zshrc.zsh-template ~/.zshrc \
      && chsh -s /bin/zsh

#### Provisioning
ENTRYPOINT /usr/bin/zsh
