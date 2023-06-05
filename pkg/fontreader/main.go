package fontreader

import "os"

func GetFontInfo(path string) (*Font, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return GetSfntFontInfo(file)
}
