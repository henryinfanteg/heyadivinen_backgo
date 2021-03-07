package main

import (
	"fmt"

	"github.com/henryinfanteg/heyadivinen_backgo/cache-redis/server"
	redis "github.com/go-redis/redis"
)

var client *redis.Client

func main() {

	conection := server.ConectionCache{Host: "localhost:6379", Database: 0, Password: ""}
	fmt.Println(conection)
	_, err := server.InitCache(&conection)
	if err != nil {
		fmt.Println("ERROR INIT CACHE", err)
		// return nil, err
	} else {
		fmt.Println("CACHE OK")
	}

	if err = server.SetDataObject("keyObj1", conection); err != nil {
		fmt.Println("ERROR SET CACHE OBJ 1", err)
	}
	if err = server.SetData("key1", "leo1"); err != nil {
		fmt.Println("ERROR SET CACHE 1", err)
	}
	if err = server.SetData("key2", "leo2"); err != nil {
		fmt.Println("ERROR SET CACHE 2", err)
	}
	
	valObj1, errObj1 := server.GetData("keyObj1");
	fmt.Println("valObj1", valObj1, errObj1)
	
	var conectionCache server.ConectionCache
	errObj2 := server.GetDataObject("keyObj1", conectionCache);
	fmt.Println("valObj2", conectionCache, errObj2)
	
	val1, err1 := server.GetData("key1");
	fmt.Println("val1", val1, err1)
	
	val2, err2 := server.GetData("key2");
	fmt.Println("val2", val2, err2)

	// ExampleNewClient()
	// fmt.Println(GetData("key"))
	// val := GetData("key4")
	// fmt.Println("val", val)
	// SetData("key3", "leoooo")
	// fmt.Println("key3", GetData("key3"))
}

func getClient() *redis.Client {
	client = redis.NewClient(&redis.Options{
		// Addr:     "192.168.1.103:6379",
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	if len(pong) != 0 {
		return client
	} else {
		return nil
	}
}

func CreateNewClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// fmt.Println("client ->", client)

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func initCache() {
	if client == nil {
		// fmt.Println("-------------------")
		// fmt.Println("CLIENT ES NULL", client)
		CreateNewClient()
	} else {
		// fmt.Println("-------------------")
		// fmt.Println("CLIENT NO ES NULL", client)
	}
}

func GetData(key string) string {
	initCache()

	val, err := client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println(key, "does not exist")
	} else if err != nil {
		panic(err)
	} else {
		// fmt.Println("key2", val2)
	}
	// fmt.Println("key", val)
	return val
}

func SetData(key string, value string) {
	initCache()

	err := client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}
