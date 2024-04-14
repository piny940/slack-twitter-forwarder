package main

import (
	"flag"
	"fmt"
	"slack-twitter-forwarder/server"

	"github.com/joho/godotenv"
)

func main() {
	env := flag.String("e", "development", "set environment")
	flag.Parse()

	loadDotenv(*env)

	if err := server.Init(); err != nil {
		panic(err)
	}
}

func loadDotenv(env string) {
	if env != "production" {
		if err := godotenv.Load(fmt.Sprintf(".env.%s", env)); err != nil {
			panic(err)
		}
	}
}
