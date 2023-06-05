package dao

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mastergo-font-linux/internal/pkg/config"
	"mastergo-font-linux/internal/pkg/resp"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type CacheFontItem struct {
	Key        string    `json:"key"`
	FileType   string    `json:"fileType"`
	Size       int64     `json:"size"`
	FilePath   string    `json:"filePath"`
	CreateTime int64     `json:"createTime"`
	CreateDate time.Time `json:"createDate"`
	Count      int       `json:"count"`
}

var cachedFonts []CacheFontItem
var cachedFontsMap map[string]*CacheFontItem

func init() {
	cachedFonts = make([]CacheFontItem, 0)
	// 读入配置
	path := filepath.Join(config.GetFontCachePath(), "cache.json")
	err := readJson(path, &cachedFonts)
	if err != nil {
		panic(err)
	}
	// 转换为map
	cachedFontsMap = make(map[string]*CacheFontItem)
	for i := 0; i < len(cachedFonts); i++ {
		cachedFontsMap[cachedFonts[i].FilePath] = &cachedFonts[i]
	}
	// 检查缓存
	ReloadCache()
}

// 重新检查缓存
func ReloadCache() {
	// 检查缓存文件是否依然存在
	for _, item := range cachedFonts {
		if _, err := os.Stat(item.FilePath); os.IsNotExist(err) {
			// 不存在则删除
			delete(cachedFontsMap, item.FilePath)
		}
	}
	SaveCacheJson()
}

func readJson(path string, cacheFontItem *[]CacheFontItem) error {
	// 文件不存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("cache.json not exist")
		err := writeJson(path, cacheFontItem)
		if err != nil {
			return err
		}
	}
	// 读取文件
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	fileByte, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	// 解析JSON
	err = json.Unmarshal(fileByte, cacheFontItem)
	if err != nil {
		return err
	}
	return nil
}

func writeJson(path string, cacheFontItem *[]CacheFontItem) error {
	// 创建文件
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	// 序列化JSON
	jsonByte, err := json.Marshal(cacheFontItem)
	if err != nil {
		return err
	}
	// 写入文件
	_, err = file.Write(jsonByte)
	if err != nil {
		return err
	}
	return nil
}

func SaveCacheJson() {
	// 重新生成cachedFonts
	cachedFonts = make([]CacheFontItem, 0)
	for _, item := range cachedFontsMap {
		cachedFonts = append(cachedFonts, *item)
	}
	path := filepath.Join(config.GetFontCachePath(), "cache.json")
	err := writeJson(path, &cachedFonts)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 获取缓存字体列表 /cache-fonts
func GetCacheFontsHandler(w http.ResponseWriter, r *http.Request) {
	// 返回字体列表
	jsonByte, err := json.Marshal(cachedFonts)
	if err != nil {
		resp.HttpError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonByte)
}

// upload-font
func UploadFontHandler(w http.ResponseWriter, r *http.Request) {
	// 读取文件
	file, header, err := r.FormFile("font")
	if err != nil {
		log.Printf("upload file error: %v", err)
		resp.HttpError(w, err)
		return
	}
	defer file.Close()
	// 获取key
	key := r.FormValue("key")
	// 计算hash
	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		log.Printf("upload file error: %v", err)
		resp.HttpError(w, err)
		return
	}
	hashByte := hash.Sum(nil)
	hashStr := hex.EncodeToString(hashByte)
	// offset归零
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		log.Println(err)
		resp.HttpError(w, err)
		return
	}
	// 创建文件
	path := filepath.Join(config.GetFontCachePath(), hashStr)
	newFile, err := os.Create(path)
	if err != nil {
		log.Println(err)
		resp.HttpError(w, err)
		return
	}
	defer newFile.Close()
	// 写入文件
	_, err = io.Copy(newFile, file)
	if err != nil {
		log.Println(err)
		resp.HttpError(w, err)
		return
	}
	// 添加到缓存
	cacheFontItem := CacheFontItem{
		Key:        key,
		FileType:   header.Header.Get("Content-Type"),
		Size:       header.Size,
		FilePath:   path,
		CreateTime: time.Now().Unix(),
		CreateDate: time.Now(),
		Count:      1,
	}
	cachedFonts = append(cachedFonts, cacheFontItem)
	cachedFontsMap[path] = &cacheFontItem
	// 保存缓存
	SaveCacheJson()
	// 返回成功
	w.Write([]byte("success"))
}

// 获取缓存字体 /cache-font
func GetCacheFontHandler(w http.ResponseWriter, r *http.Request) {
	// query参数 path
	path := r.URL.Query().Get("path")
	if path == "" {
		resp.HttpError(w, fmt.Errorf("path is empty"))
		return
	}
	// 获取缓存
	cacheFontItem, ok := cachedFontsMap[path]
	if !ok {
		resp.HttpError(w, fmt.Errorf("path not found"))
		return
	}
	// 检查文件是否存在
	if _, err := os.Stat(cacheFontItem.FilePath); os.IsNotExist(err) {
		// 重新检查缓存
		ReloadCache()
		resp.HttpError(w, fmt.Errorf("file not exist"))
		return
	}
	// 打开文件
	file, err := os.Open(cacheFontItem.FilePath)
	if err != nil {
		resp.HttpError(w, err)
		return
	}
	defer file.Close()
	// 返回文件
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.FormatInt(cacheFontItem.Size, 10))
	w.Header().Set("Cache-Control", "max-age=31536000")
	_, err = io.Copy(w, file)
	if err != nil {
		resp.HttpError(w, err)
		return
	}
}
