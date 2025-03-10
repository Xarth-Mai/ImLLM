package server

import (
	"github.com/Xarth-Mai/ImLLM/internal/server/api/openai"
	"github.com/Xarth-Mai/ImLLM/internal/server/handles"
	"github.com/Xarth-Mai/ImLLM/internal/server/middleware"
	"github.com/Xarth-Mai/ImLLM/internal/server/utils"
	"net/http"
	"strings"
)

// Run starts the server and handles the api requests
func Run(userPasswd map[string]string) http.HandlerFunc {
	var userToken = utils.GenerateUserMap(userPasswd)

	return func(w http.ResponseWriter, r *http.Request) {
		apiPath := strings.TrimPrefix(r.URL.Path, "/api/")

		switch apiPath {
		case "login":
			handles.HandleLogin(w, r, &userPasswd, &userToken)
		case "logout":
			middleware.WebAuth(handles.HandleLogout(&userToken), userToken)(w, r)
		case "chat":
			middleware.WebAuth(handles.HandleChat(), userToken)(w, r)
		case "openai":
			middleware.ApiAuth(openai.HandleOpenAI, userPasswd)(w, r)
		default:
			http.Error(w, "Unknown API endpoint", http.StatusNotFound)
		}
	}
}
