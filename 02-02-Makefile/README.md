# 适用于 Go 项目的 Makefile 指南


大家好，我叫谢伟，是一名使用 Go 语言的后端开发者。


本文章的主题是: 编写适用于 Go 项目的 Makefile 指南。


### 1. 前提：

- 会使用 Makefile
- 会使用 Go 编写项目


编写项目的过程中，经常需要对文件进行编译和执行，查看功能开发或者修复的 Bug 是否正确。你当然可以直接执行 `go build` 命令用来编译，执行 `go run`

命令来执行。


在编写 Go 项目其实还会经常执行些诸如 测试、格式检查、库下载安装等命令。

当然你也可以编写 shell 脚本来执行这些命令，进一步进行了简化。


其实有更好的选择，即 Makefile。 在诸多的开源项目中经常能看到 Makefile 的身影。当你的项目中文件发生变化，都可以使用 Makefile 执行命令来自动构建

### 2. Makefile 语法


```
PROJECT="example"

default:
    echo ${PROJECT}

install:
    @govendor sync -v
    
test: install
    @go test ./...

.PHONY: default install test


```

上文是一个非常简单的 Makefile 文件，通过这些命令的编写，直接执行 `make` , `make install`, `make test` 等就能完成对应的命令。


格式介绍：

```
<target> : <prerequisites> 
[tab]  <commands> 
```

- target : 即自定义的想要执行的命令
- prerequisites: 前置条件，即执行 target 命令之前执行的命令
- commands : 具体的执行的命令
- .PHONY 伪指令，内置的关键字
- 不带参数，默认执行第一个 target
- @ 表示禁止回声，即终端不会打印真实的执行命令
- `#` 表示注释
- ${val} 表示变量，和 shell 脚本中的变量的声明和使用一致
- 允许使用 通配符

### 3. Go 项目

Go 中支持内置的 `go` 命令，可以用来执行：测试、编译、运行、语法检查等命令

一个完善的 Go 项目经常会执行哪些命令？

- go vet 静态检查
- go test 运行单元测试
- go fmt 格式化
- go build 编译
- go run 运行
...


所以一个适用于 Go 项目的 Makefile 也应该支持这些命令。

- make default : 编译
- make fmt: 格式化
- make vet: 静态检查
- make test: 运行测试
- make install: 下载依赖库
- make clean: 移除编译的二进制文件

所以整体如下：

``` 
BINARY="example"
VERSION=1.0.0
BUILD=`date +%FT%T%z`

PACKAGES=`go list ./... | grep -v /vendor/`
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default:
	@go build -o ${BINARY} -tags=jsoniter

list:
	@echo ${PACKAGES}
	@echo ${VETPACKAGES}
	@echo ${GOFILES}

fmt:
	@gofmt -s -w ${GOFILES}

fmt-check:
	@diff=$$(gofmt -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

install:
	@govendor sync -v

test:
	@go test -cpu=1,2,4 -v -tags integration ./...

vet:
	@go vet $(VETPACKAGES)

docker:
    @docker build -t wuxiaoxiaoshen/example:latest .

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: default fmt fmt-check install test vet clean


```



### 4. 补充

Makefile 构建工具，大大的简化了构建项目的难度。

真实的生产环境下，需要使用到CI/CD（持续集成和持续部署）， 所以 Makefile 也通常用来和 CI 工具配合使用。


比如新合并的代码，先触发单元测试，静态检查等，在执行 CI 脚本，成功之后，再构建镜像，推送镜像到服务器上，完成持续集成和持续部署一整套流程。


Makefile 通常配合 travis 使用。

比如：

``` 
language: go
go:
  - "1.11"
  - "1.11.x"
env:
  - GO111MODULE=on
notifications:
  email:
    recipients:
      - wuxiaoshen@shu.edu.cn
    on_success: change # default: change
    on_failure: always # default: always

before_install:
  - go test -cpu=1,2,4 -v -tags integration ./...
  - go vet $(go list ./... | grep -v /vendor/)

script:
  - make fmt
  - make fmt-check
  - make vet
  - make list  
  - go test -race  ./... -coverprofile=coverage.txt -covermode=atomic


```

希望对大家有所启发。

<完>


参考：

- [Makefile](http://www.ruanyifeng.com/blog/2015/02/make.html)
- [TravisCI](https://docs.travis-ci.com/)