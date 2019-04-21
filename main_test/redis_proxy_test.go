package main_test

import (
	"fmt"
	"os"
	"redis_proxy/app"
	"redis_proxy/app/cache"
	"redis_proxy/app/db"
	"redis_proxy/app/models"
	"testing"
	"time"
)

var (
	SIZE = os.Getenv("CACHE_SIZE")
	TTL  = os.Getenv("CACHE_TTL")
	URL  = os.Getenv("REDIS_URL")
)

func TestMain(m *testing.M) {
	a := app.App{}
	a.Initialize()
	InitDBVals()
	printEnv()
	code := m.Run()
	os.Exit(code)
}

func TestClient(t *testing.T) {
	client := db.RedisClient()
	_, err := client.Ping().Result()
	if err != nil {
		t.Errorf("Expected response from client")
	}
}

func TestCacheLRU_Baseline(t *testing.T) {
	fmt.Println("\n This test shows the cache after each GET")

	c := cache.NewCache()

	for i := 0; i < 3; i++ {
		c.Check("first")
		c.Display()
	}

	for i := 0; i < 3; i++ {
		c.Check("second")
		c.Display()
	}

	for i := 0; i < 3; i++ {
		c.Check("third")
		c.Display()
	}

	c.Check("seventh")
	c.Display()
	c.Check("fifth")
	c.Display()
	c.Check("first")
	c.Display()

	if c.Queue.Length != 5 {
		t.Errorf("Expected cache length of 5")
	}
}

func TestCacheLRU_TTL(t *testing.T) {
	c := cache.NewCache()

	c.Check("first")
	c.Check("second")
	c.Check("second")
	c.Check("second")
	time.Sleep(5 * time.Second)
	c.Check("second")
	c.Check("second")

}

func InitDBVals() {
	example := []*models.Post{
		{Title: "first", Body: "one"},
		{Title: "second", Body: "two"},
		{Title: "third", Body: "three"},
		{Title: "fourth", Body: "four"},
		{Title: "fifth", Body: "five"},
		{Title: "sixth", Body: "six"},
		{Title: "seventh", Body: "seven"},
		{Title: "eighth", Body: "eight"},
		{Title: "nineth", Body: "nine"},
		{Title: "tenth", Body: "ten"},
	}

	for _, item := range example {
		db.CreatePost(*item)
	}
}

func printEnv() {
	fmt.Println("size: ", SIZE)
	fmt.Println("ttl: ", TTL)
	fmt.Println("url: ", URL)
}
