package main

import (
	"flag"
	"os"
	"snapmate/db"
	"snapmate/logger"
	"snapmate/snaphots"
)

var (
	help   bool
	isHook bool
)

func main() {
	l := logger.NewLogger()
	parseFlags()

	if help {
		flag.Usage()
		return
	}

	if !isHook {
		l.Error("This program should only be run as a pacman hook")
		flag.Usage()
		os.Exit(1)
	}

	err := db.Migrate()
	if err != nil {
		l.Error("Could not migrate database. Exiting")
		l.Error(err.Error())
		os.Exit(1)
	}

	err = snaphots.CreateSnapshot()
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}
}

func parseFlags() {
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&isHook, "hook", false, "Indicates if pacman ran this program as a hook")
	flag.Parse()
}
