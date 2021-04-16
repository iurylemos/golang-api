package controller_users

import "net/http"

func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find users"))
}
