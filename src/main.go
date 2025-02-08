package main

import (
	"flag"
	"snapmate/config"
	"snapmate/logger"
	"snapmate/snaphots"
)

var (
	help       bool
	isHook     bool
	seedConfig bool
)

func main() {
	l := logger.NewLogger()
	parseFlags()

	if help {
		flag.Usage()
		return
	}

	if seedConfig {
		l.Info("Seeding config file")
		err := config.SeedConfig()
		if err != nil {
			l.Error(err.Error())
			return
		}
		l.Info("Config file seeded")
		return
	}

	if !isHook {
		l.Error("This program should only be run as a pacman hook or with the -seed-config flag")
		flag.Usage()
		return
	}

	err := snaphots.CreateSnapshot()
	if err != nil {
		l.Error(err.Error())
		return
	}
}

func parseFlags() {
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&isHook, "hook", false, "Indicates if pacman ran this program as a hook")
	flag.BoolVar(&seedConfig, "seed-config", false, "Seed the config file with default values")
	flag.Parse()
}
