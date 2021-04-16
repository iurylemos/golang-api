package controller_users

import "net/http"

func UpdateUSer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}
