package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"

	utils "github.com/shurcooL/github_flavored_markdown"
)

const (
	VERSION = "0.2.0"
)

var (
	file    string
	bind    string
	version bool
)

var html = `
<html>
  <head>
    <meta charset="utf-8">
    <link href="https://assets-cdn.github.com/assets/frameworks-343a7fdeaa4388a32c78fff00bca4c2f2b7d112375af9b44bdbaed82c48ad4ee.css" media="all" rel="stylesheet" type="text/css" />
    <link href="https://assets-cdn.github.com/assets/github-82746a5e80e1762d01af3e079408b886361d5fe5339de04edb1cd6df16c24eb2.css" media="all" rel="stylesheet" type="text/css" />
    <link href="//cdnjs.cloudflare.com/ajax/libs/octicons/2.1.2/octicons.css" media="all" rel="stylesheet" type="text/css" />
    <style>
      body {
        width: 800px;
        margin: auto auto;
      }
    </style>
  </head>
  <body>
    <article class="markdown-body entry-content" style="padding: 30px;">
      {{.}}
    </article>
  </body>
</html>
`

func init() {
	flag.StringVar(&bind, "bind", ":8080", "interface to bind to, eg. 0.0.0.0:8080")
	flag.BoolVar(&version, "version", false, "prints out the version")
}

func main() {
	flag.Parse()
	if version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	http.HandleFunc("/", Handler)
	log.Printf("Serving file %s on interface %s\n", file, bind)
	log.Fatal(http.ListenAndServe(bind, nil))
}

func Handler(res http.ResponseWriter, req *http.Request) {
	path := "." + req.URL.Path // prepend dot to make all file search relative to current dir
	pattern := "\\.md$"
	match, _ := regexp.MatchString(pattern, path)

	if !match {
		// return 404 here
		http.NotFound(res, req)
	} else {
		fileInfo, err := os.Stat(path)
		if err != nil {
			log.Printf("error while statting file %s\n%s", path, err)
			http.NotFound(res, req)
		} else {

			file, err := os.Open(path)
			if err != nil {
				log.Println(err)
				http.NotFound(res, req)
			}

			// Read the file here, parse it, do whatever.
			// Let's first just return the raw content
			b := make([]byte, fileInfo.Size())
			log.Printf("reading file %s\n", file.Name())
			n, err := file.Read(b)
			if err != nil {
				log.Printf("error while reading file %s\n%s", file.Name(), err)
				http.NotFound(res, req)
			}

			log.Printf("%d bytes read", n)
			content, err := ParseMD(b)
			if err != nil {
				log.Printf("error while parsing file %s\n%s", file.Name(), err)
				http.NotFound(res, req)
			} else {
				fmt.Fprintf(res, string(content))
			}
		}
	}
}

func ParseMD(b []byte) ([]byte, error) {
	tpl, err := template.New("html").Parse(html)
	if err != nil {
		return nil, err
	}

	md := utils.Markdown(b)
	x := make([]byte, 0)
	buf := bytes.NewBuffer(x)
	err = tpl.Execute(buf, string(md))
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
