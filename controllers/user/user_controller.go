package user

import (
	"errors"
	"github.com/aidanlloydtucker/wso-backend-go-demo/controllers"
	"github.com/aidanlloydtucker/wso-backend-go-demo/controllers/user/generated"
	base "github.com/aidanlloydtucker/wso-backend-go-demo/lib/generate/api"
	"github.com/aidanlloydtucker/wso-backend-go-demo/models"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type Controller struct {
	controllers.BaseController
	userModel *models.UserModel
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		userModel: &models.UserModel{
			BaseModel: models.BaseModel{
				DB: db,
			},
		},
	}
}

func (c *Controller) FetchAllUsers(ctx *base.Context) (*generated.FetchAllUsersResp, *base.Error) {
	var users []models.User
	err := c.userModel.GetAllUsers(&users)

	if err != nil {
		return nil, base.InternalServerError(err)
	}

	return generated.ToFetchAllUsersResp(users), nil
}

func (t *Controller) GetUser(userID string, ctx *base.Context) (*generated.GetUserResp, *base.Error) {
	var userIDParsed uint
	var err error

	if userID == "me" {
		userIDParsed, err = ctx.GetUserIDSelf()
		if err != nil {
			return nil, base.InternalServerError(err)
		}
	} else {
		userIDToInt, err := strconv.Atoi(userID)
		if err != nil {
			return nil, base.ErrToError(http.StatusBadRequest, errors.New("could not parse user id"))
		}
		userIDParsed = uint(userIDToInt)
	}

	var user models.User
	err = t.userModel.GetUserByID(userIDParsed, &user)
	if err != nil {
		return nil, base.InternalServerError(err)
	}

	return generated.ToGetUserResp(user), nil
}
