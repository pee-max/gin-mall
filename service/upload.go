package service

import (
	"gin_mall/conf"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarToLocalStatic(file multipart.File, uId uint, userName string) (string, error) {
	bId := strconv.Itoa(int(uId))
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExists(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "user" + bId + "/" + userName + ".jpg", nil
}

func DirExists(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CreateDir(path string) bool {
	err := os.MkdirAll(path, 755)
	if err != nil {
		return false
	}
	return true
}
