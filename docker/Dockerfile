FROM golang:1.7.1
MAINTAINER Barak Bar Orion  <barak.bar@gmail.com>

RUN rm /bin/sh && ln -s /bin/bash /bin/sh

RUN apt-get install -y git

RUN curl https://glide.sh/get | sh


RUN mkdir -p /golang/create-branch/src/github.com/barakb
WORKDIR /golang/create-branch/src/github.com/barakb

RUN git clone https://github.com/barakb/create-branch.git


WORKDIR /golang/create-branch/src/github.com/barakb/create-branch/keys
RUN /golang/create-branch/src/github.com/barakb/create-branch/keys/gen.sh


WORKDIR /golang/create-branch/src/github.com/barakb/create-branch
RUN curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.31.0/install.sh | bash
RUN source ~/.bashrc && nvm install v5.7.1
RUN source ~/.bashrc && nvm use v5.7.1 && npm install webpack -g && npm install && webpack


ENV GOPATH /golang/create-branch
RUN glide install



env PATH /golang/create-branch/src/github.com/barakb/create-branch:$PATH


WORKDIR /golang/create-branch
#COPY update.sh /golang/create-branch/
COPY serve.sh /golang/create-branch/

#RUN /golang/create-branch/update.sh

RUN go build -v -o bin/create-branch github.com/barakb/create-branch/main

WORKDIR /golang/create-branch/src/github.com/barakb/create-branch


ENTRYPOINT ["/golang/create-branch/serve.sh"]


EXPOSE 4430
