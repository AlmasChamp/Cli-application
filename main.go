package main

import (
	"log"
	"os"

	joker "github.com/AlmasChamp/Cli-application.git/service"
)

func main() {
	args := os.Args[1:]
	if err := joker.Start(args); err != nil {
		log.Println(err)
	}

}
