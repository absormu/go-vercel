package handler

import (
	"fmt"
	"net/http"
)

func Provider(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Provider Absor")
}
