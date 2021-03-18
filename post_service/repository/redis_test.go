package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

// func ExampleRedisDo1() {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "10.0.0.202:6383",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})

// 	r, err := rdb.Do(context.Background(), "set", fmt.Sprintf("comment_count_post_%d", 111),
// 		22+1, "ex", 24*3600, "nx").Result()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// fmt.Println(r.(string))
// 	r1, _ := rdb.Do(context.Background(), "set", fmt.Sprintf("comment_count_post_%d", 112),
// 		22+1, "ex", 24*3600, "nx").Result()
// 	fmt.Println(r1)

// 	r, err = rdb.SetNX(context.Background(), "aaaa", 111, 0).Result()
// 	fmt.Println(r)
// 	fmt.Println(err)
// 	// Output: sadaaa
// 	// 111
// }

func ExamplePipeline() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.0.0.202:6383",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// go RunPipeline1(rdb)
	go RunPipeline2(rdb)
	// go RunPipeline3(rdb)

	time.Sleep(time.Second * 1)

	// Output: xxxx
	// xas
}

func RunPipeline1(rdb *redis.Client) {
	ctx := context.Background()
	pipe := rdb.Pipeline()
	pipe.Get(ctx, "1").Result()
	pipe.Get(ctx, "2").Result()
	pipe.Get(ctx, "3").Result()
	pipe.Get(ctx, "4").Result()
	pipe.Get(ctx, "3").Result()
	pipe.Get(ctx, "4").Result()
	pipe.Get(ctx, "1").Result()
	pipe.Get(ctx, "2").Result()
	pipe.Get(ctx, "3").Result()
	pipe.Get(ctx, "4").Result()
	pipe.Get(ctx, "3").Result()
	pipe.Get(ctx, "4").Result()
	cmders, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println("err", err)
	}
	for _, cmder := range cmders {
		cmd := cmder.(*redis.StringCmd)
		r, err := cmd.Result()
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println("r:", r)
	}
}

func RunPipeline2(rdb *redis.Client) {
	ctx := context.Background()
	// _, err := rdb.HGetAll(ctx, "33").Result()
	// fmt.Println(err)
	r, err := rdb.HMGet(ctx, "5", "1").Result()
	fmt.Println("r:", r)
	// fmt.Println(err)
	pipe := rdb.Pipeline()
	// pipe.Get(ctx, "2").Result()
	// pipe.Get(ctx, "3").Result()
	// pipe.Get(ctx, "4").Result()
	pipe.HMGet(ctx, "11", "1", "2")
	pipe.HMGet(ctx, "22", "1", "2")
	pipe.HMGet(ctx, "33", "1", "2")
	pipe.HMGet(ctx, "44", "1", "2")
	cmders, err := pipe.Exec(ctx)
	// if err != nil {
	fmt.Println("err2", err)
	// }
	for _, cmder := range cmders {
		cmd := cmder.(*redis.SliceCmd)
		type A struct {
			A1 int `redis:"1"`
			A2 int `redis:"2"`
		}
		var a A
		err := cmd.Scan(&a)
		// if err != nil {
		fmt.Println("errr2", err)
		// }
		fmt.Printf("%+v \n", a)
		// if r == nil {
		// 	fmt.Println("nil:", r)
		// }
		// fmt.Println("r2:", r)
	}
}

func RunPipeline3(rdb *redis.Client) {
	ctx := context.Background()
	pipe := rdb.Pipeline()
	pipe.Get(ctx, "2").Result()
	pipe.Get(ctx, "3").Result()
	pipe.Get(ctx, "4").Result()
	pipe.Get(ctx, "1").Result()
	cmders, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println("err", err)
	}
	for _, cmder := range cmders {
		cmd := cmder.(*redis.StringCmd)
		r, err := cmd.Result()
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println("r3:", r)
	}
}

func BenchmarkHMGETB(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.0.0.202:6383",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		rdb.HMGet(ctx, "11", "1", "2", "3").Result()
		// fmt.Println(r)
	}
}

func BenchmarkHGETALLB(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.0.0.202:6383",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		rdb.HGetAll(ctx, "11").Result()
	}
}
