package routes

import (
	"encoding/json"
	"net/http"

	"github.com/brunopita/twitter-example/twitter-api/controller"
	"github.com/brunopita/twitter-example/twitter-pg/tdao"
)

func TopFiveFollow(w http.ResponseWriter, req *http.Request) {
	var err error
	var result []tdao.User
	result, err = controller.TopFiveFollowController(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func PostsForHour(w http.ResponseWriter, req *http.Request) {
	var err error
	var result []tdao.QttyHourHashtag

	result, err = controller.PostsForHourController(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func TotalPostHashtagByLocate(w http.ResponseWriter, req *http.Request) {
	var err error
	var result []tdao.QttyHashtagLocate

	result, err = controller.TotalPostHashtagByLocate(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
