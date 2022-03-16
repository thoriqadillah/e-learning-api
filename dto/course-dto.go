package dto

type Course struct {
	Kode        string `json:"kode"`
	Name        string `json:"name"`
	MahasiswaID uint64 `json:"mahasiswa_id,omitempty"`
}
