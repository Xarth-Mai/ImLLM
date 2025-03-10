package main

import (
	"github.com/Xarth-Mai/ImLLM/internal/config"
	"github.com/Xarth-Mai/ImLLM/internal/server"
	"log"
	"net/http"
)

var addr string
var port string
var userPasswd map[string]string

const webFS = "./web/"

func main() {
	initErr := config.Init(&addr, &port, &userPasswd)
	if initErr != nil {
		log.Printf("Config init error: %v", initErr)
		return
	}

	http.HandleFunc("/api/", server.Run(userPasswd))
	http.Handle("/", http.FileServer(http.Dir(webFS)))

	log.Printf("Server is starting, listening on http://%s:%s", addr, port)

	err := http.ListenAndServe(addr+":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
