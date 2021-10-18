package assets

import (
	"embed"
	"io/fs"
)

//go:embed static/*
var staticFS embed.FS

func GetStaticFS() fs.FS {
	return staticFS
}