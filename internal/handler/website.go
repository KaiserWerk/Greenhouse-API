package handler

import (
	"fmt"
	"github.com/KaiserWerk/Greenhouse-Manager/internal/caching"
	"net/http"
	"time"

	"github.com/KaiserWerk/Greenhouse-Manager/internal/templating"
)

func (h HttpHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		LastAccess string
	}{
		LastAccess: caching.GetLastAccess().Format(time.RFC3339),
	}
	if err := templating.ExecuteTemplate(w, "index.gohtml", data); err != nil {
		fmt.Printf("could not execute template: %s\n", err.Error())
	}
}
