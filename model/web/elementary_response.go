package web

type ElementaryResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Alamat         string `json:"alamat"`
	Tanggal_lahir  string `json:"tanggal_lahir"`
	Tempat_lahir   string `json:"tempat_lahir"`
	Jenis_kelamin  string `json:"jenis_kelamin"`
	Agama          string `json:"agama"`
	Golongan_darah string `json:"golongan_darah"`
	No_telepon     string `json:"no_telepon"`
}
