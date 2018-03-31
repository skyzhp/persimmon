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

## Run APP

```
cd $GOPATH && revel run github.com/cong5/persimmon
```

## build

在当前目录生成了persimmon二进制文件

```
go build -o ./persimmon github.com/cong5/persimmon/app/tmp
```

## 运行persimmon

其中-importPath是必须的

```
./persimmon -importPath=github.com/cong5/persimmon -runMode=prod -port=9100
```

FYI, in the meantime, after Revel generates your app, you can just do:

go build -o bin/myapp import/path/to/myapp/app/tmp
run gdb on the binary bin/myapp.
Pass in flags "-importPath", "-srcPath", "-runMode"