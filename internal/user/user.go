package user

import (
	"github.com/Xarth-Mai/ImLLM/internal/middleware"
	"github.com/Xarth-Mai/ImLLM/internal/user/chat"
	"github.com/Xarth-Mai/ImLLM/internal/utils"
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
			HandleLogin(w, r, &userPasswd, &userToken)
		case "logout":
			middleware.WebAuth(HandleLogout(&userToken), userToken)(w, r)
		case "chat":
			middleware.WebAuth(chat.HandleChat, userToken)(w, r)
		default:
			middleware.WebAuth(HandleOK, userToken)(w, r)
		}
	}
}
