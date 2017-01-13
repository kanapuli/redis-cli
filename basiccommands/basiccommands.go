package basics

import (
	"fmt"
	"log"

	redis "gopkg.in/redis.v5"
)

//SetCommand Set Command Function
func SetCommand(client *redis.Client, arg1 string, arg2 string) {
	err := client.Set(arg1, arg2, 0)
	fmt.Println(err)
	//loop.LoopFunction(client)
}

//GetCommand get Command Function
func GetCommand(client *redis.Client, arg1 string) {
	val, err := client.Get(arg1).Result()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Key doesnt exists")
	}
	fmt.Println(val)
	//loop.LoopFunction(client)
}
