package main

import (
	"os"
	"log"
	"time"
	r "github.com/dancannon/gorethink"
)

var session *r.Session

func InitDB() {
	//var err error

	session, err := r.Connect(map[string]interface{} {
		"address" : os.Getenv("RETHINKDB_URL"),
		"database": "todos",
		"maxIdle" : 10,
		"idleTimeout": time.Second  * 10,
	})

	if err != nil {
		log.Println(err)
	}
}
