package routes

import (
	"encoding/json"
	"net/http"

	"github.com/brunopita/twitter-example/twitter-api/controller"
)

func TopFiveFollow(resp http.ResponseWriter, req *http.Request) {
	result := controller.TopFiveFollowController()
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return		
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(js) 
}