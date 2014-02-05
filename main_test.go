package main

import (
"log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Home", func() {
	It("returns a 200 Status Code", func() {
		Request("GET", "/todos", HandleIndex)
        log.Println(response)
		Expect(response.Code).To(Equal(200))
	})
})
