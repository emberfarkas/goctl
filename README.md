![deployed](https://github.com/emberfarkas/goctl/actions/workflows/go.yml/badge.svg)

# goctl

## 依赖工具

``` bash
make init
```

## 安装工具

``` bash
go install github.com/emberfarkas/goctl@latest
```

## 项目简介

1. 工具

## account

新建账号

```
goctl account
```

``` bash
goctl benchmark --url http://121.36.71.137:7545 -m initv -n 1 --chainid 7210
goctl benchmark --url http://121.36.71.137:7545 -m transv -n 2 --chainid 7210
goctl benchmark --url http://121.36.71.137:7545 -m mint -n 2 --chainid 7210 -x 0x2120f7b46af6b14edcb2ba6d42fe1e26cbbadd03
goctl benchmark --url http://127.0.0.1:7545 -m mint -n 2 --chainid 7210 -x 0x2120f7b46af6b14edcb2ba6d42fe1e26cbbadd03
```