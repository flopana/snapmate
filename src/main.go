package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	help       bool
	isHook     bool
	seedConfig bool
)

func main() {
	l := newLogger(true)
	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&isHook, "hook", false, "Indicates if pacman ran this program as a hook")
	flag.BoolVar(&seedConfig, "seed-config", false, "Seed the config file with default values")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	if seedConfig {
		l.debug("Seeding config file")
		SeedConfig()
		return
	}

	if !isHook {
		l.error("This program should only be run as a pacman hook or with the -seed-config flag")
		flag.Usage()
		return
	}

	_ = GetConfig()

	ppid := os.Getppid()
	l.debug("Parent PID: ", ppid)
	l.debug("Parent CMD: ", getProcessArgs(ppid))
}

func getProcessArgs(pid int) string {
	cmd := exec.Command("ps", "-p", fmt.Sprintf("%d", pid), "-o", "args=")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}
