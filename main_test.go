package main

import (
    "encoding/json"
    "log"
    "bytes"
    //"io/ioutil"
    "github.com/bitly/go-simplejson"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)
var _ = Describe("Todo", func() {

    var (
         js *simplejson.Json
        body []byte
        err error
        //todos []Todo
    )

    Describe("List all todos", func() {
        It("returns a 200 Status Code", func() {
            Request("GET", "/todos", HandleIndex)
            Expect(response.Code).To(Equal(200))
        })
    })

    Describe("Create a Todo", func() {

        BeforeEach(func() {
            todo := Todo{"keep things green"}
            body, err = json.Marshal(todo)
            if err != nil {
                log.Println("Unable to marshal todo")
            }
        })

        It("returns a 200 Status Code", func() {
            PostRequest("POST", "/todos", HandleNewTodo, bytes.NewReader(body))
            Expect(response.Code).To(Equal(200))
        })
    })

    Describe("Returns a single created todo", func() {
        It("returns a 200 Status Code", func() {
            Request("GET", "/todos", HandleIndex)
            log.Println(response)
            log.Println(response.Body)
            Expect(response.Code).To(Equal(200))
            js, err = simplejson.NewJson([]byte(response.Body.String()))

            log.Println(js)


            log.Println(js.GetIndex(1).Get("Name"))


            /*body, err = ioutil.ReadAll(response.Body)
            log.Println("=====")
            log.Println(body)
            err = json.Unmarshal(body, &todos)
              log.Println(err)
            log.Println(todos)
            */
        })
    })
})






