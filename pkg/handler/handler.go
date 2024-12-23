package handler

import (
	"donTecoTest/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Hello route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": true,
			"data":   "hello",
		})
	})

	// Можно сделать GET запросом, привычнее поиск через POST и json body
	router.POST("/employee/get-by-name", h.FindEmployeeByName)
	router.POST("/employee/get-list", h.GetListEmployee)
	return router
}
