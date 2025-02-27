package snaphots

import (
	"os/exec"
	"snapmate/config"
	"snapmate/db"
	"snapmate/logger"
	"strings"
)

func timeshiftDeleteSnapshot(snapshot *db.Snapshot) error {
	l := logger.NewLogger()
	out, err := exec.Command("timeshift", "--delete", "--snapshot", snapshot.Name).Output()
	if err != nil {
		l.Warn("Could not delete snapshot")
		out := string(out)
		l.Error(out)
		if strings.Contains(out, "Could not find snapshot") {
			l.Warn("Snapshot not found, maybe it was already deleted")
			return nil
		}

		return err
	}

	l.Info("Deleted snapshot: ", snapshot.Name)
	return nil
}

func deleteOldestSnapshots() error {
	l := logger.NewLogger()
	conf := config.GetConfig()

	if !conf.DeleteSnapshots {
		l.Info("deleteSnapshots is false, not deleting snapshots")
		return nil
	}
	l.Info("Deleting snapshots if necessary")

	snapshots, err := db.GetOldestSnapshots()
	if err != nil {
		return err
	}

	l.Debug("Number of snapshots: ", len(snapshots))
	l.Debug("MaxSnapshots: ", conf.MaxSnapshots)

	if len(snapshots) <= conf.MaxSnapshots {
		l.Info("Number of snapshots is less or equal than maxSnapshots, not deleting any")
		return nil
	}

	for i := 0; i < len(snapshots)-conf.MaxSnapshots; i++ {
		snapshot := &snapshots[i]
		err := timeshiftDeleteSnapshot(snapshot)
		if err != nil {
			return err
		}

		err = db.DeleteSnapshot(snapshot)
		if err != nil {
			return err
		}
	}

	return nil
}
