package models

import (
	"testing"
)

func TestModels(t *testing.T) {
	t.Run("Create log", func(t *testing.T) {
		newLog := Log{
			Action: 1,
		}
		err := newLog.Save()
		if err != nil {
			t.Errorf("Wanted nil but got %q", err.Error())
		}
	})
	t.Run("Get logs", func(t *testing.T) {
		logs, err := new(Log).List()
		if err != nil {
			t.Errorf("Wanted nil but got %q", err.Error())
		}
		if len(logs) == 0 {
			t.Errorf("Wanted non empty array but got %v", logs)
		}
	})
}
