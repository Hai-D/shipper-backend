package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 从环境变量中读取端口，默认值为 8081
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8081"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from shipper-backend running on port %s", port)
	})
	fmt.Printf("Starting server on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
