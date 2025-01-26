package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type Conf struct {
	MaxSnapshots    int
	DeleteSnapshots bool
	LogToStdout     bool
}

func GetConfig() Conf {
	defaultValueConf := getDefaultConfig()
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Could not get user home directory")
		fmt.Println(err)
		return defaultValueConf
	}
	path := home + "/.config/snapmate/config.ini"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Config file not found, using default config")
		return defaultValueConf
	}

	// Read config file
	inidata, err := ini.Load(path)
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

	section = inidata.Section("logging")
	config.LogToStdout = section.Key("logToStdout").MustBool(defaultValueConf.LogToStdout)

	return config
}

func getDefaultConfig() Conf {
	return Conf{
		MaxSnapshots:    5,
		DeleteSnapshots: true,
		LogToStdout:     true,
	}
}

func SeedConfig() {
	defaultValueConf := getDefaultConfig()
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Could not get user home directory")
		fmt.Println(err)
	}
	path := home + "/.config/snapmate/config.ini"

	// Check if config file already exists
	if _, err := os.Stat(path); err == nil {
		fmt.Println("Config file already exists")
		return
	}

	iniFile := ini.Empty()
	iniFile.Section("snapshots").Key("maxSnapshots").SetValue(fmt.Sprintf("%d", defaultValueConf.MaxSnapshots))
	iniFile.Section("snapshots").Key("deleteSnapshots").SetValue(fmt.Sprintf("%t", defaultValueConf.DeleteSnapshots))
	iniFile.Section("logging").Key("logToStdout").SetValue(fmt.Sprintf("%t", defaultValueConf.LogToStdout))

	err = iniFile.SaveTo(path)
	if err != nil {
		fmt.Println("Could not save config file")
		fmt.Println(err)
	}
}
