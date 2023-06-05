package fontreader

import (
	"mastergo-font-linux/pkg/fontcn"
	"os"
	"strings"

	"golang.org/x/image/font/sfnt"
)

type Font struct {
	PostscriptName  string `json:"postscriptName"`
	Family          string `json:"family"`
	LocalizedFamily string `json:"localizedFamily"`
	Style           string `json:"style"`
	LocalizedStyle  string `json:"localizedStyle"`
	Weight          uint16 `json:"weight"`
	Stretch         uint16 `json:"stretch"`
	Italic          bool   `json:"italic"`
	Monospace       bool   `json:"monospace"`
}

// 根据Style的值获取Weight的值
func GetWeight(style string) uint16 {
	switch style {
	case "Thin":
		return 100
	case "ExtraLight":
		return 200
	case "Light":
		return 300
	case "Regular":
		return 400
	case "Medium":
		return 500
	case "SemiBold":
		return 600
	case "Bold":
		return 700
	case "ExtraBold":
		return 800
	case "Black":
		return 900
	default:
		return 400
	}
}

func GetSfntFontInfo(file *os.File) (*Font, error) {
	font, err := sfnt.ParseReaderAt(file)
	if err != nil {
		return nil, err
	}
	postScriptName, _ := font.Name(nil, sfnt.NameIDPostScript)
	family, _ := font.Name(nil, sfnt.NameIDTypographicFamily)
	if family == "" {
		family, _ = font.Name(nil, sfnt.NameIDFamily)
	}
	style, _ := font.Name(nil, sfnt.NameIDTypographicSubfamily)
	if style == "" {
		style, _ = font.Name(nil, sfnt.NameIDSubfamily)
	}

	isItalic := strings.Contains(strings.ToLower(style), "italic")

	isMonospace := font.PostTable().IsFixedPitch

	fontInfo := &Font{
		PostscriptName:  postScriptName,
		Family:          family,
		LocalizedFamily: fontcn.PraseCNFamily(family),
		Style:           style,
		LocalizedStyle:  fontcn.PraseCNStyle(style),
		Weight:          GetWeight(style),
		Stretch:         5,
		Italic:          isItalic,
		Monospace:       isMonospace,
	}
	return fontInfo, err
}
