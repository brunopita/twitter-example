package tdao

import (
	"database/sql"
	"log"

	"github.com/brunopita/twitter-example/twitter-pg/utils"
	"github.com/dghubble/go-twitter/twitter"
)

type Tweet struct {
	Id       int64
	Message  string
	IdUser   int64
	CreateAt string
	Hashtag  string
}

type User struct {
	Id        int64
	Name      string
	Followers int
	Locate    string
}

type QttyHourHashtag struct {
	Hour    string
	Qtty    int
	Hashtag string
}

func InsertTweet(tweet Tweet, db *sql.DB) error {
	var query = "INSERT INTO tb_tweet values ($1, $2, $3, $4, $5) ON CONFLICT (id) DO NOTHING"
	createAt, err := utils.StringToTimestampPostgres(tweet.CreateAt)
	if err != nil {
		return err
	}
	_, err = db.Exec(query, tweet.Id, tweet.Message, tweet.IdUser, createAt, tweet.Hashtag)
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(user User, db *sql.DB) error {
	var query = "INSERT INTO tb_user values ($1, $2, $3, $4) ON CONFLICT (id) DO NOTHING"
	_, err := db.Exec(query, user.Id, user.Name, user.Followers, user.Locate)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int64, db *sql.DB) error {
	var query = "DELETE FROM tb_user where id = $1"
	db.QueryRow(query, id)
	return nil
}

func DeleteTweet(id int64, db *sql.DB) error {
	var query = "DELETE FROM tb_tweet where id = $1"
	db.QueryRow(query, id)
	return nil
}

func GetTopFiveUserFollowers(db *sql.DB) ([]User, error) {
	var result []User
	var query = "SELECT name, followers, locate FROM tb_user ORDER BY followers DESC limit 5"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Name, &u.Followers, &u.Locate)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

func GetQttyForHourByHashtag(db *sql.DB) ([]QttyHourHashtag, error) {
	var result []QttyHourHashtag
	var query = "select  hashtag, extract(hour from createat), count(1) from tb_tweet group by 2, hashtag order by 2,1;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var h QttyHourHashtag
		err := rows.Scan(&h.Hashtag, &h.Hour, &h.Qtty)
		if err != nil {
			return nil, err
		}
		result = append(result, h)
		log.Printf("%v", h)
	}
	return result, nil
}

func BuildUser(val *twitter.User) User {
	var user User
	user.Name = val.Name
	user.Id = val.ID
	user.Followers = val.FollowersCount
	user.Locate = val.Location
	return user
}

func BuildTweet(val *twitter.Tweet, hashtag string) Tweet {
	var tweet Tweet
	tweet.Id = val.ID
	tweet.CreateAt = val.CreatedAt
	tweet.Hashtag = hashtag
	tweet.IdUser = val.User.ID
	return tweet
}
