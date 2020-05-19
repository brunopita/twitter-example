package tdao

import (
	"database/sql"

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

type QttyHashtagLocate struct {
	Qtty    int
	Hashtag string
	Locate  string
}

func InsertTweet(tweet Tweet, db *sql.Tx) error {
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

func InsertUser(user User, db *sql.Tx) error {
	var query = "INSERT INTO tb_user values ($1, $2, $3, $4) ON CONFLICT (id) DO NOTHING"
	_, err := db.Exec(query, user.Id, user.Name, user.Followers, user.Locate)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int64, db *sql.Tx) error {
	var query = "DELETE FROM tb_user where id = $1"
	db.QueryRow(query, id)
	return nil
}

func DeleteTweet(id int64, db *sql.Tx) error {
	var query = "DELETE FROM tb_tweet where id = $1"
	db.QueryRow(query, id)
	return nil
}

func GetTopFiveUserFollowers(db *sql.Tx) ([]User, error) {
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

func GetQttyForHourByHashtag(db *sql.Tx) ([]QttyHourHashtag, error) {
	var result []QttyHourHashtag
	var query = "select  hashtag, extract(hour from createat), count(1) from tb_tweet group by 2, hashtag order by 2,1;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var r QttyHourHashtag
		err := rows.Scan(&r.Hashtag, &r.Hour, &r.Qtty)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

func GetQttyPostForHashtagByLocate(db *sql.Tx) ([]QttyHashtagLocate, error) {
	var result []QttyHashtagLocate
	var query = "select count(*), hashtag, u.locate from tb_tweet inner join tb_user as u on iduser = u.id group by 2,3 order by hashtag;"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var r QttyHashtagLocate
		err := rows.Scan(&r.Qtty, &r.Hashtag, &r.Locate)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
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
