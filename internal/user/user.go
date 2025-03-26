package user

import (
	"github.com/Xarth-Mai/ImLLM/internal/dialog"
	"github.com/Xarth-Mai/ImLLM/internal/middleware"
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
		case "dialog":
			middleware.WebAuth(dialog.HandleDialog, userToken)(w, r)
		default:
			middleware.WebAuth(HandleOK, userToken)(w, r)
		}
	}
}
