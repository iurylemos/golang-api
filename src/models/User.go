package models

import "time"

// using for first time omitempty
// when for pass this user to json and the id been in blanck, this not going to pass
// it don't going pass this for json, it go removed the id of json
type User struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"Criadoem,omitempty"`
}
