package dto

type RegisterReq struct {
	Name     string `json:"name"`
	Account  string `json:"account"`
	Password string `json:"password"`
}
