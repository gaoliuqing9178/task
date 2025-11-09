package dto

type UpdateUserInfoReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
