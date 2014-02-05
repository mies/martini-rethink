package main

import (
"log"
	//"time"
"github.com/codegangsta/martini"
"github.com/codegangsta/martini-contrib/render"
"github.com/codegangsta/martini-contrib/binding"
rethink "github.com/dancannon/gorethink"
)

type Todo struct {
	Name string `form:"name"`
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
	log.Println("GET /todos")
	r.JSON(200, result)
}

/*
func HandleNewTodo(binding.Json(Todo{}), func(todo Todo, r render.Render) {

}
*/

func HomeHandler() string{
	log.Println("home")
	return "sdsd"
}

func main() {
	session := InitDB()

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", HomeHandler)

	m.Get("/todos", HandleIndex)

	m.Post("/todos", binding.Json(Todo{}), func(todo Todo, r render.Render) {
		_, err := rethink.Table("todos").Insert(todo).RunWrite(session)
		if err != nil {
			log.Println(err)
			return
		}

		r.JSON(200, map[string]interface{}{"todo added":"ok"})

		})


	m.Run()
}
