package search

import (
	"fmt"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
)

func SearchByHashtag(hastag string, client *twitter.Client) (*twitter.Search, *http.Response, error) {
	return client.Search.Tweets(&twitter.SearchTweetParams{
		Query: fmt.Sprintf("%s%s", "%23", hastag),
		Count: 100,
	})
}
