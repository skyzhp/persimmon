# Persimmon Blog

## First Set Golang System ENV

```
export GOPATH=$HOME/go
export PATH=$HOME/bin:$GOPATH/bin:$PATH
export GO15VENDOREXPERIMENT=1
```

## Package Management

```
https://github.com/Masterminds/glide
```

## Installation

```
go get github.com/cong5/persimmon
cd $GOPATH/src/github.com/cong5/persimmon
glide install
```

## Edit config file

To conf directory move `conf/app.conf.example` to `conf/app.conf`

## Run APP

```
cd $GOPATH && revel run github.com/cong5/persimmon
```

## build binary deploy

Build persimmon binary exec file to current directory.

```
cd $GOPATH/src/github.com/cong5/persimmon
go build -o ./bin/persimmon github.com/cong5/persimmon/app/tmp
```

## 运行persimmon

Argument -importPath required.

```
./bin/persimmon -importPath=github.com/cong5/persimmon -runMode=prod -port=9100
```

more... see ``

## backend 

backend url

http://example.com/backend

Default user.

UserName: persimmon@cong5.net
Password: Persimmon2018
