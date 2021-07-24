// +build !prod

package web

import (
	"fmt"
	"io/fs"
	"os"
)

func GetFileSystem() fs.FS {
	fmt.Println("Filesystem served by os package")
	return os.DirFS("web/dist")
}
