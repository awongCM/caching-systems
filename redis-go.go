package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	// String values
	client.Set("newkey", "ofsomevalue", 0)
	value, err := client.Get("newkey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	// Binary values
	client.Set("pages:about", "about us", 0)
	value, err = client.Get("pages:about").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	// Counter values
	client.Set("counter", 100, 0)
	value, err = client.Get("counter").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	client.Incr("counter")
	value, err = client.Get("counter").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	client.IncrBy("counter", 50)
	value, err = client.Get("counter").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	// MSET/MGET
	client.MSet("a", 10, "b", 20, "c", 30)
	mvalue, err := client.MGet("a", "b", "c").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(mvalue)

	// ALTERING/QUERYING KEY
	client.Set("mykey", "hello", 0)
	mykeyexists, err := client.Exists("mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(mykeyexists)

	client.Del("mykey")
	mykeyexists, err = client.Exists("mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(mykeyexists)

	keytype, err := client.Type("mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keytype)

	// SET TIMEOUT
	client.Set("mykey", "somevalue", 0)
	client.Expire("mykey", 6*time.Second)
	beforevalue, err := client.Get("mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("before expiry: " + beforevalue)

	time.Sleep(6 * time.Second)

	// TODO
	// aftervalue, err := client.Get("mykey").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("after expiry: " + aftervalue)

	// TTL
	client.Set("mykey", "100", 0)
	client.Expire("mykey", 10*time.Second)
	ttlValue := client.TTL("mykey")
	if err != nil {
		panic(err)
	}
	fmt.Println(ttlValue)

	// LIST OPERATIONS

	// remove existing values from list, if any
	client.Del("myList")

	client.RPush("myList", "a")
	client.RPush("myList", "b")
	client.LPush("myList", "first")
	fmt.Println(client.LRange("myList", 0, -1))

	client.RPush("myList", []string{strconv.Itoa(1), strconv.Itoa(2), strconv.Itoa(3), strconv.Itoa(4), strconv.Itoa(5), "foo bar"})
	fmt.Println(client.LRange("myList", 0, -1))

	client.RPush("numberlist", []int{1, 2, 3}) //something's not right here
	fmt.Println(client.LRange("numberlist", 0, -1))
	popedvalue, err := client.RPop("numberlist").Result()
	fmt.Println("Popped value: " + popedvalue)
	popedvalue, err = client.RPop("numberlist").Result()
	fmt.Println("Popped value: " + popedvalue)
	popedvalue, err = client.RPop("numberlist").Result()
	fmt.Println("Popped value: " + popedvalue)

	// CAPPED LIST
	client.RPush("mycappedList", []int{1, 2, 3, 4, 5}) // and here...
	client.LTrim("mycappedList", 0, 2)
	fmt.Println("Capped list:.... ")
	fmt.Println(client.LRange("mycappedList", 0, -1))

}
