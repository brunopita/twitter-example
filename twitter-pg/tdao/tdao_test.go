package tdao

import (
	"testing"

	"github.com/brunopita/twitter-example/twitter-pg/spg"
)

var tweet = Tweet{
	Id:       1,
	IdUser:   1,
	CreateAt: "Tue May 12 23:45:05 +0000 2020",
	Hashtag:  "teste",
	Message:  "alow",
}
var user = User{
	Name:      "Test",
	Followers: 15,
	Locate:    "Brazil",
	Id:        1,
}

func TestInsertUser(t *testing.T) {
	db, err := spg.GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = InsertUser(user, db)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteUser(t *testing.T) {
	db, err := spg.GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = DeleteUser(user.Id, db)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertTweet(t *testing.T) {
	db, err := spg.GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = InsertTweet(tweet, db)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTweet(t *testing.T) {
	db, err := spg.GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = DeleteTweet(tweet.Id, db)
	if err != nil {
		t.Error(err)
	}
}
func TestTopFiveFollowers(t *testing.T) {
	db, err := spg.GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	_, err = GetTopFiveUserFollowers(db)
	if err != nil {
		t.Error(err)
	}
}

func TestGetQttyForHourByHashtag(t *testing.T) {
	db, err := spg.GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	_, err = GetQttyForHourByHashtag(db)
	if err != nil {
		t.Error(err)
	}
}
