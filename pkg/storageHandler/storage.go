package storageHandler

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func SaveImage(base64Str, filename string) (string, error) {
	err := deleteImageFile(filename)
	img, _ := base64.StdEncoding.DecodeString(base64Str)
	err = ioutil.WriteFile(".\\storage\\images\\"+filename+".png", img, 0666)
	return filename + ".png", err
}
func deleteImageFile(filename string) error {
	path := ".\\storage\\images\\" + filename + ".png"
	if _, err := os.Stat(path); err == nil {
		return os.Remove(path)
	} else {
		return nil
	}
}
func SaveVideo(base64Str, filename string) (string, error) {
	err := deleteVideoFile(filename)
	video, _ := base64.StdEncoding.DecodeString(base64Str)
	err = ioutil.WriteFile(".\\storage\\videos\\"+filename+".mp4", video, 0666)
	return filename + ".mp4", err
}
func deleteVideoFile(filename string) error {
	path := ".\\storage\\videos\\" + filename + ".mp4"
	if _, err := os.Stat(path); err == nil {
		return os.Remove(path)
	} else {
		return nil
	}
}
