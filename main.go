package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type Fruit struct {
	FruitName   string `json:"name"`
	FruitExpire int    `json:"expire"`
}

func main() {
	fmt.Println("Welcome")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Fruit{FruitName: "Mango", FruitExpire: 25})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(json)
	}

	err = client.Set("fruit01", json, 0).Err()

	if err != nil {
		fmt.Println(err)
	}

	get, err := client.Get("fruit01").Result()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(get)
	}

	ping, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ping)
	}

	err = client.Set("name", "Ravi", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, errors := client.Get("name").Result()
	if errors != nil {
		fmt.Println(errors)
	} else {
		fmt.Println(val)
	}
}
