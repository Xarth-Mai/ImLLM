package server

import (
	"github.com/Xarth-Mai/ImLLM/internal/server/api/openai"
	"github.com/Xarth-Mai/ImLLM/internal/server/handles"
	"github.com/Xarth-Mai/ImLLM/internal/server/utils"
	"net/http"
	"strings"
)

// Run starts the server and handles the api requests
func Run(userPasswd map[string]string) http.HandlerFunc {
	var UserMap = utils.GenerateUserMap(userPasswd)
	var userToken = UserMap
	var userTokenSeed = UserMap

	return func(w http.ResponseWriter, r *http.Request) {
		apiPath := strings.TrimPrefix(r.URL.Path, "/api/")

		switch apiPath {
		case "login":
			handles.HandleLogin(w, r, &userPasswd, &userTokenSeed, &userToken)
		case "logout":
			handles.HandleLogout(w, r, &userToken)
		case "chat":
			handles.HandleChat(w, r, &userToken)
		case "openai":
			openai.HandleOpenAI(w, r, &userToken)
		default:
			http.Error(w, "Unknown API endpoint", http.StatusNotFound)
		}
	}
}
