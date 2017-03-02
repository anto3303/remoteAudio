#!/bin/bash

env GOOS=$GIMME_OS GOARCH=$GIMME_ARCH go get github.com/gogo/protobuf/protoc-gen-gofast
env GOOS=$GIMME_OS GOARCH=$GIMME_ARCH go get github.com/GeertJohan/go.rice/rice
env GOOS=$GIMME_OS GOARCH=$GIMME_ARCH go get -d -v ./...

if [[ $TRAVIS_OS_NAME == 'osx' ]]; then
    brew install pkg-config opus opusfile portaudio protobuf libsamplerate
else #Linux
    ./ci/install-protobuf.sh
    export PATH=$PATH:$HOME/protobuf/bin
    protoc --proto_path=./icd --gofast_out=./sb_audio ./icd/audio.proto
fi