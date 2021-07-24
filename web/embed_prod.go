// +build prod

package web

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed dist
var webFs embed.FS

func init() {
	_, err := webFs.Open("dist/index.html")
	if err != nil {
		panic("no dist/index.html found for web. You could use this point to check if the build has been ran")
	}
}

func GetFileSystem() fs.FS {
	fmt.Println("Filesystem served by embed package")
	sfs, _ := fs.Sub(webFs, "dist")
	return sfs
}
