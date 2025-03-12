package main

import (
	"github.com/Xarth-Mai/ImLLM/internal/api/openai"
	"github.com/Xarth-Mai/ImLLM/internal/config"
	"github.com/Xarth-Mai/ImLLM/internal/middleware"
	"github.com/Xarth-Mai/ImLLM/internal/user"
	"log"
	"net/http"
)

const webFS = "./web/"
const configPath = "./config.yaml"

func main() {
	UserConfig, initErr := config.InitConfig(configPath)
	if initErr != nil {
		log.Printf("Config init error: %v", initErr)
		return
	}

	http.HandleFunc("/user/", user.HandleUser(UserConfig.UserPasswd))
	http.HandleFunc("/openai/", middleware.ApiAuth(openai.HandleOpenAI, UserConfig.UserPasswd))
	http.Handle("/", http.FileServer(http.Dir(webFS)))

	log.Printf("Server is starting, listening on http://%s:%s", UserConfig.Addr, UserConfig.Port)

	if err := http.ListenAndServe(UserConfig.Addr+":"+UserConfig.Port, nil); err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
