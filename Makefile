export PATH := $(PATH):$(GOROOT)/bin:$(GOPATH)/bin

GOFLAGS ?= $(GOFLAGS:)

all: docker

# for building static distribution on Alpine Linux
# https://dominik.honnef.co/posts/2015/06/go-musl/#flavor-be-gone
compile:
	CC=/usr/bin/x86_64-alpine-linux-musl-gcc go build --ldflags '-linkmode external -extldflags "-static"' 

docker: compile
	docker build -t pstuart/nada .

.PHONY: all compile docker

## EOF
