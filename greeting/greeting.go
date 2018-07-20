package greeting

import (
	"fmt"
	"io"
	"net/http"
)

// Greet : So hard to greet someone new
func Greet(buffer io.Writer, name string) {
	fmt.Fprintf(buffer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
