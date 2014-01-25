package main

import (
	//"os"
	"log"
	"time"
	r "github.com/dancannon/gorethink"
)

//var session *r.Session

func InitDB() *r.Session {
	//var err error

	session, err := r.Connect(map[string]interface{} {
		//"address" : os.Getenv("RETHINKDB_URL"),
		"address" : "localhost:28015",
		"database": "test",
		"maxIdle" : 10,
		"idleTimeout": time.Second  * 10,
	})

	if err != nil {
		log.Println(err)
	}
	err = r.DbCreate("todos").Exec(session)
    if err != nil {
    	log.Println(err)
    }

	response, err := r.Db("test").TableCreate("todos").RunWrite(session)
	if err != nil {
		log.Println(err)
	}
	log.Println(response)

	return session
}
