package controller

import "testing"

func TestController(t *testing.T) {
	result := TopFiveFollowController()
	if result == nil {
		t.Error("Nullo")
	}
} 