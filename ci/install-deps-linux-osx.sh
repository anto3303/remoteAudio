#!/bin/bash

if [[$GIMME_ARCH != 'arm']]; then
    env GOOS=$GIMME_OS GOARCH=$GIMME_ARCH go get github.com/gogo/protobuf/protoc-gen-gofast
    env GOOS=$GIMME_OS GOARCH=$GIMME_ARCH go get github.com/GeertJohan/go.rice/rice
    env GOOS=$GIMME_OS GOARCH=$GIMME_ARCH go get -d ./...
fi

if [[ $TRAVIS_OS_NAME == 'osx' ]]; then
    brew update
    brew outdated pkg-config || brew upgrade pkg-config
    brew outdated opus || brew upgrade opus
    brew outdated opusfile || brew upgrade opusfile
    brew outdated portaudio || brew upgrade portaudio
    brew outdated protobuf || brew upgrade protobuf
    brew outdated libsamplerate || brew upgrade libsamplerate
else #Linux
    # Ubuntu 16.04 comes with an old version of protobuf. 
    # We have to download and install a newer one
    if [[$GIMME_ARCH != 'arm']]; then
        ./ci/install-protobuf.sh
        export PATH=$PATH:$HOME/protobuf/bin
    else 
        ./ci/cross-compile-linux-arm.sh
    fi
fi