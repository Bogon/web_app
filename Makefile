.PHONY: all build buildmac run runmac gotool clean help

BINARY="webapp.io"

all: gotool build

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/${BINARY}

buildmac:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 sudo  go build  -ldflags "-s -w" -o ./bin/${BINARY} -buildvcs=false

run:
	@go run ./ ./conf/config.yaml

runmac:
	@sudo go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件(linux环境下)"
	@echo "make buildmac - 编译 Go 代码, 生成二进制文件(mac 环境下)"
	@echo "make run - 直接运行 Go 代码(linux环境下)"
	@echo "make runmac - 直接运行 Go 代码(mac环境下)"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
