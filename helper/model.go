package helper

import (
	"go-sisko/model/domain"
	"go-sisko/model/web"
)

func ToElementaryResponse(elementary domain.Elementary) web.ElementaryResponse {
	return web.ElementaryResponse{
		Id:             elementary.Id,
		Name:           elementary.Name,
		Alamat:         elementary.Alamat,
		Tanggal_lahir:  elementary.Tanggal_lahir,
		Tempat_lahir:   elementary.Tempat_lahir,
		Jenis_kelamin:  elementary.Jenis_kelamin,
		Agama:          elementary.Agama,
		Golongan_darah: elementary.Golongan_darah,
		No_telepon:     elementary.No_telepon,
	}
}

func ToElementaryResponses(elementarys []domain.Elementary) []web.ElementaryResponse {
	var elementaryResponses []web.ElementaryResponse
	for _, elementary := range elementarys {
		elementaryResponses = append(elementaryResponses, ToElementaryResponse(elementary))
	}
	return elementaryResponses
}
