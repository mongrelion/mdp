package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	utils "github.com/shurcooL/github_flavored_markdown"
)

func main() {
	http.HandleFunc("/", Handler)
	log.Printf("Listening on port %d\n", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
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
	b, err := ioutil.ReadFile("./README.md")
	if err != nil {
		return nil, err
	}

	return utils.Markdown(b), nil
}
