package user

import (
	"github.com/Xarth-Mai/ImLLM/internal/user/handles"
	"github.com/Xarth-Mai/ImLLM/internal/user/middleware"
	"github.com/Xarth-Mai/ImLLM/internal/user/utils"
	"net/http"
	"strings"
)

// HandleUser starts the server and handles the api requests
func HandleUser(userPasswd map[string]string) http.HandlerFunc {
	var userToken = utils.GenerateUserMap(userPasswd)

	return func(w http.ResponseWriter, r *http.Request) {
		apiPath := strings.TrimPrefix(r.URL.Path, "/user/")

		switch apiPath {
		case "login":
			handles.HandleLogin(w, r, &userPasswd, &userToken)
		case "logout":
			middleware.WebAuth(handles.HandleLogout(&userToken), userToken)(w, r)
		case "chat":
			middleware.WebAuth(handles.HandleChat(), userToken)(w, r)
		default:
			http.Error(w, "Unknown API endpoint", http.StatusNotFound)
		}
	}
}
