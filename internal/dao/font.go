package dao

import (
	"encoding/json"
	"fmt"
	"io"
	"mastergo-font-linux/internal/pkg/resp"
	"mastergo-font-linux/pkg/fontreader"
	"net/http"
	"os"

	"github.com/flopp/go-findfont"
)

type FontItem struct {
	fontreader.Font
	Path               string  `json:"path"`
	UsedPostScriptName string  `json:"postscriptName"`
	CoverUrl           *string `json:"coverUrl"`
}

var LocalFontsMap = make(map[string]*FontItem)

func init() {
	// 初始化字体记录
	fonts := findfont.List()
	for _, font := range fonts {
		fontInfo, err := fontreader.GetFontInfo(font)
		if err != nil {
			// 跳过
			continue
		}
		fontItem := FontItem{
			Font:               *fontInfo,
			Path:               font,
			UsedPostScriptName: fontInfo.PostscriptName,
			CoverUrl:           nil,
		}
		LocalFontsMap[font] = &fontItem
	}
}

// local-fonts
func GetLocalFontsHandler(w http.ResponseWriter, r *http.Request) {
	fonts := findfont.List()
	fontItems := make([]FontItem, 0)
	for _, font := range fonts {
		if LocalFontsMap[font] != nil {
			fontItems = append(fontItems, *LocalFontsMap[font])
		} else {
			fontInfo, err := fontreader.GetFontInfo(font)
			if err != nil {
				// 跳过
				continue
			}
			fontItem := FontItem{
				Font:               *fontInfo,
				Path:               font,
				UsedPostScriptName: fontInfo.PostscriptName,
				CoverUrl:           nil,
			}
			LocalFontsMap[font] = &fontItem
			fontItems = append(fontItems, fontItem)
		}
	}
	jsonByte, err := json.Marshal(fontItems)
	if err != nil {
		resp.HttpError(w, err)
		return
	}
	// JSON类型
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonByte)
}

// font-file
func GetFontFileHandler(w http.ResponseWriter, r *http.Request) {
	// query参数 path
	path := r.URL.Query().Get("path")
	if path == "" {
		resp.HttpError(w, fmt.Errorf("path is empty"))
		return
	}
	// 是否为缓存字体
	if _, ok := cachedFontsMap[path]; ok {
		GetCacheFontHandler(w, r)
		return
	}
	// 安全检查
	if _, ok := LocalFontsMap[path]; !ok {
		resp.HttpError(w, fmt.Errorf("path is not in local-fonts"))
		return
	}
	// 读取字体文件
	fontFile, err := os.Open(path)
	if err != nil {
		resp.HttpError(w, err)
		return
	}
	defer fontFile.Close()
	fontByte, err := io.ReadAll(fontFile)
	if err != nil {
		resp.HttpError(w, err)
		return
	}
	// 返回字体文件
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fontByte)
}

// ziyou-fonts
func GetZiYouFontsHandler(w http.ResponseWriter, r *http.Request) {
	// 返回空数组
	w.Header().Set("Content-Type", "application/json")
	res := resp.H{
		"data": resp.H{
			"fonts": []string{},
		},
	}
	jsonByte, err := json.Marshal(res)
	if err != nil {
		resp.HttpError(w, err)
		return
	}
	w.Write(jsonByte)
}
