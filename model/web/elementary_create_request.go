package web

type ElementaryCreateRequest struct {
	Name           string `validate:"required,min=1,max=100" json:"name"`
	Alamat         string `validate:"required,min=1,max=100" json:"alamat"`
	Tanggal_lahir  string `validate:"required,min=1,max=100" json:"tanggal_lahir"`
	Tempat_lahir   string `validate:"required,min=1,max=100" json:"tempat_lahir"`
	Jenis_kelamin  string `validate:"required,min=1,max=100" json:"jenis_Kelamin"`
	Agama          string `validate:"required,min=1,max=100" json:"agama"`
	Golongan_darah string `validate:"required,min=1,max=100" json:"golongan_darah"`
	No_telepon     string `validate:"required,min=1,max=100" json:"no_telepon"`
}
