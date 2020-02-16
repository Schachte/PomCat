package router

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Method string

const (
	POST  Method = "POST"
	GET   Method = "GET"
	PUT   Method = "PUT"
	PATCH Method = "PATCH"
)

var RequestMapping = map[string]Method{
	"POST": POST,
	"GET":  GET,
}

type CustomRouter struct {
	handlers map[string]func(rq http.ResponseWriter, r *http.Request)
	database *sql.DB
}

func InitializeRouter(db *sql.DB) *CustomRouter {
	router := new(CustomRouter)
	router.handlers = make(map[string]func(rq http.ResponseWriter, r *http.Request))
	router.database = db
	return router
}

func (c *CustomRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	pathSlice := strings.Split(r.URL.Path, "/")

	currentReqMethod, ok := RequestMapping[r.Method]

	if !ok {
		log.Println("There was an issue retrieving this type of request")
		return
	}

	if handler, ok := c.handlers[Key(currentReqMethod, pathSlice[1])]; ok {
		fmt.Printf("Currently handling %s with %s\n", r.Method, pathSlice[1])
		handler(rw, r)
		return
	}

	fmt.Printf("Sorry, there was an error handling this request %s", pathSlice[0])
}

func (c *CustomRouter) POST(path string, h http.HandlerFunc) {
	c.handlers[Key(POST, path)] = h
}

func (c *CustomRouter) GET(path string, h http.HandlerFunc) {
	fmt.Printf("Registering %s with %s\n", path, GET)
	c.handlers[Key(GET, path)] = h
}

func Key(method Method, path string) string {
	return fmt.Sprintf("%s:%s", path, method)
}
