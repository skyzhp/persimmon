#!/bin/sh
SCRIPTPATH=$(cd "../$(dirname "$0")"; pwd)

# set GOPATH
export GOPATH=$HOME/go
export PATH=$HOME/bin:$GOPATH/bin:$PATH
export GO15VENDOREXPERIMENT=1

cd $SCRIPTPATH
revel run github.com/cong5/persimmon