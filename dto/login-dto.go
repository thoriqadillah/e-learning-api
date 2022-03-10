package dto

type Login struct {
	NIM      string `json:"nim"  binding:"required,min=15,max=15"`
	Password string `json:"password,omitempty" binding:"required,max=15,alphanum"`
}
