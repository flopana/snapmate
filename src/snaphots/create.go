package snaphots

import (
	"os"
	"os/exec"
	"snapmate/db"
	"snapmate/logger"
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

	err = deleteOldestSnapshots()
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
	l.Info("Snapshot comment: ", pacmanArgs)
	l.Debug("Inserting snapshot into database")
	snapshot, err := db.CreateSnapshot(snapshotName, pacmanArgs)
	if err != nil {
		return err
	}
	l.Debug("Snapshot inserted into database with ID: ", snapshot.ID)

	return nil
}
