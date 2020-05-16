package client

import (
	"github.com/brunopita/go-common/commonsys"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var env *commonsys.Environment

func init() {
	env = commonsys.GetEnvironment()
}

func GetTwitterClient() *twitter.Client {
	config := &clientcredentials.Config{
		ClientID:     env.GetOrDefault("TW_CLIENT_ID", "QjvLHbuPh9mp8wsF9VIKi4UwS"),
		ClientSecret: env.GetOrDefault("TW_CLIENT_SECRET", "1W7w5CWvEnSIFR7YZlcf1aYaHihCJC6rRcwwd5ZDylffGuMYGm"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	httpClient := config.Client(oauth2.NoContext)
	client := twitter.NewClient(httpClient)
	return client
}
