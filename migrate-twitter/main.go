package main

import (
	"log"

	"github.com/brunopita/twitter-example/twitter-consumer/spg"
)

func main()  {
	db, err := spg.GetConnection()
	if err != nil {
		log.Fatalln(Err)
	}
	db.Close()
}