package go_specs_greet

import (
	"fmt"
	"net/http"

	"github.com/marvin-min/testcontainers-go-demo/specifications"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, specifications.Greet(name))
}
