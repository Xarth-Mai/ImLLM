package handles

import "net/http"

func HandleChat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
