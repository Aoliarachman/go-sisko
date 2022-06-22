package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"go-sisko/exception"
	"go-sisko/helper"
	"go-sisko/model/domain"
	"go-sisko/model/web"
	"go-sisko/repository"
)

type ElementaryServiceImpl struct {
	ElementaryRepository repository.ElementaryRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewElementaryService(elementaryRepository repository.ElementaryRepository, DB *sql.DB, validate *validator.Validate) *ElementaryServiceImpl {

	return &ElementaryServiceImpl{
		ElementaryRepository: elementaryRepository,
		DB:                   DB,
		Validate:             validate,
	}
}

func (service *ElementaryServiceImpl) Create(ctx context.Context, request web.ElementaryCreateRequest) web.ElementaryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	elementary := domain.Elementary{
		Name:           request.Name,
		Alamat:         request.Alamat,
		Tanggal_lahir:  request.Tanggal_lahir,
		Tempat_lahir:   request.Tempat_lahir,
		Jenis_kelamin:  request.Jenis_kelamin,
		Agama:          request.Agama,
		Golongan_darah: request.Golongan_darah,
		No_telepon:     request.No_telepon,
	}
	elementary = service.ElementaryRepository.Save(ctx, tx, elementary)

	return helper.ToElementaryResponse(elementary)
}

func (service *ElementaryServiceImpl) Update(ctx context.Context, request web.ElementaryUpdateRequest) web.ElementaryResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	elementary, err := service.ElementaryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	elementary.Name = request.Name
	elementary.Alamat = request.Alamat
	elementary.Tanggal_lahir = request.Tanggal_lahir
	elementary.Tempat_lahir = request.Tempat_lahir
	elementary.Jenis_kelamin = request.Jenis_kelamin
	elementary.Agama = request.Agama
	elementary.Golongan_darah = request.Golongan_darah
	elementary.No_telepon = request.No_telepon
	elementary = service.ElementaryRepository.Update(ctx, tx, elementary)

	return helper.ToElementaryResponse(elementary)
}

func (service *ElementaryServiceImpl) Delete(ctx context.Context, elementaryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	elementary, err := service.ElementaryRepository.FindById(ctx, tx, elementaryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.ElementaryRepository.Delete(ctx, tx, elementary)
}

func (service *ElementaryServiceImpl) FindById(ctx context.Context, elementaryId int) web.ElementaryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	elementary, err := service.ElementaryRepository.FindById(ctx, tx, elementaryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToElementaryResponse(elementary)
}

func (service *ElementaryServiceImpl) FindAll(ctx context.Context) []web.ElementaryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	elementarys := service.ElementaryRepository.FindAll(ctx, tx)

	return helper.ToElementaryResponses(elementarys)
}
