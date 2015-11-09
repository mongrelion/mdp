package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	utils "github.com/shurcooL/github_flavored_markdown"
)

var (
	file string
	bind string
)

func init() {
	flag.StringVar(&bind, "bind", ":8080", "interface to bind to, eg. 0.0.0.0:8080")
	flag.StringVar(&file, "file", "README.md", "file to render on web interface")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", Handler)
	log.Printf("Listening on port %s\n", bind)
	log.Fatal(http.ListenAndServe(bind, nil))
}

func Handler(res http.ResponseWriter, req *http.Request) {
	readme, err := GetReadme()
	if err != nil {
		fmt.Fprintf(res, "Something went wrong:\n%s", err)
		return
	}

	fmt.Fprintf(res, string(readme))
}

func GetReadme() ([]byte, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return utils.Markdown(b), nil
}
