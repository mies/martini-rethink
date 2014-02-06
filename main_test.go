package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("List all todos", func() {
	It("returns a 200 Status Code", func() {
		Request("GET", "/todos", HandleIndex)
		Expect(response.Code).To(Equal(200))
	})
})

var _ = Describe("List all todos", func() {
    It("returns a 200 Status Code", func() {
        Request("POST", "/todos", HandleNewTodo)
        Expect(response.Code).To(Equal(200))
    })
})

