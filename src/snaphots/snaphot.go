package snaphots

import (
	"fmt"
	"os/exec"
	"snapmate/config"
	"snapmate/logger"
)

func CreateSnapshot(pacmanArgs string, ppid int, conf config.Conf) error {
	err := timeshiftCreateSnapshot(pacmanArgs)
	if err != nil {
		return err
	}

	return nil
}

func timeshiftCreateSnapshot(pacmanArgs string) error {
	out, err := exec.Command("timeshift", "--create", "--comments", pacmanArgs).Output()
	if err != nil {
		l := logger.NewLogger()
		l.Error("Could not create snapshot")
		l.Error(string(out))

		return err
	}
	fmt.Println(string(out))
	return nil
}
