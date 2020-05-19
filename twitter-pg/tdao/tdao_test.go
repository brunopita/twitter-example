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

	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}

	err = InsertUser(user, tx)
	if err != nil {
		t.Error(err)
		tx.Rollback()
	}
	tx.Commit()
}

func TestDeleteUser(t *testing.T) {
	db, err := spg.GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}
	err = DeleteUser(user.Id, tx)
	if err != nil {
		tx.Rollback()
		t.Error(err)
	}
	tx.Commit()
}

func TestInsertTweet(t *testing.T) {
	db, err := spg.GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}

	err = InsertTweet(tweet, tx)
	if err != nil {
		t.Error(err)
		tx.Rollback()
	}
	tx.Commit()
}

func TestDeleteTweet(t *testing.T) {
	db, err := spg.GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}

	err = DeleteTweet(tweet.Id, tx)
	if err != nil {
		t.Error(err)
		tx.Rollback()
	}
	tx.Commit()
}
func TestTopFiveFollowers(t *testing.T) {
	db, err := spg.GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}

	_, err = GetTopFiveUserFollowers(tx)
	if err != nil {
		t.Error(err)
		tx.Rollback()
	}
	tx.Commit()
}

func TestGetQttyForHourByHashtag(t *testing.T) {
	db, err := spg.GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}

	_, err = GetQttyForHourByHashtag(tx)
	if err != nil {
		t.Error(err)
		tx.Rollback()
	}
	tx.Commit()
}

func TestGetQttyPostForHashtagByLocate(t *testing.T) {
	db, err := spg.GetConnection("192.168.0.153", "15432", "twitter", "teste@123", "twitter")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}

	_, err = GetQttyPostForHashtagByLocate(tx)
	if err != nil {
		t.Error(err)
		tx.Rollback()
	}
	tx.Commit()
}
