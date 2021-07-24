# go:embed and the go-standards project layout

While experimenting with Go and trying out the go:embed directive to serve a VueJS application I was running
into an issue with trying to stick the [go-standards project layout][l-gspj]

This is because `go:embed` doesnt let you embed from parent directories and since in this specific scenario I 
was trying to have multiple services split into cmd/ subfolders this seemed to be at odds with each other.


> Patterns may not contain ‘.’ or ‘..’ or empty path elements, nor may they begin or end with a slash.

_[source][l-embed-docs]_

This led me to a deep search into some other projects that may have ran into a similar problem and I found this
commit on github https://github.com/lawrencejones/pgsink/commit/3b7c58f79e21b6b6dfffcdd8e7403bb1523747cb

Realising that by having a go package made of the web folder, I could reference it, so combining with using
build tags to either serve by the `os` package when developing or serve by `embed` when building I had a solution
that worked well.

Considering that I found many more questions asking how to do it when I searched, I figured I would provide an example.

## Run

__Run with using the embed.FS__

```
go run -tags prod cmd/http/main.go
```

__Run with the os.Dir__

```
go run cmd/http/main.go
```

Browse to http://127.0.0.1:8080 to see the index.html


[l-gspj]: https://github.com/golang-standards/project-layout
[l-embed-docs]: https://pkg.go.dev/embed#hdr-Directives