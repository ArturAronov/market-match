package handlers

import "net/http"

func GetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to recieve order: Expexted method GET, recieved" + r.Method))
	}
}
