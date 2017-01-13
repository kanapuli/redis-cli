package loops

import (
	"fmt"
	"log"
	"os"

	basics "github.com/redis-cli/basiccommands"

	redis "gopkg.in/redis.v5"
)

//LoopFunction  - Keeps the terminal running forever
func LoopFunction(connection *redis.Client) {
	fmt.Printf("%s", "goredis# ")
	var command, arg1, arg2 string
	fmt.Scanln(&command, &arg1, &arg2)
	//Set command
	if command == "set" {
		basics.SetCommand(connection, arg1, arg2)
	} else if command == "get" {
		basics.GetCommand(connection, arg1)
	} else {
		fmt.Println("Press Any key to exit.")
		b := make([]byte, 10)
		if _, err := os.Stdin.Read(b); err != nil {
			log.Fatal(err)

		}
	}

}
