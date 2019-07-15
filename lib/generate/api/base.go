package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  *Error  `json:"error,omitempty"`
}

type Error struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func ErrToError(code int, err error) *Error {
	if gorm.IsRecordNotFoundError(err) {
		return &Error{
			ErrorCode: http.StatusNotFound,
			Message: err.Error(),
		}
	}

	return &Error{
		ErrorCode: code,
		Message: err.Error(),
	}
}

func InternalServerError(err error) *Error {
	if gorm.IsRecordNotFoundError(err) {
		return &Error{
			ErrorCode: http.StatusNotFound,
			Message: err.Error(),
		}
	}

	return &Error{
		ErrorCode: http.StatusInternalServerError,
		Message: err.Error(),
	}
}

type Context struct {
	gin *gin.Context
}

func NewContext(gin *gin.Context) *Context {
	return &Context{
		gin: gin,
	}
}

func (ctx *Context) GetUserIDSelf() (uint, error) {
	idItf, exists := ctx.gin.Get("id")
	if !exists {
		return 0, errors.New("userID could not be found")
	}

	id, exists := idItf.(uint)
	if !exists {
		return 0, errors.New("userID could not be converted")
	}

	return id, nil
}