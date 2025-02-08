package snaphots

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"snapmate/logger"
	"strings"
)

func CreateSnapshot() error {
	pacmanArgs, err := getProcessArgs(os.Getppid())
	if err != nil {
		return err
	}

	err = timeshiftCreateSnapshot(pacmanArgs)
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

func getProcessArgs(pid int) (string, error) {
	l := logger.NewLogger()
	cmd := exec.Command("ps", "-p", fmt.Sprintf("%d", pid), "-o", "args=")
	out, err := cmd.Output()
	if err != nil {
		l.Error("Could not get process args")
		return "", err
	}
	return string(out), nil
}
