package errorHandler

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Message map[string]string

func LoadLocale() error {
	Message = make(map[string]string)
	dir := "./locales"
	if _, err := os.Stat(dir); err != nil {
		dir = "../../locales"
	}
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			lang := info.Name()
			file, err := ioutil.ReadFile(dir + "/" + lang)
			if err != nil {
				return err
			}
			data := make(map[string]string)
			err = json.Unmarshal(file, &data)
			if err != nil {
				return err
			}
			Message = data
			return nil
		})
	return err
}

func GetMessage(code string) string {
	if s, ok := Message[code]; ok {
		return s
	}
	return "unknown errCode: " + code
}
