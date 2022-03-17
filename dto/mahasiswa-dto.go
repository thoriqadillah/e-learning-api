package dto

//Struct for updatable data in mahasiswa table
type Mahasiswa struct {
	Name     string `json:"name,omitempty" binding:"omitempty"`
	Email    string `json:"email,omitempty" binding:"omitempty,email"`
	Password string `json:"password,omitempty" binding:"omitempty,max=15,alphanum"`
}
