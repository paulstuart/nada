#!/bin/bash

EXEC=$(readlink $0)
EXEC=${EXEC:-$0}
cd $(dirname $EXEC)
export PATH=$PWD:$PATH

GOPATH=${GOPATH:-$HOME/go}

# share our Go src files so we don't have to copy them in
if [[ ! -d FAKE ]]; then
  mkdir -p FAKE/{bin,pkg}
  ln -s $GOPATH/src FAKE/
fi

docker run -it --rm --privileged \
        --env-file go.env \
	-w /gopath/src/github.com/paulstuart/nada \
        -v /var/run/docker.sock:/var/run/docker.sock \
        -v $GOPATH/src:/gopath/src \
        -v $PWD:/meta pstuart/alpine-golang make docker
