package main

import (
	"net/http"

	"github.com/zogot/go-embed-example/web"
)

func main() {
	// This is using the standard http library but if you return a native fs.FS you
	// can generally convert to any of the other http libraries easily.

	// Make the call to the web package
	wfs := web.GetFileSystem()

	// Put into a handle for the root path, given to the http.FS
	http.Handle("/", http.FileServer(http.FS(wfs)))

	// Start the server
	http.ListenAndServe(":8080", nil)
}
