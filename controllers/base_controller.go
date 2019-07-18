package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
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

var Base = BaseController{}

func (BaseController) RespondError(code int, err error, c *gin.Context) {
	if gorm.IsRecordNotFoundError(err) {
		code = http.StatusNotFound
	}

	c.AbortWithStatusJSON(code, BaseResponse{
		Status: code,
		Error: &RespError{
			ErrorCode: code,
			Message:   err.Error(),
		},
	})

	c.Error(err)
}

func GetUIntParam(key string, ctx *gin.Context) (uint, error) {
	param := ctx.Param(key)
	paramInt, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}

	return uint(paramInt), nil
}

func GetUserID(ctx *gin.Context) uint {
	return (ctx.MustGet("user_id")).(uint)
}