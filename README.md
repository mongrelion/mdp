# Markdown Previewer

### Roadmap
  - [x] Stylesheets. Right now it looks very crappy. Maybe just import Github's css.
  - [ ] Set title to the same as the file that is being read.
  - [ ] Load filename from query string: `/foo.md` instead of using flag.

### Installation
Grab the latest release for your platform [here](https://github.com/mongrelion/mdp/releases).
It's a binary file and it works out of the box.

Alternatively, if you have Go installed, you can install it using the `go install` command:

```
$ go intall https://github.com/mongrelion/mdp
```

### Usage
By default `mdp` will serve the **README.md** file in the current directory on **http://localhost:8080**
but of course these defaults are overridable:

```
$ mpd --help
Usage of mdp:
-bind string
  interface to bind to, eg. 0.0.0.0:8080 (default ":8080")
-file string
  file to render on web interface (default "README.md")
-version
  prints out the version
```
