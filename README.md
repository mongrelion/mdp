# Markdown Previewer

### Roadmap
  - [x] Stylesheets. Right now it looks very crappy. Maybe just import Github's css.
  - [ ] Set title to the same as the file that is being read.
  - [x] Load filename from query string: `/foo.md` instead of using flag.

### Installation
Grab the latest release for your platform [here](https://github.com/mongrelion/mdp/releases).
It's a binary file and it works out of the box.

Alternatively, if you have Go installed, you can install it using the `go install` command:

```
$ go intall https://github.com/mongrelion/mdp
```

### Usage
Run it with
```
$ mdp
```

By default it listens on http://localhost:8080/
For example, if you want to render the `README.md` file then you must access it via http://localhost:8080/README.md.  
Subdirectories are also valid: http://localhost:8080/post/2016/foobar.md

```
$ mpd --help
Usage of mdp:
-bind string
  interface to bind to, eg. 0.0.0.0:8080 (default ":8080")
-version
  prints out the version
```
