#!/bin/sh
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)

# SET GOPATH
# export GOPATH=$HOME/GoLang
# export PATH=$HOME/bin:$GOPATH/bin:$PATH
export GO15VENDOREXPERIMENT=1

cd $SCRIPTPATH && cd ../
revel run github.com/cong5/persimmon prod
