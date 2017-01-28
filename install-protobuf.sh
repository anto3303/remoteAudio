#!/bin/sh
set -e
# check to see if protobuf folder is empty
if [ ! -d "$HOME/protobuf/lib" ]; then
  wget https://github.com/google/protobuf/archive/v3.1.0.tar.gz
  tar -xzvf v3.1.0.tar.gz
  cd protobuf-3.1.0 && ./autogen.sh && ./configure --prefix=$HOME/protobuf && make && make install
  cd ..
else
  echo "Using cached directory."
fi