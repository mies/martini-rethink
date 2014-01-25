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

func main() {
	session := InitDB()

    m := martini.Classic()
    m.Use(render.Renderer())

    m.Get("/todos", func(r render.Render)  {
		rows, err := rethink.Table("todos").Run(session)
		if err != nil {
			log.Println(err)
		}

		log.Println(rows)

		r.JSON(200, map[string]interface{}{"hello": "world"})
    })

	m.Post("/todos", binding.Form(Todo{}), func(todo Todo, r render.Render) {

		log.Println(todo)
		_, err := rethink.Table("todos").Insert(todo).RunWrite(session)
		if err != nil {
			log.Println(err)
		}

		r.JSON(200, map[string]interface{}{"yes":"sir"})

	})


    m.Run()
}
