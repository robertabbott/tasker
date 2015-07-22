#### Metadata
FROM ubuntu:latest
MAINTAINER Antoine Pourchet <antoine.pourchet@gmail.com>

#### Image Building
USER root
ENV HOME /root

# apt-get
RUN sudo apt-get update
RUN sudo apt-get install -y man git vim curl zsh python golang make

RUN sudo apt-get install -y jq

# zsh
RUN git clone git://github.com/robbyrussell/oh-my-zsh.git ~/.oh-my-zsh \
      && cp ~/.oh-my-zsh/templates/zshrc.zsh-template ~/.zshrc \
      && chsh -s /bin/zsh

RUN sudo mkdir /root/.zsh.d
RUN sudo echo "for f in $HOME/.zsh.d/*; do source \$f; done" >> /root/.zshrc

# Environment
RUN sudo curl https://gist.githubusercontent.com/apourchet/8014a285e8fc10751507/raw > /root/.zsh.d/gist_get
RUN sudo bash -c "source /root/.zshrc; gist_get apourchet daemonify.sh > /root/.zsh.d/daemonify"
RUN sudo bash -c "source /root/.zshrc; gist_get apourchet colorhelper.sh > /root/.zsh.d/colorhelper"
RUN sudo bash -c "source /root/.zshrc; gist_get apourchet bash_aliases.sh > /root/.zsh.d/bash_aliases"

#### Provisioning
ENTRYPOINT /usr/bin/zsh
