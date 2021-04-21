package models

// this struct senha is the format of request for change password
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
