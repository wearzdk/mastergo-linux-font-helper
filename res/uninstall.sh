#!/bin/bash

# 不允许使用root用户执行
if [ "$EUID" -eq 0 ]; then
  echo "请使用普通用户权限执行";
  exit 1;
fi

# 安装路径
INSTALL_PATH=$HOME/.local/bin

# 停止服务
systemctl --user stop mastergo-font
systemctl --user disable mastergo-font

# 删除服务文件
rm $HOME/.config/systemd/user/mastergo-font.service

# 删除可执行文件
rm $INSTALL_PATH/mastergo-font-linux

systemctl --user daemon-reload

# 提示
echo "卸载完成"