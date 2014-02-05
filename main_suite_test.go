package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
    "log"
	"testing"
	"net/http"
	"net/http/httptest"
    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/render"

)

var (
  response *httptest.ResponseRecorder
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main Suite")
}

func Request(method string, route string, handler martini.Handler) {
  m := martini.Classic()
  m.Get(route, handler)
  m.Use(render.Renderer())
  request, _ := http.NewRequest(method, route, nil)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
}

