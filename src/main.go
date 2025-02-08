package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
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

	ppid := os.Getppid()
	pacmanArgs, err := getProcessArgs(ppid)
	if err != nil {
		l.Error("Could not get pacman args")
		return
	}

	l.Debug("Parent PID: ", ppid)
	l.Debug("Parent CMD: ", pacmanArgs)

	err = snaphots.CreateSnapshot(pacmanArgs, ppid, config.GetConfig())
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

func getProcessArgs(pid int) (string, error) {
	cmd := exec.Command("ps", "-p", fmt.Sprintf("%d", pid), "-o", "args=")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
