package fontcn

import "strings"

// 翻译字体 family subfamily 等

var CNFamilyRawMap = map[string]string{
	"SimSun":                "宋体",
	"SimHei":                "黑体",
	"Microsoft Yahei":       "微软雅黑",
	"Microsoft JhengHei":    "微软正黑体",
	"KaiTi":                 "楷体",
	"NSimSun":               "新宋体",
	"FangSong":              "仿宋",
	"PingFang SC":           "苹方",
	"PingFangSC-Light":      "苹方-细体",
	"PingFangSC-Ultralight": "苹方-极细体",
	"PingFangSC-Semibold":   "苹方-中粗体",
	"PingFangSC-Medium":     "苹方-中黑体",
	"PingFangSC-Regular":    "苹方-常规体",
	"PingFangSC-Thin":       "苹方-纤细体",
	"STHeiti":               "华文黑体",
	"STKaiti":               "华文楷体",
	"STSong":                "华文宋体",
	"STFangsong":            "华文仿宋",
	"STZhongsong":           "华文中宋",
	"STHupo":                "华文琥珀",
	"STXinwei":              "华文新魏",
	"STLiti":                "华文隶书",
	"STXingkai":             "华文行楷",
	"Hiragino Sans GB":      "冬青黑体简",
	"Lantinghei SC":         "兰亭黑-简",
	"Hanzipen SC":           "翩翩体-简",
	"Hannotate SC":          "手札体-简",
	"Songti SC":             "宋体-简",
	"Wawati SC":             "娃娃体-简",
	"Weibei SC":             "魏碑-简",
	"Xingkai SC":            "行楷-简",
	"Yapi SC":               "雅痞-简",
	"Yuanti SC":             "圆体-简",
	"YouYuan":               "幼圆",
	"LiSu":                  "隶书",
	"STXihei":               "华文细黑",
	"STCaiyun":              "华文彩云",
	"FZShuTi":               "方正舒体",
	"FZYaoti":               "方正姚体",
	"Source Han Sans CN":    "思源黑体",
	"Source Han Serif SC":   "思源宋体",
	"WenQuanYi Micro Hei":   "文泉驿微米黑",
	"HYQihei 40S":           "汉仪旗黑",
	"HYQihei 50S":           "汉仪旗黑",
	"HYQihei 60S":           "汉仪旗黑",
	"HYDaSongJ":             "汉仪大宋简",
	"HYKaiti":               "汉仪楷体",
	"HYJiaShuJ":             "汉仪家书简",
	"HYPPTiJ":               "汉仪PP体简",
	"HYLeMiaoTi":            "汉仪乐喵体简",
	"HYXiaoMaiTiJ":          "汉仪小麦体",
	"HYChengXingJ":          "汉仪程行体",
	"HYHeiLiZhiTiJ":         "汉仪黑荔枝",
	"HYYaKuHeiW":            "汉仪雅酷黑W",
	"HYDaHeiJ":              "汉仪大黑简",
	"HYShangWeiShouShuW":    "汉仪尚魏手书W",
	"FZYaSongS-B-GB":        "方正粗雅宋简体",
	"FZBaoSong-Z04S":        "方正报宋简体",
	"FZCuYuan-M03S":         "方正粗圆简体",
	"FZDaBiaoSong-B06S":     "方正大标宋简体",
	"FZDaHei-B02S":          "方正大黑简体",
	"FZFangSong-Z02S":       "方正仿宋简体",
	"FZHei-B01S":            "方正黑体简体",
	"FZHuPo-M04S":           "方正琥珀简体",
	"FZKai-Z03S":            "方正楷体简体",
	"FZLiBian-S02S":         "方正隶变简体",
	"FZLiShu-S01S":          "方正隶书简体",
	"FZMeiHei-M07S":         "方正美黑简体",
	"FZShuSong-Z01S":        "方正书宋简体",
	"FZShuTi-S05S":          "方正舒体简体",
	"FZShuiZhu-M08S":        "方正水柱简体",
	"FZSongHei-B07S":        "方正宋黑简体",
	"FZSong":                "方正宋三简体",
	"FZWeiBei-S03S":         "方正魏碑简体",
	"FZXiDengXian-Z06S":     "方正细等线简体",
	"FZXiHei I-Z08S":        "方正细黑一简体",
	"FZXiYuan-M01S":         "方正细圆简体",
	"FZXiaoBiaoSong-B05S":   "方正小标宋简体",
	"FZXingKai-S04S":        "方正行楷简体",
	"FZYaoTi-M06S":          "方正姚体简体",
	"FZZhongDengXian-Z07S":  "方正中等线简体",
	"FZZhunYuan-M02S":       "方正准圆简体",
	"FZZongYi-M05S":         "方正综艺简体",
	"FZCaiYun-M09S":         "方正彩云简体",
	"FZLiShu II-S06S":       "方正隶二简体",
	"FZKangTi-S07S":         "方正康体简体",
	"FZChaoCuHei-M10S":      "方正超粗黑简体",
	"FZNew BaoSong-Z12S":    "方正新报宋简体",
	"FZNew ShuTi-S08S":      "方正新舒体简体",
	"FZHuangCao-S09S":       "方正黄草简体",
	"FZShaoEr-M11S":         "方正少儿简体",
	"FZZhiYi-M12S":          "方正稚艺简体",
	"FZXiShanHu-M13S":       "方正细珊瑚简体",
	"FZCuSong-B09S":         "方正粗宋简体",
	"FZPingHe-S11S":         "方正平和简体",
	"FZHuaLi-M14S":          "方正华隶简体",
	"FZShouJinShu-S10S":     "方正瘦金书简体",
	"FZXiQian-M15S":         "方正细倩简体",
	"FZZhongQian-M16S":      "方正中倩简体",
	"FZCuQian-M17S":         "方正粗倩简体",
	"FZPangWa-M18S":         "方正胖娃简体",
	"FZSongYi-Z13S":         "方正宋一简体",
	"FZJianZhi-M23S":        "方正剪纸简体",
	"FZLiuXingTi-M26S":      "方正流行体简体",
	"FZXiangLi-S17S":        "方正祥隶简体",
	"FZCuHuoYi-M25S":        "方正粗活意简体",
	"FZPangTouYu-M24S":      "方正胖头鱼简体",
	"FZKaTong-M19S":         "方正卡通简体",
	"FZYiHei-M20S":          "方正艺黑简体",
	"FZShuiHei-M21S":        "方正水黑简体",
	"FZGuLi-S12S":           "方正古隶简体",
	"FZYouXian-Z09S":        "方正幼线简体",
	"FZQiTi-S14S":           "方正启体简体",
	"FZXiaoZhuanTi-S13T":    "方正小篆体",
	"FZYingBiKaiShu-S15S":   "方正硬笔楷书简体",
	"FZZhanBiHei-M22S":      "方正毡笔黑简体",
	"FZYingBiXingShu-S16S":  "方正硬笔行书简体",
}

var CNFamilyMap map[string]string

func init() {
	// 左侧英文小写对照
	CNFamilyMap = make(map[string]string)
	for k, v := range CNFamilyRawMap {
		CNFamilyMap[strings.ToLower(k)] = v
	}

}

func PraseCNFamily(family string) string {
	familyLower := strings.ToLower(family)
	if v, ok := CNFamilyMap[familyLower]; ok {
		return v
	}
	return family
}

var CNStyleMap = map[string]string{
	"UltraLight":  "极细体",
	"ExtraLight":  "极细体",
	"Thin":        "细体",
	"Light":       "细体",
	"Regular":     "常规",
	"Bold":        "粗体",
	"Italic":      "斜体",
	"Bold Italic": "粗斜体",
	"Medium":      "中黑体",
	"Black":       "黑体",
	"Heavy":       "黑体",
	"ExtraBlack":  "极黑体",
	"UltraBlack":  "极黑体",
}

func PraseCNStyle(style string) string {
	if v, ok := CNStyleMap[style]; ok {
		return v
	}
	return style
}
