package main

import (
	"encoding/json"
	"syscall"

	"github.com/ambiware-labs/loqa-core/skills/examples/internal/host"
)

type request struct {
	Message string `json:"message"`
}

type response struct {
	Reply string `json:"reply"`
}

func run() {
	var req request
	_ = json.Unmarshal([]byte(getEnv("LOQA_EVENT_PAYLOAD")), &req)

	reply := response{Reply: "Hello from TinyGo!"}
	payload, _ := json.Marshal(reply)

	if host.Publish("skill.example.response", payload) {
		host.Log("response published")
	} else {
		host.Log("publish failed")
	}
}

func getEnv(key string) string {
	if v, ok := syscall.Getenv(key); ok {
		return v
	}
	return ""
}

//go:export run
func main() {
	run()
}
