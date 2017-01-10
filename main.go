package main

import (
	"fmt"
	"os"

	"gopkg.in/redis.v5"
)

func main() {
	/*redisCommand := os.Args[1]
	redisHost := os.Args[2]
	redisPort := os.Args[3]*/
	fmt.Printf("%s", "goredis# ")
	var redisCommand, redisHost, redisPort string
	fmt.Scanln(&redisCommand, &redisHost, &redisPort)
	fmt.Println("Attempting to connect")
	fmt.Println(redisCommand)
	if redisCommand == "connect" {
		res, client, err := connectClient(redisHost, redisPort)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
		//call  the  loop   function
		loopFunction(client)
	}

	//readkey code
	/*	fmt.Println("Press Any key to exit.")
		b := make([]byte, 10)
		if _, err := os.Stdin.Read(b); err != nil {
			panic(err)
		}*/
}
func loopFunction(connection *redis.Client) {
	fmt.Printf("%s", "goredis# ")
	var command, arg1, arg2 string
	fmt.Scanln(&command, &arg1, &arg2)
	//Set command
	if command == "set" {
		setCommand(connection, arg1, arg2)
	} else if command == "get" {
		getCommand(connection, arg1)
	} else {
		fmt.Println("Press Any key to exit.")
		b := make([]byte, 10)
		if _, err := os.Stdin.Read(b); err != nil {
			panic(err)
		}
	}

}

//Redis Client Connection function
func connectClient(host string, port string) (response string, connection *redis.Client, Error error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		return "Error connecting redis", nil, err

	}
	response, connection, Error = pong, client, nil
	return
}

//Set Command Function
func setCommand(client *redis.Client, arg1 string, arg2 string) {
	err := client.Set(arg1, arg2, 0)
	fmt.Println(err)
	loopFunction(client)
}

//get Command Function
func getCommand(client *redis.Client, arg1 string) {
	val, err := client.Get(arg1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	loopFunction(client)
}
