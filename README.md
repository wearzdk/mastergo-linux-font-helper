# MasterGo Linux字体插件
## 特性

适用于Linux系统的MasterGo本地字体插件。

已实现的功能：

- [x] 获取本地字体
- [x] 字体缓存
- [ ] 字由
- [ ] 字体预览

目前暂不支持字体预览，具体效果如下

<img src="/media/zskks/数据/项目库/开源项目/mastergo-font-linux/mastergo-font-linux/README.assets/image-20230605140720645.png" alt="image-20230605140720645" style="float:left;" />

## 安装

前往[Releases](https://github.com/zskzskabcd/mastergo-linux-font-helper/releases)页面，下载最新版本

解压文件，通过终端打开目录，运行 `./install.sh` 命令，即可完成 MasterGo Linux 字体插件的安装。

安装完成后，可以通过运行 `systemctl --user status mastergo-font` 命令，确认插件已经成功安装。

## 解决鼠标中键问题

在 Linux 系统下，鼠标中键通常用于快速粘贴。尽管这个功能十分便捷，但是在 MasterGo 设计软件等特定软件上，可能会出现冲突。

要解决这个问题，可以使用这个油猴脚本

## License

MIT