package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
	"go-with-compose/app"
	"go-with-compose/app/db"
	"log"
	"net/http"
	"time"
)

func main() {
	ExampleClient()

	router := app.Routes()

	walkFunc := func(method string, route string, hander http.Handler, middlewares ...func(handler http.Handler) http.Handler) error {
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf(err.Error())
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}

func ExampleClient() {

	client := db.RedisClient()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	//err = client.Set("key", "value", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := client.Get("key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	//
	//val2, err := client.Get("key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}
	// Output: key value
	// key2 does not exist
}

type Object struct {
	Str string
	Num int
}

func Example_basicUsage() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
			"server2": ":6380",
		},
	})

	codec := &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}

	codec.Set(&cache.Item{
		Key:        key,
		Object:     obj,
		Expiration: time.Hour,
	})

	var wanted Object
	if err := codec.Get(key, &wanted); err == nil {
		fmt.Println(wanted)
	}

	// Output: {mystring 42}
}
