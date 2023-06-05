#!/bin/bash

# 需要在hack目录下执行
# 检查GOOS GOARCH 是否设置
if [ -z "$GOOS" ]; then
  export GOOS=$(go env GOOS)
fi

if [ -z "$GOARCH" ]; then
  export GOARCH=$(go env GOARCH)
fi

# 仅构建当前系统的可执行文件
go build -o ./mastergo-font-${GOOS}-${GOARCH} -trimpath -ldflags="-s -w" ../main.go

mkdir -p ../output

cp ../res/install.sh ./install.sh
cp ../res/mastergo-font.service ./mastergo-font.service


# 打包
tar -czf ../output/mastergo-font-${GOOS}-${GOARCH}.tar.gz ./mastergo-font-${GOOS}-${GOARCH} ./install.sh ./mastergo-font.service

rm ./install.sh
rm ./mastergo-font.service
rm ./mastergo-font-${GOOS}-${GOARCH}