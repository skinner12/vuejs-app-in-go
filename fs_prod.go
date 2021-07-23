// +build prod

package main

import (
	"embed"
	"io/fs"

	log "github.com/sirupsen/logrus"
)

//go:embed frontend/dist
var addFrontend embed.FS

func getFrontend() fs.FS {
	f, err := fs.Sub(addFrontend, "frontend/dist")
	if err != nil {
		log.Fatalln(err)
	}

	return f
}
