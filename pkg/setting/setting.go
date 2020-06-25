package setting

import (
	"dynamic/pkg/utils"
	"time"
)

const (
	AppNameConfPath = "conf/app.ini"
	RedisConfPath   = "conf/redis.ini"
	AppConfName     = "app"
	ServerConfName  = "server"
	RedisConfName   = "redis"
)

type APP struct {
	PageSize        int
	JwtSecret       int
	PrefixUrl       string
	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &APP{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

func InitConfig() {
	utils.Setup(AppNameConfPath, AppConfName, AppSetting)
	utils.Setup(AppNameConfPath, ServerConfName, ServerSetting)
	utils.Setup(RedisConfPath, RedisConfName, RedisSetting)
}
