package spg

import (
	"testing"
)

func TestSpg(t *testing.T) {
	db, err := GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Error(err)
	}
}
