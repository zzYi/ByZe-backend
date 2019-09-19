package upload

import (
	"ByZe/pkg/file"
	"ByZe/pkg/setting"
	"ByZe/pkg/util"
	"fmt"
	"os"
	"strings"
)

func GetFileName(name string) string {
	ext := file.GetExt(name)
	filename := strings.TrimSuffix(name, ext)
	filename = util.EncodeMD5(filename)

	return filename + ext
}

func getFilePath() string {
	return setting.AppSetting.FileSavePath
}

func GetFileFullPath() string {
	return setting.AppSetting.RuntimeRootPath+getFilePath()
}

func CheckFileExt(name string) bool {
	ext := file.GetExt(name)

	for _, allowExt := range setting.AppSetting.FileAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckFile(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}