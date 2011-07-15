package main

import (
	"exp/template"
	"path/filepath"
	. "github.com/paulbellamy/mango"
	"bytes"
	"fmt"
)

var Templates = new(template.Set)

func init() {
	fs, _ := filepath.Glob("templates/**/*.html")
	for _, tmpl := range fs {
		Templates.Add(template.MustParseFile(tmpl))
	}
	fmt.Println(Templates)
}

func VideoShow(env Env) (Status, Headers, Body) {

	buffer := bytes.NewBufferString("")
	vid := &Video{1, "YouKu", "The Lord of War", "EQWEQWEQWEQWEQWE"}
	Templates.Execute(buffer, "VideoShow.html", vid)

	return 200, Headers{}, Body(buffer.String())
}


func NotFound(env Env) (Status, Headers, Body) {
	return 200, Headers{}, Body("Not Found")
}

func main() {

	stack := new(Stack)
	stack.Address = ":3000"

	routes := map[string]App{"/videos/(.*)": VideoShow}
	stack.Middleware(Routing(routes))

	stack.Run(NotFound)
}
