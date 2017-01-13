package main

import (
	"fmt"
	"log"

	loop "github.com/redis-cli/loops"

	redis "gopkg.in/redis.v5"
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
	if redisCommand == "ping" {
		res, client, err := connectClient(redisHost, redisPort)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		fmt.Println(res)
		//call  the  loop   function
		loop.LoopFunction(client)
		defer client.Close()
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
