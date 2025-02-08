package snaphots

import (
	"bufio"
	"errors"
	"os/exec"
	"snapmate/config"
	"snapmate/logger"
	"strings"
)

func CreateSnapshot(pacmanArgs string, ppid int, conf config.Conf) error {
	err := timeshiftCreateSnapshot(pacmanArgs)
	if err != nil {
		return err
	}

	return nil
}

func timeshiftCreateSnapshot(pacmanArgs string) error {
	l := logger.NewLogger()
	out, err := exec.Command("timeshift", "--create", "--comments", pacmanArgs).Output()
	if err != nil {
		l.Error("Could not create snapshot")
		l.Error(string(out))

		return err
	}
	l.Info("Snapshot created")
	snapshotName, err := parseTimeshiftOutput(string(out))
	if err != nil {
		l.Error("Could not parse timeshift output")
		l.Error(string(out))

		return err
	}
	l.Info("Created snapshot: ", snapshotName)

	return nil
}

/**
 * This function is used to parse the output of the timeshift command.
 * It looks for the line that contains the snapshot name and returns it.
 */
func parseTimeshiftOutput(out string) (string, error) {
	// Tagged snapshot '2025-02-08_12-22-58': ondemand
	scanner := bufio.NewScanner(strings.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Tagged snapshot '") {
			snapshot := strings.Split(line, "'")[1]
			return snapshot, nil
		}
	}

	return "", errors.New("could not parse timeshift output")
}
