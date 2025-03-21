package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type Conf struct {
	MaxSnapshots    int
	DeleteSnapshots bool
	MinTimeBetween  int // Minimum time between snapshots in minutes
	DebugLog        bool
	DatabasePath    string // Path for the SQLite database
}

const Path = "/etc/snapmate/config.ini"

func GetConfig() Conf {
	defaultValueConf := getDefaultConfig()

	if _, err := os.Stat(Path); os.IsNotExist(err) {
		fmt.Println("Config file not found, using default config")
		return defaultValueConf
	}

	// Read config file
	inidata, err := ini.Load(Path)
	if err != nil {
		fmt.Println("Could not read config file")
		fmt.Println(err)
		return defaultValueConf
	}

	// Parse config file
	config := Conf{}

	section := inidata.Section("snapshots")
	config.MaxSnapshots = section.Key("maxSnapshots").MustInt(defaultValueConf.MaxSnapshots)
	config.DeleteSnapshots = section.Key("deleteSnapshots").MustBool(defaultValueConf.DeleteSnapshots)
	config.MinTimeBetween = section.Key("minTimeBetween").MustInt(defaultValueConf.MinTimeBetween)

	section = inidata.Section("logging")
	config.DebugLog = section.Key("debugLog").MustBool(defaultValueConf.DebugLog)

	section = inidata.Section("database")
	config.DatabasePath = section.Key("path").String()

	return config
}

func getDefaultConfig() Conf {
	return Conf{
		MaxSnapshots:    5,
		DeleteSnapshots: true,
		MinTimeBetween:  60,
		DebugLog:        true,
		DatabasePath:    "/home/snapmate.db",
	}
}
