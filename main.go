package main

import (
	  "log"
	  "time"
	  "github.com/codegangsta/martini"
	  "github.com/codegangsta/martini-contrib/render"
)

type Todo struct {
	name string
	duedate time.Date
}

func main() {
    m := martini.Classic()
    m.Use(render.Renderer())

    m.Get("/todos", func(r render.Render) {
        r.HTML(200, "list", GetAll(db))
    })

    m.Run()
}
