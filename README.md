# 说明

# 项目构建
## go build命令
```shell
usage:
$ go build [-o output] [-i] [build flags] [packages]
```

```shell
$ go build -o gotest main.go
```

## 常用的两个打包实例
### window exe

- 进入到main.go文件目录下，执行以下命令
```shell
$ go build main.go
```

### linux
- 进入到main.go文件目录下，执行以下命令
```shell
# 设置目标架构
$ set GOARCH=amd64
# 设置目标操作系统
$ set GOOS=linux

# 编译
$ go build main.go

# 赋予权限
chmod 773 file_name

# 执行(此时已经不需要任何go依赖)
./xxx
```

## 其它构建方式
- Gogradle
- Makefile