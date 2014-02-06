package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
    "io"
	"net/http"
	"net/http/httptest"
    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/render"
    "github.com/codegangsta/martini-contrib/binding"

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

func PostRequest(method string, route string, handler martini.Handler, body io.Reader) {
  m := martini.Classic()
  m.Post(route, binding.Json(Todo{}), handler)
  m.Use(render.Renderer())
  request, _ := http.NewRequest(method, route, body)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
}