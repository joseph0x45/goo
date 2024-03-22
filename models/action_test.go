package models

import (
	"database/sql"
	"github.com/thewisepigeon/goo/database"
	"testing"
)

func TestAction(t *testing.T) {
	database.ResetDB()
	t.Run("Create new action", func(t *testing.T) {
		newAction := &Action{
			Name:           "test_action",
			WorkDir:        ".",
			Command:        "echo 'test command'",
			RecoverCommand: "echo 'action failed'",
		}
		err := newAction.Save()
		if err != nil {
			t.Fatalf("Failed to create action %q", err.Error())
		}
	})
	t.Run("Get duplicate action name uniqueness", func(t *testing.T) {
		ok, err := new(Action).IsNotDuplicateName("test_action")
		if err != nil {
			t.Fatalf("Error while checking for name uniqueness: %q", err.Error())
		}
		if ok {
			t.Errorf("Wanted false but got %v", ok)
		}
	})
	t.Run("Get unique action name uniqueness", func(t *testing.T) {
		ok, err := new(Action).IsNotDuplicateName("unique_name")
		if err != nil {
			t.Fatalf("Error while checking for name uniqueness: %q", err.Error())
		}
		if !ok {
			t.Errorf("Wanted true but got %v", ok)
		}
	})
	t.Run("Create action with duplicate name", func(t *testing.T) {
		newAction := &Action{
			Name: "test_action",
		}
		err := newAction.Save()
		if err == nil {
			t.Fatal("Wanted error but got nil")
		}
	})
	t.Run("Get action by name", func(t *testing.T) {
		action, err := new(Action).GetByName("test_action")
		if err != nil {
			t.Errorf("Wanted nil but got %q", err.Error())
		}
		if action.Name != "test_action" {
			t.Errorf("Wanted 'test_action' as action name but got %q", action.Name)
		}
	})
	t.Run("List actions", func(t *testing.T) {
		_, err := new(Action).List()
		if err != nil {
			t.Errorf("Wanted nil but got %q", err.Error())
		}
	})
	t.Run("Delete action", func(t *testing.T) {
		err := new(Action).Delete("1")
		if err != nil {
			t.Errorf("Wanted nil but got %q", err)
		}
	})
	t.Run("Get deleted action by name", func(t *testing.T) {
		_, err := new(Action).GetByName("test_action")
		if err == nil {
			t.Error("Wanted error but got nil")
		}
		if err != sql.ErrNoRows {
			t.Errorf("Wanted sql.ErrNoRows but got %q", err.Error())
		}
	})
	database.ResetDB()
}
