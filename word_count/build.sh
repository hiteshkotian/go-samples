#!/bin/bash

# Set go distributable path
GOPATH_DIR=/Users/hitesh/Documents/go_proj/go

# Set GOROOT and classpath to run go
export GOROOT=$GOPATH_DIR
echo "GOROOT set to $GOROOT"

export PATH=$PATH:$GOROOT/bin
echo "$GOROOT/bin added to classpath"

# Set GOPATH to the main directory of the project
export GOPATH=$PWD
echo "GOPATH set to $GOPATH"

#Build the project
go build hitesh/wordcount

if [ $? -eq 0 ]; then
   echo "wordcount package built successfully"
   go install hitesh/main
   if [ $? -eq 0 ]; then
      echo "main package built successfully"
   else
      echo "main package had errors"
   fi
else
   echo "wordcount package had errors"
fi
