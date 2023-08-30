package web

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:build
var content embed.FS

func Build() http.FileSystem {
	build, err := fs.Sub(content, "build")
	if err != nil {
		log.Fatal(err)
	}

	return http.FS(build)
}
