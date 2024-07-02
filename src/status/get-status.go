package status

import (
	"fmt"
	"io"
	"net/http"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Status OK")
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Status OK")
}
