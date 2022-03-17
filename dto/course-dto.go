package dto

//Struct for updatable data in course table
type Course struct {
	Kode        string `json:"kode"`
	Name        string `json:"name"`
	MahasiswaID uint64 `json:"mahasiswa_id,omitempty"`
}
