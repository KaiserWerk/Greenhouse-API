package handler

import (
	"fmt"
	"net/http"

	"github.com/KaiserWerk/Greenhouse-Manager/internal/templates"
)

func (h HttpHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.gohtml", nil); err != nil {
		fmt.Printf("could not execute template: %s\n", err.Error())
	}
}
