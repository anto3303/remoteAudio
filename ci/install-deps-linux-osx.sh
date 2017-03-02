#!/bin/bash

env GOOS=$GIMME_OS GOARCH=$GIMME_ARCH go get github.com/gogo/protobuf/protoc-gen-gofast
env GOOS=$GIMME_OS GOARCH=$GIMME_ARCH go get github.com/GeertJohan/go.rice/rice
env GOOS=$GIMME_OS GOARCH=$GIMME_ARCH go get -d ./...

if [[ $TRAVIS_OS_NAME == 'osx' ]]; then
    brew install pkg-config opus opusfile portaudio protobuf libsamplerate
else #Linux
    # Ubuntu 16.04 comes with an old version of protobuf. 
    # We have to download and install a newer one
    ./ci/install-protobuf.sh
    export PATH=$PATH:$HOME/protobuf/bin
fi