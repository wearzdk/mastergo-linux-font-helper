#!/bin/bash

# 需要在hack目录下执行

# amd64
GOOS=linux GOARCH=amd64 go build -o ./mastergo-font-linux-amd64 -trimpath -ldflags="-s -w" ../main.go
# arm64
GOOS=linux GOARCH=arm64 go build -o ./mastergo-font-linux-arm64 -trimpath -ldflags="-s -w" ../main.go
# i386
GOOS=linux GOARCH=386 go build -o ./mastergo-font-linux-386 -trimpath -ldflags="-s -w" ../main.go

mkdir -p ../bin
mkdir -p ../output

# 安装文件
cp ../res/install.sh ../bin/install.sh
cp ../res/uninstall.sh ../bin/uninstall.sh
cp ../res/mastergo-font.service ../bin/mastergo-font.service

cp ./mastergo-font-linux-amd64 ../bin

# 打包 amd64
cd ../bin
tar -czf ../output/mastergo-font-linux-amd64.tar.gz ./mastergo-font-linux-amd64 ./install.sh ./mastergo-font.service ./uninstall.sh
cd ../hack

# 打包 arm64
cp ./mastergo-font-linux-arm64 ../bin
cd ../bin
tar -czf ../output/mastergo-font-linux-arm64.tar.gz ./mastergo-font-linux-arm64 ./install.sh ./mastergo-font.service ./uninstall.sh
cd ../hack

# 打包 i386
cp ./mastergo-font-linux-386 ../bin
cd ../bin
tar -czf ../output/mastergo-font-linux-386.tar.gz ./mastergo-font-linux-386 ./install.sh ./mastergo-font.service ./uninstall.sh
cd ../hack

# 删除临时文件
rm -rf ../bin
rm mastergo-font-linux-*