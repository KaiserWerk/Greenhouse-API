package assets

import (
	"embed"
	"io/fs"
	"path/filepath"
)

//go:embed templates/*
var templateFS embed.FS

//go:embed static/*
var staticFS embed.FS

func GetTemplateFile(name string) ([]byte, error) {
	return templateFS.ReadFile(filepath.Join("templates", name))
}

func GetStaticFS() fs.FS {
	return staticFS
}