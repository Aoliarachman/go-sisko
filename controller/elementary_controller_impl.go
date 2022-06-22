package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-sisko/helper"
	"go-sisko/model/web"
	"go-sisko/service"
	"net/http"
	"strconv"
)

type ElementaryControllerImpl struct {
	ElementaryService service.ElementaryService
}

func NewElementaryController(elementaryService service.ElementaryService) ElementaryController {
	return &ElementaryControllerImpl{ElementaryService: elementaryService}
}

func (controller *ElementaryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	elementaryCreateRequest := web.ElementaryCreateRequest{}
	helper.ReadFromRequestBody(request, &elementaryCreateRequest)

	elementaryResponse := controller.ElementaryService.Create(request.Context(), elementaryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   elementaryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ElementaryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	elementaryUpdateRequest := web.ElementaryUpdateRequest{}
	helper.ReadFromRequestBody(request, &elementaryUpdateRequest)

	elementaryId := params.ByName("elementaryId")
	id, err := strconv.Atoi(elementaryId)
	helper.PanicIfError(err)

	elementaryUpdateRequest.Id = id

	elementaryResponse := controller.ElementaryService.Update(request.Context(), elementaryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   elementaryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ElementaryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	elementaryId := params.ByName("elementaryId")
	id, err := strconv.Atoi(elementaryId)
	helper.PanicIfError(err)

	controller.ElementaryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ElementaryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	elementaryId := params.ByName("elementaryId")
	id, err := strconv.Atoi(elementaryId)
	helper.PanicIfError(err)

	elementarysResponse := controller.ElementaryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   elementarysResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ElementaryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	elementaryResponse := controller.ElementaryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   elementaryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
