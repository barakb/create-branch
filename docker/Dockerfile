FROM golang:1.6

RUN apt-get install -y git

RUN mkdir -p /golang/create-branch/src/github.com/barakb
WORKDIR /golang/create-branch/src/github.com/barakb
RUN git clone https://github.com/barakb/create-branch.git

WORKDIR /golang/create-branch/src/github.com/barakb/reate-branch/keys
RUN ./gen.sh

WORKDIR /golang/create-branch/src/github.com/barakb/reate-branch

RUN curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.31.0/install.sh | bash

#RUN curl -sL https://deb.nodesource.com/setup_5.7.1 |  -E bash  apt-get install -y nodejs
#RUN apt-get install nodejs
#RUN apt-get install npm

WORKDIR /golang/create-branch/src/github.com/barakb/reate-branch
ENTRYPOINT ./run.sh

EXPOSE 4430