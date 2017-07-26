package main

import (
    "fmt"
    "net/http"
    "html/template"

    "github.com/julienschmidt/httprouter"
)
type User struct {
    Id int
	Name string
}
func main() {
    r := httprouter.New()
    r.GET("/", HomeHandler)

    r.POST("/new", PostsCreateHandler)


    fmt.Println("Starting server on :8080")
    http.ListenAndServe(":8080", r)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    tmpl := template.Must(template.ParseFiles("home.html"))
    users := []User{
		{1, "Tom Hanks"},
    	{2, "Will Smith"},
    	{3, "Jason Statham"},
	}
    tmpl.Execute(rw, struct{ Users []User }{users})
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    r.ParseForm()
    fmt.Fprintln(rw, r.Form.Get("name"))
}
