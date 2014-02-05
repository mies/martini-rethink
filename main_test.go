package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Home", func() {
	It("returns a 200 Status Code", func() {
		Request("GET", "/todos", HandleIndex)
		Expect(response.Code).To(Equal(200))
	})
})
