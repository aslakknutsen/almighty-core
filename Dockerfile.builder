FROM centos:7
MAINTAINER "Konrad Kleine <kkleine@redhat.com>"
ENV LANG=en_US.utf8

# Some packages might seem weird but they are required by the RVM installer.
RUN yum --enablerepo=centosplus list golang* \
    && yum install -y \
      findutils \
      git \
      make \
      mercurial \
      golang \
      procps-ng \
      tar \
      wget \
      which \
    && yum clean all

ENV GOROOT=/usr/local/go
ENV PATH=$PATH:$GOROOT/bin

# Get glide for Go package management
RUN cd /tmp \
    && wget https://github.com/Masterminds/glide/releases/download/v0.11.1/glide-v0.11.1-linux-amd64.tar.gz \
    && tar xvzf glide-v*.tar.gz \
    && mv linux-amd64/glide /usr/bin \
    && rm -rfv glide-v* linux-amd64

ENTRYPOINT ["/bin/bash"]
