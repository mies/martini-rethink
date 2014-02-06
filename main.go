package main

import (
"log"
"github.com/codegangsta/martini"
"github.com/codegangsta/martini-contrib/render"
"github.com/codegangsta/martini-contrib/binding"
rethink "github.com/dancannon/gorethink"
)

type Todo struct {
	Name string
}

var session = InitDB()

func HandleIndex(r render.Render) {

	result := []Todo{}
	rows, err := rethink.Table("todos").Run(session)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var t Todo
		err := rows.Scan(&t)
		if err != nil {
			log.Println(err)
		}
		result = append(result, t)
	}
	r.JSON(200, result)
}


func HandleNewTodo(todo Todo, r render.Render) {
	_, err := rethink.Table("todos").Insert(todo).RunWrite(session)
	if err != nil {
		log.Println(err)
		return
	}
	r.JSON(200, map[string]interface{}{"todo added":"ok"})
}

func main() {

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/todos", HandleIndex)
	m.Post("/todos", binding.Json(Todo{}), HandleNewTodo )

	m.Run()
}
