// +build !prod

package main

import (
	"io/fs"
	"os"
)

func getFrontend() fs.FS {
	return os.DirFS("frontend/dist")
}
