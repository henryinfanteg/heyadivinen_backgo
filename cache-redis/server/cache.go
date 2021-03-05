package server

import (
	"encoding/json"
	"fmt"

	redis "github.com/go-redis/redis"
)

var client *redis.Client

func InitCache(conection *ConectionCache) (string, error) {
	if client == nil {
		fmt.Println("CLIENT ES NULL", client)
		client = redis.NewClient(&redis.Options{
			Addr:     conection.Host,
			DB:       conection.Database,
			Password: conection.Password,
		})
	} else {
		fmt.Println("CLIENT NO ES NULL", client)
	}
	return client.Ping().Result()
}

func GetData(key string) (string, error) {
	return client.Get(key).Result()
	// val, err := client.Get(key).Result()
	// if err == redis.Nil {
	// 	fmt.Println(key, "does not exist")
	// } else if err != nil {
	// 	// panic(err)
	// 	fmt.Println("ERRORRRRRRRR", err)
	// 	return nil, err
	// } else {
	// 	json.Unmarshal([]byte(val), &value)
	// }
	// return nil
}

func GetDataObject(key string, value interface{}) error {
	// return client.Get(key).Result()
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println(key, "does not exist")
	} else if err != nil {
		// panic(err)
		fmt.Println("ERRORRRRRRRR", err)
		return err
		} else {
			json.Unmarshal([]byte(val), &value)
			fmt.Println("CONVIRTIO", value)
	}
	return nil
}

func SetData(key string, value string) error {

	// valueNew, _ := json.Marshal(value)
	return client.Set(key, value, 0).Err()
}

func SetDataObject(key string, value interface{}) error {

	bytes, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
		// return
	}
	// fmt.Println(string(bytes))

	// valueNew, _ := json.Marshal(value)
	return client.Set(key, string(bytes), 0).Err()
}

// func SetData3(key string, fields map[string]interface{}) {
// 	initCache()
// 	// valueNew, _ := json.Marshal(value)

// 	// err := client.Set(key, valueNew, 0).Err()
// 	// err := client.HSet(key, json.Marshal(value), 0).Err()
// 	client.HMSET(key, fields)
// 	err := client.HMSET(key, fields).Err()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func RemoveData(key string) error {
// 	// pong, err := initCache()
// 	_, err := initCache()

// 	// err := client.Del(key).Err()
// 	if err != nil {
// 		return err
// 	}
// 	return client.Del(key).Err()
// }
