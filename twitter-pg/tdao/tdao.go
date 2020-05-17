package tdao

import (
	"database/sql"

	"github.com/brunopita/twitter-example/twitter-pg/utils"
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

type HourQtty struct {
	Hour string
	Qtty int
}

func InsertTweet(tweet Tweet, db *sql.DB) error {
	var query = "INSERT INTO tb_tweet values ($1, $2, $3, $4, $5)"
	createAt, err := utils.StringToTimestampPostgres(tweet.CreateAt)
	if err != nil {
		return err
	}
	db.QueryRow(query, tweet.Id, tweet.Message, tweet.IdUser, createAt, tweet.Hashtag)
	return nil
}

func InsertUser(user User, db *sql.DB) error {
	var query = "INSERT INTO tb_user values ($1, $2, $3, $4)"
	db.QueryRow(query, user.Id, user.Name, user.Followers, user.Locate)
	return nil
}

func GetTopFiveUserFollowers(db *sql.DB) ([]User, error) {
	var result []User
	var query = "SELECT name, followers FROM USER ORDER BY followers DESC limit 5"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Name, &u.Followers)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

func GetPostByHour(db *sql.DB) ([]HourQtty, error) {
	var result []HourQtty
	var query = "SELECT date_trunc('hour', createAt), count(1) FROM tb_tweet GROUP BY 1"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var h HourQtty
		err := rows.Scan(&h.Hour, &h.Qtty)
		if err != nil {
			return nil, err
		}
		result = append(result, h)
	}
	return result, nil
}
