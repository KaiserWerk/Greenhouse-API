package handler

import (
	"fmt"
	"net/http"
)

func (h HttpHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "This is the start page of the Greenhouse API")
}
