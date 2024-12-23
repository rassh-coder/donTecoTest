package handler

import (
	"donTecoTest/pkg/models"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"io"
	"log"
	"net/http"
)

// Handler поиска сотрудника по имени
func (h *Handler) FindEmployeeByName(c *gin.Context) {
	var inputs models.EmployeeInputFields
	jDB, _ := io.ReadAll(c.Request.Body)

	if err := json.Unmarshal(jDB, &inputs); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"data":   "invalid json body",
		})
		return
	}

	if inputs.Name == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"data":   "required field 'name' is missing",
		})
		return
	}
	emp, err := h.Service.Employee.FindByName(inputs.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status": false,
				"data":   "record is not found",
			})
			return
		}
		log.Printf("can't find employee: %s", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": false,
			"data":   "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   emp,
	})
	return
}

// Handler списка сотрудников
func (h *Handler) GetListEmployee(c *gin.Context) {
	var inputs struct {
		Limit  uint `json:"limit"`
		Offset uint `json:"offset"`
	}
	jDB, _ := io.ReadAll(c.Request.Body)

	if err := json.Unmarshal(jDB, &inputs); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"data":   "invalid json body",
		})
		return
	}

	emps, err := h.Service.Employee.GetList(inputs.Limit, inputs.Offset)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status": false,
				"data":   "records is not found",
			})
			return
		}
		log.Printf("can't fetch employees: %s", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": false,
			"data":   "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   emps,
	})
	return
}
