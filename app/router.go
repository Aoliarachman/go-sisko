package app

import (
	"github.com/julienschmidt/httprouter"
	"go-sisko/controller"
	"go-sisko/exception"
)

func NewRouter(elementaryController controller.ElementaryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/elementarys", elementaryController.FindAll)
	router.GET("/api/elementarys/:elementaryId", elementaryController.FindById)
	router.POST("/api/elementarys", elementaryController.Create)
	router.PUT("/api/elementarys/:elementaryId", elementaryController.Update)
	router.DELETE("/api/elementarys/:elementaryId", elementaryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
