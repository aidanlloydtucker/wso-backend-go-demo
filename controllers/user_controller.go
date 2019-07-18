package controllers

import (
	"errors"
	"fmt"
	"github.com/aidanlloydtucker/wso-backend-go-demo/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type UserController struct {
	BaseController
	userModel *models.UserModel
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		userModel: &models.UserModel{
			BaseModel: models.BaseModel{
				DB: db,
			},
		},
	}
}

func (t *UserController) FetchAllUsers(c *gin.Context) {
	var users []models.User
	err := t.userModel.GetAllUsers(&users)

	if err != nil {
		t.RespondError(http.StatusInternalServerError, err, c)
		return
	}

	t.RespondOK(users, c)
}

func (t *UserController) GetUser(c *gin.Context) {
	userIDStr := c.Param("user_id")

	var userID uint

	if userIDStr == "me" {
		userID = (c.MustGet("user_id")).(uint)
	} else {
		userIDInt, err := strconv.Atoi(userIDStr)
		userID = uint(userIDInt)
		if err != nil {
			t.RespondError(http.StatusBadRequest, errors.New("could not parse user id"), c)
			return
		}
	}

	var user models.User
	err := t.userModel.GetUserByID(uint(userID), &user)
	if err != nil {
		t.RespondError(http.StatusInternalServerError, err, c)
		return
	}

	t.RespondOK(user, c)
}

func (t *UserController) UpdateUser(c *gin.Context) {
	userID, err := GetUIntParam("user_id", c)
	if err != nil {
		t.RespondError(http.StatusBadRequest, errors.New("could not parse user id"), c)
		return
	}

	if userID != GetUserID(c) {
		t.RespondError(http.StatusForbidden, errors.New("can only update self"), c)
		return
	}

	var update map[string]interface{}
	err = c.ShouldBind(&update)
	if err != nil {
		t.RespondError(http.StatusBadRequest, errors.New("could not parse user id"), c)
		return
	}

	fmt.Printf("%+v\n", update)

	err = t.userModel.UpdateUser(userID, update)
	if err != nil {
		t.RespondError(http.StatusInternalServerError, err, c)
		return
	}

	t.RespondOK(nil, c)
}
