package snaphots

import (
	"os"
	"os/exec"
	"snapmate/config"
	"snapmate/db"
	"snapmate/logger"
	"time"
)

func CreateSnapshot() error {
	conf := config.GetConfig()

	createSnapshot, err := checkForMinimumTimeBetween(conf)
	if err != nil {
		return err
	}

	if !createSnapshot {
		err := deleteOldestSnapshots()
		if err != nil {
			return err
		}
		return nil
	}

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

func checkForMinimumTimeBetween(conf config.Conf) (bool, error) {
	newestSnapshot, err := db.GetNewestSnapshot()
	if err != nil {
		return false, err
	}

	if newestSnapshot != nil {
		now := time.Now()
		diff := now.Sub(newestSnapshot.CreatedAt)
		diffMinutes := int(diff.Minutes())
		if diffMinutes < conf.MinTimeBetween {
			l := logger.NewLogger()
			l.Info("Minimum time between snapshots not reached, not creating snapshot")
			l.Debug("Minutes since last snapshot: ", diffMinutes)
			l.Debug("Minimum time between snapshots: ", conf.MinTimeBetween)

			return false, nil
		}
	}

	return true, nil
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
