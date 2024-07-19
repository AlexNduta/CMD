package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

type commandRequest struct {
	Command string
}

type commandResponse struct {
	Output string
	Error  string
}

func main() {
	http.HandleFunc("/api/cmd", func(w http.ResponseWriter, r *http.Request) {
		var req commandRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body: %v", err)
			return
		}

		cmd := exec.Command("bash", "-c", req.Command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Command not found: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", out)
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
