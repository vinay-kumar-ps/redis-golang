package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)
func main() {
	fmt.Println("Hello, World!")

	client := redis.NewClient(&redis.Options{
Addr: "localhost:6379",
Password: "",
DB: 0,
})
ping,err := client.Ping(context.Background()).Result()
if err!=nil {
	fmt.Println(err.Error())
	return
}
fmt.Println("Ping response:",ping)

err=client.Set(context.Background(),"name","elliot",0).Err()
if err!= nil{
	fmt.Println(err.Error())
	return
}
val,err :=client.Get(context.Background(),"name").Result()
if err!= nil{
	fmt.Println("Error getting value",err.Error())
	return
}
fmt.Println("name retrived from redis",val)
}