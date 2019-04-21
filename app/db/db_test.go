package db

import (
	"redis_proxy/app/models"
	"testing"
)

func TestCreatePost(t *testing.T) {

	post := models.Post{Title: "CreatePostTest", Body: "this one is created in the db pkg test"}
	CreatePost(post)

	if GetPost("CreatePostTest").Body != "this one is created in the db pkg test" {
		t.Errorf("Expected title to be 'CreatePostTest'")
	}
}
