package web

import (
	"embed"
	"io/fs"
)

//go:embed build
var staticFiles embed.FS

// GetStaticFiles returns the embedded filesystem of static files
func GetStaticFiles() fs.FS {
	fsys, err := fs.Sub(staticFiles, "build")
	if err != nil {
		panic(err)
	}
	return fsys
}
