package user

import "net/http"

func HandleOK(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
