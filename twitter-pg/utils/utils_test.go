package utils

import "testing"

func TestStringToTimestampPostgres(t *testing.T) {
	var time = "Tue May 12 23:45:05 +0000 2020"
	_, err := StringToTimestampPostgres(time)
	if err != nil {
		t.Error(err)
	}
}
