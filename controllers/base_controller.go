package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type BaseController struct{}

type BaseResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  *RespError  `json:"error,omitempty"`
}

type RespError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func (BaseController) RespondOK(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, BaseResponse{
		Status: http.StatusOK,
		Data:   data,
		Error:  nil,
	})
}

func (BaseController) RespondError(code int, err error, c *gin.Context) {
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, BaseResponse{
			Status: http.StatusNotFound,
			Error: &RespError{
				ErrorCode: http.StatusNotFound,
				Message:   err.Error(),
			},
		})

		return
	}

	c.JSON(code, BaseResponse{
		Status: code,
		Error: &RespError{
			ErrorCode: code,
			Message:   err.Error(),
		},
	})
}
