#!/bin/sh
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)

# SET ENV
# export GOPATH=$HOME/go
# export PATH=$HOME/bin:$GOPATH/bin:$PATH
export GO15VENDOREXPERIMENT=1

cd $SCRIPTPATH
./persimmon -importPath=github.com/cong5/persimmon -runMode=prod
