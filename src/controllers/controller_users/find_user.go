package controller_users

import "net/http"

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find user"))
}
