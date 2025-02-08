package config

import (
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type Conf struct {
	MaxSnapshots    int
	DeleteSnapshots bool
	AskUser         bool // Ask user if a snapshot should be taken
	MinTimeBetween  int  // Minimum time between snapshots in minutes
	DebugLog        bool
}

const ConfigPath = "/etc/snapmate/config.ini"

func GetConfig() Conf {
	defaultValueConf := getDefaultConfig()

	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		fmt.Println("Config file not found, using default config")
		return defaultValueConf
	}

	// Read config file
	inidata, err := ini.Load(ConfigPath)
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
	config.AskUser = section.Key("askUser").MustBool(defaultValueConf.AskUser)
	config.MinTimeBetween = section.Key("minTimeBetween").MustInt(defaultValueConf.MinTimeBetween)

	section = inidata.Section("logging")
	config.DebugLog = section.Key("debugLog").MustBool(defaultValueConf.DebugLog)

	return config
}

func getDefaultConfig() Conf {
	return Conf{
		MaxSnapshots:    5,
		DeleteSnapshots: true,
		AskUser:         false,
		MinTimeBetween:  60,
		DebugLog:        true,
	}
}

func SeedConfig() error {
	defaultValueConf := getDefaultConfig()

	// Check if config file already exists
	if _, err := os.Stat(ConfigPath); err == nil {
		return errors.New("config file already exists")
	}

	iniFile := ini.Empty()
	iniFile.Section("snapshots").Key("maxSnapshots").SetValue(fmt.Sprintf("%d", defaultValueConf.MaxSnapshots))
	iniFile.Section("snapshots").Key("deleteSnapshots").SetValue(fmt.Sprintf("%t", defaultValueConf.DeleteSnapshots))
	iniFile.Section("snapshots").Key("askUser").SetValue(fmt.Sprintf("%t", defaultValueConf.AskUser))
	iniFile.Section("snapshots").Key("minTimeBetween").SetValue(fmt.Sprintf("%d", defaultValueConf.MinTimeBetween))

	iniFile.Section("logging").Key("debugLog").SetValue(fmt.Sprintf("%t", defaultValueConf.DebugLog))

	err := iniFile.SaveTo(ConfigPath)
	if err != nil {
		return err
	}

	return nil
}
