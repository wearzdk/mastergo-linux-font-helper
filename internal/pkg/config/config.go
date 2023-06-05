package config

import (
	"os"
	"path/filepath"
)

func GetAppPath() string {
	var path string
	configPath, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	path = filepath.Join(configPath, "mastergo-font-linux")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return path
}

// 字体缓存目录
func GetFontCachePath() string {
	return filepath.Join(GetAppPath(), "font-cache")
}

func init() {
	// 初始化字体缓存目录
	path := GetFontCachePath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
