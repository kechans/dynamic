package utils

import (
	"github.com/go-ini/ini"
	"log"
)

var cfg *ini.File

// Setup initialize the configuration instance
func Setup(conPath string, section string, v interface{}) {
	var err error
	cfg, err = ini.Load(conPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	log.Printf("init confPath :[%s] section :[%s] configName :[%v]", conPath, section, v)
	GetAppMap(section, v)
}

// mapTo map section
func GetAppMap(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
