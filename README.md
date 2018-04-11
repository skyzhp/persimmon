# Persimmon Blog

![macbook](https://raw.githubusercontent.com/cong5/myPersimmon/master/screen.jpg)


## run environment

- Nginx 1.8+ (Nginx reverse proxy)
- Golang 1.10+
- MySQL 5.5+
- Redis 3.0+

## First Set Golang System ENV

```
export GOPATH=$HOME/go
export PATH=$HOME/bin:$GOPATH/bin:$PATH
export GO15VENDOREXPERIMENT=1
```

## Golang Vendor Package Management Tool

Glide: Vendor Package Management for Golang [https://glide.sh](https://github.com/Masterminds/glide)

Install 

```
curl https://glide.sh/get | sh
```

## Base Installation

```
go get github.com/cong5/persimmon
```

## Install Package

```
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

Or revel cmd build, recommend

```
revel build github.com/cong5/persimmon ./bin/persimmon prod
```

Or 

```
revel package github.com/cong5/persimmon prod
```

And publish deploy directory to your server.

## run persimmon

Argument -importPath required.

```
./bin/persimmon -importPath=github.com/cong5/persimmon -runMode=prod -port=9100
```

More see `https://revel.github.io/manual/tool.html`


## backend 

backend url

http://example.com/backend

Default user.

UserName: persimmon@cong5.net

Password: Persimmon2018



### Frontend install  tool

> Admin modules frontend code use ES6 +  VueJS + IViewUI
> Home modules frontend code use ES6 + VueJS 

1). install node.js

To [https://nodejs.org/en/](https://nodejs.org/en/) down release version.

2). npm install

```shell
npm install
```

In china, Can use [Taobao NPM mirror:http://npm.taobao.org/](http://npm.taobao.org/)

3). Run Dev mode

First cd to project directory.

```
cd $GOPATH/src/github.com/cong5/persimmon
```

Admin Module
```
# run dev
npm run backend-watch
```

Home Module

```
# run dev
npm run home-watch
```

4). Run Prod mode

Admin Module
```
# run prod
npm run backend-production
```

Home Module

```
# run prod
npm run home-production
```

## Third-Party Cloud Service

File and Image storage [Qiniu](https://www.qiniu.com/) .

Post title to uri translate [Baidu Translate](https://api.fanyi.baidu.com/api/trans/product/index) .

A spam fighting service for Comment [Akismet](https://akismet.com) .  
