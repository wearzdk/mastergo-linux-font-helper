#!/bin/bash

# 不允许使用root用户执行
if [ "$EUID" -eq 0 ]; then
  echo "请使用普通用户权限执行";
  exit 1;
fi

# 安装路径
INSTALL_PATH=$HOME/.local/bin

arch=$(uname -m);
if [ "$arch" == "x86_64" ]; then
  arch="amd64";
elif [ "$arch" == "i386" ]; then
  arch="386";
elif [ "$arch" == "aarch64" ]; then
  arch="arm64";
else
  echo "不受支持的CPU架构: $arch";
  exit 1;
fi


# 复制可执行文件
cp ./mastergo-font-linux-${arch} $INSTALL_PATH/mastergo-font-linux

# 设置权限
chown $USER:$USER $INSTALL_PATH/mastergo-font-linux
chmod +x $INSTALL_PATH/mastergo-font-linux

# 创建服务文件目录
mkdir -p $HOME/.config/systemd/user/

# 复制服务文件
cp ./mastergo-font.service $HOME/.config/systemd/user/mastergo-font.service

# 启动服务
systemctl --user daemon-reload
systemctl --user enable mastergo-font
systemctl --user start mastergo-font

# 提示
echo "安装完成 可通过运行 systemctl --user status mastergo-font 查看服务状态"