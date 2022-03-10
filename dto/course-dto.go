package dto

type CourseUpdate struct {
	Kode        string `json:"kode"`
	Name        string `json:"name"`
	MahasiswaID uint64 `json:"mahasiswa_id,omitempty"`
}
