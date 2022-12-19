package setting

import "time"

type App struct {
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

// Setup initialize the configuration instance
func Setup() {
	// App settings
	AppSetting.PageSize = 10
	AppSetting.PrefixUrl = "http://127.0.0.1:9999"

	AppSetting.RuntimeRootPath = "runtime/"

	AppSetting.LogSavePath = "logs/"
	AppSetting.LogSaveName = "log"
	AppSetting.LogFileExt = "log"
	AppSetting.TimeFormat = "20060102"

	// Server settings
	ServerSetting.RunMode = "debug"
	ServerSetting.HttpPort = 9999
	ServerSetting.ReadTimeout = 60
	ServerSetting.WriteTimeout = 60

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}
