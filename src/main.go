package main

import (
	"flag"
	"fmt"
	"os"
	"snapmate/db"
	"snapmate/logger"
	"snapmate/snaphots"
)

var (
	help        bool
	isHook      bool
	versionFlag bool
)

func main() {
	l := logger.NewLogger()
	parseFlags()

	if help {
		flag.Usage()
		return
	}

	if versionFlag {
		fmt.Printf("%s version %s", name, version)
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

	skipSnapshot := false
	if os.Getenv("SKIP_SNAPSHOT") != "" {
		l.Info("Environment variable SKIP_SNAPSHOT is set, not creating snapshot")
		skipSnapshot = true
	}

	err = snaphots.CreateSnapshot(skipSnapshot)
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}
}

func parseFlags() {
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&isHook, "hook", false, "Indicates if pacman ran this program as a hook")
	flag.BoolVar(&versionFlag, "version", false, "Show version info")
	flag.Parse()
}
