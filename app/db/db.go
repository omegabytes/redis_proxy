package db

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"redis_proxy/app/models"
)

var (
	ADDR = os.Getenv("REDIS_URL")
)

func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     ADDR,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetPost(id string) *models.Post {
	var post *models.Post

	c := RedisClient()
	defer c.Close()

	val, err := c.Get(id).Bytes()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("no record for ", id)
		} else {
			panic(err)
		}
	} else {
		if err := json.Unmarshal(val, &post); err != nil {
			panic(err)
		}
	}
	return post
}

// CreatePost creates a new list entry.
func CreatePost(t models.Post) {
	c := RedisClient()
	defer c.Close()

	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	err = c.Set(t.Title, b, 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Printf("[DEBUG] POST OK %v\n", *GetPost(t.Title))
}
