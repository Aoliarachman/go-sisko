package web

type ElementaryUpdateRequest struct {
	Id             int    `validate:"required"`
	Name           string `validate:"required,max=100,min=1" json:"name"`
	Alamat         string `validate:"required,max=100,min=1" json:"alamat"`
	Tanggal_lahir  string `validate:"required,max=100,min=1" json:"tanggal_lahir"`
	Tempat_lahir   string `validate:"required,max=100,min=1" json:"tempat_lahir"`
	Jenis_kelamin  string `validate:"required,max=100,min=1" json:"jenis_Kelamin"`
	Agama          string `validate:"required,max=100,min=1" json:"agama"`
	Golongan_darah string `validate:"required,max=100,min=1" json:"golongan_darah"`
	No_telepon     string `validate:"required,max=100,min=1" json:"no_telepon"`
}
