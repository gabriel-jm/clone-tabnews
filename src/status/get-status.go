package status

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ServerStatus struct {
	UpdatedAt string `json:"updatedAt"`
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Status OK")
	w.Header().Set("Content-Type", "application/json")

	jsonBytes, _ := json.Marshal(ServerStatus{
		UpdatedAt: time.Now().Format(time.RFC3339),
	})
	io.WriteString(w, string(jsonBytes))
}
