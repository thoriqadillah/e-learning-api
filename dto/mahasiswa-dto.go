package dto

type MahasiswaUpdate struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty" binding:"max=15,alphanum" `
}