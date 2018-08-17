package main

import (
	"fmt"
	"net/http"

	"github.com/stackimpact/stackimpact-go"
)

func handler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err == nil {
		defer res.Body.Close()
	}

	fmt.Fprintf(w, "Loaded some data from some API!")
}

func main() {
	agent := stackimpact.Start(stackimpact.Options{
		AgentKey: "1dc294ef4175e52d95e9a3918dff78c2024c2db3",
		AppName: "Test",
		AppVersion: "1.0.0",
		AppEnvironment: "production",
	})
	http.HandleFunc(agent.ProfileHandlerFunc("/", handler))
	http.ListenAndServe(":8080", nil)
}