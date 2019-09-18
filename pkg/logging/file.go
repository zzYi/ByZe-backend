package logging

import (
	"ByZe/pkg/setting"
	"fmt"
	"time"
)

// GetLogFilePath get the log file save path
func GetLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

// GetLogFileName get the save name of the log file
func GetLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}
