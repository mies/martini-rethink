package main

import (
    "encoding/json"
    "log"
    "bytes"
    //"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
    //. "github.com/onsi/gomega/matchers"
)
var _ = Describe("Todo", func() {

    var (
        body []byte
        err error
        //todos []Todo
    )

    Context("List all todos", func() {
        It("returns a 200 Status Code", func() {
            Request("GET", "/todos", HandleIndex)
            Expect(response.Code).To(Equal(200))
        })
    })

    Context("Create a Todo", func() {

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

    Context("Returns a single created todo", func() {
        It("returns a 200 Status Code", func() {
            Request("GET", "/todos", HandleIndex)
            log.Println(response)
            log.Println(response.Body)
            Expect(response.Code).To(Equal(200))

            Expect(response.Body).To(MatchJSON(`[{"Name":"keep things green"}]`))
        })
    })
})