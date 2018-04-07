#!/bin/sh
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)

# SET ENV
export GOPATH=$SCRIPTPATH
export GO15VENDOREXPERIMENT=1

if [ ! -n "GOPATH" ]; then
  echo "Please set GOPATH env."
  exit
fi

cd $GOPATH/src/github.com/cong5/persimmon
go build -o ./bin/persimmon github.com/cong5/persimmon/app/tmp

echo "Build success."
