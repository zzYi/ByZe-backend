package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// App appsettings
type App struct {
	PageSize        int
	JwtSecret       string
	RuntimeRootPath string
	FileSavePath    string
	FileAllowExts   []string
	LogSavePath     string
	LogSaveName     string
	TimeFormat      string
	LogFileExt      string
}

// AppSetting appsettings
var AppSetting = &App{}

// Server serversettings
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// ServerSetting serversetting
var ServerSetting = &Server{}

// Database databasesettings
// type Database struct {
// 	User        string
// 	Password    string
// 	Host        string
// 	Type        string
// 	Name        string
// 	TablePrefix string
// }

// // DatabaseSetting databasesettings
// var DatabaseSetting = &Database{}

var cfg *ini.File

// SetUp initialize the configuration instance
func SetUp() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	// mapTo("database", DatabaseSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
