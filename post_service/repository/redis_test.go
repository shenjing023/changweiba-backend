package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func ExampleRedisDo1() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.0.0.202:6383",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	r, err := rdb.Do(context.Background(), "set", fmt.Sprintf("comment_count_post_%d", 111),
		22+1, "ex", 24*3600, "nx").Result()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(r.(string))
	r1, _ := rdb.Do(context.Background(), "set", fmt.Sprintf("comment_count_post_%d", 112),
		22+1, "ex", 24*3600, "nx").Result()
	fmt.Println(r1)

	r, err = rdb.SetNX(context.Background(), "aaaa", 111, 0).Result()
	fmt.Println(r)
	fmt.Println(err)
	// Output: sadaaa
	// 111
}
