package db

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:kerjakansekarang@/rest_golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect db")
	}

	return db
}

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:4321",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect redis")
	}

	return client
}
