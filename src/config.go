package main

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
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Could not get user home directory")
		return err
	}
	path := home + "/.config/snapmate/config.ini"

	// Check if config file already exists
	if _, err := os.Stat(path); err == nil {
		return errors.New("config file already exists")
	}

	iniFile := ini.Empty()
	iniFile.Section("snapshots").Key("maxSnapshots").SetValue(fmt.Sprintf("%d", defaultValueConf.MaxSnapshots))
	iniFile.Section("snapshots").Key("deleteSnapshots").SetValue(fmt.Sprintf("%t", defaultValueConf.DeleteSnapshots))
	iniFile.Section("snapshots").Key("askUser").SetValue(fmt.Sprintf("%t", defaultValueConf.AskUser))
	iniFile.Section("snapshots").Key("minTimeBetween").SetValue(fmt.Sprintf("%d", defaultValueConf.MinTimeBetween))
	iniFile.Section("logging").Key("debugLog").SetValue(fmt.Sprintf("%t", defaultValueConf.DebugLog))

	err = iniFile.SaveTo(path)
	if err != nil {
		return err
	}

	return nil
}
