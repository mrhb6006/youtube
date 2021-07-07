package errorHandler

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var statusMessage map[string]map[string]string

func LoadLocale() error {
	statusMessage = make(map[string]map[string]string)
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
			statusMessage[lang[0:len(lang)-5]] = data
			return nil
		})
	return err
}

func GetMessage(code, lang string) string {
	if messages, ok := statusMessage[lang]; ok {
		return messages[code]
	}
	return "unknown errCode: " + code
}
