package models

import (
	"github.com/thewisepigeon/goo/pkg"
	"testing"
)

func TestKeys(t *testing.T) {
	key := pkg.GenerateRandomString(15)
	t.Run("Create key", func(t *testing.T) {
		newKey := &Key{
			Key: key,
		}
		err := newKey.Save()
		if err != nil {
			t.Errorf("Wanted nil but got %q", err.Error())
		}
	})
	t.Run("Get keys", func(t *testing.T) {
		keys, err := new(Key).GetKeys()
		if err != nil {
			t.Errorf("Wanted nil but got %q", err.Error())
		}
		if len(keys) == 0 {
			t.Errorf("Wanted non empty array but got %v", keys)
		}
	})
	t.Run("Check key existence", func(t *testing.T) {
		ok, err := new(Key).IsValid(key)
		if err != nil {
			t.Errorf("Wanted nil but got %q", err.Error())
		}
		if !ok {
			t.Errorf("Wanted true but got %v", ok)
		}
	})
	t.Run("Delete key", func(t *testing.T) {
		err := new(Key).DeleteKey("1")
		if err != nil {
			t.Errorf("Wanted nil but got %q", err.Error())
		}
	})
	t.Run("Check deleted key existence", func(t *testing.T) {
		ok, err := new(Key).IsValid(key)
		if err != nil {
			t.Errorf("Wanted nil but got %q", err.Error())
		}
		if ok {
			t.Errorf("Wanted false but got %v", ok)
		}
	})
}
