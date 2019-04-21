package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"net/http"
	"redis_proxy/app/cache"
	"redis_proxy/app/db"
	"redis_proxy/app/models"
)

var (
	c = cache.NewCache()
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "postID")
	post := c.Check(postId)
	c.Display()

	render.JSON(w, r, post)
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var post models.Post
	if err := json.Unmarshal(body, &post); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
	}

	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}

	db.CreatePost(post)

	response := make(map[string]string)
	response["message"] = "Created successfully"
	render.JSON(w, r, response)
}
