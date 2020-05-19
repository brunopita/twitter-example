package controller

import (
	"context"
	"testing"
)

func TestTopFiveeController(t *testing.T) {
	_, err := TopFiveFollowController(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestPostsForHourController(t *testing.T) {
	_, err := PostsForHourController(context.Background())
	if err != nil {
		t.Error(err)
	}
}
