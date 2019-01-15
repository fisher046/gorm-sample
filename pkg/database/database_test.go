package database

import (
	"testing"
)

func TestGetDB(t *testing.T) {
	db1 := GetDB()
	if db1 == nil {
		t.Fatal("should not be nil")
	}

	db2 := GetDB()
	if db2 != db1 {
		t.Fatalf("expected: %v, actual: %v", db1, db2)
	}
}
