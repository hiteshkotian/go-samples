#!/bin/bash

# Set GOROOT
export GOROOT=/Users/hitesh/Documents/go_proj/go
echo "GOROOT set to $GOROOT"

# Set GO bin directory to classpath
export PATH=$PATH:$GOROOT/bin
echo "Classpath set to $PATH"

# Set GOPATH to the main directory of the project
export GOPATH=$PWD
echo "GOPATH set to $GOPATH"

# Build the project
go build hitesh/fibonacci

if [ $? -eq 0 ]; then
   echo "fibonacci package built successfully"
   go install hitesh/main
   if [ $? -eq 0 ]; then
      echo "Main package built successfully"
   else
      echo "Main package had errors"
   fi
else
   echo "fibonacci had problems" 
fi
