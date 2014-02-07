package main

import (
    "encoding/json"
    "log"
    "bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("List all todos", func() {
	It("returns a 200 Status Code", func() {
		Request("GET", "/todos", HandleIndex)
		Expect(response.Code).To(Equal(200))
	})
})

var _ = Describe("Insert new todo", func() {

    var (
        body []byte
        err error
    )

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

