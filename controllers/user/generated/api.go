package generated

import (
	base "github.com/aidanlloydtucker/wso-backend-go-demo/lib/generate/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID uint `json:"id"`
	Type           string `json:"type"`
	Name           string `json:"name"`
	CellPhone      string `json:"cell_phone"`
	CampusPhoneExt string `json:"campus_phone_ext"`
	UnixID         string `json:"unix_id"`
	WilliamsEmail  string `json:"williams_email"`
	Title          string `json:"title"`
	Visible        bool   `json:"visible"`
	ClassYear      int    `json:"class_year"`

	// Equivalent to belongs_to Department
	DepartmentID int        `json:"department_id"`

	DormVisible bool   `json:"dorm_visible"`
	HomeTown    string `json:"home_town"`
	HomeZip     string `json:"home_zip"`
	HomePhone   string `json:"home_phone"`
	HomeState   string `json:"home_state"`
	HomeCountry string `json:"home_country"`
	HomeVisible bool   `json:"home_visible"`

	Major                     string `json:"major"`
	SUBox                     string `json:"su_box"`
	Entry                     string `json:"entry"`
	Admin                     bool   `json:"admin"`
	FactrakAdmin              bool   `json:"factrak_admin"`
	HasAcceptedFactrakPolicy  bool   `json:"has_accepted_factrak_policy"`
	HasAcceptedDormtrakPolicy bool   `json:"has_accepted_dormtrak_policy"`

	Pronoun              string `json:"pronoun"`
	AtWilliams           bool   `json:"at_williams"`
	OffCycle             bool   `json:"off_cycle"`
	FactrakSurveyDeficit int    `json:"factrak_survey_deficit"`

	OptOutEphcatch      bool `json:"opt_out_ephcatch"`
	EphcatchEligibility bool `json:"ephcatch_eligibility"`
}

type UserAPI interface {
	FetchAllUsers(ctx *base.Context) (*FetchAllUsersResp, *base.Error)
	GetUser(userID string, ctx *base.Context) (*GetUserResp, *base.Error)
}

type UserController struct {
	api UserAPI
}

func NewUserController(api UserAPI) *UserController {
	return &UserController{
		api: api,
	}
}

type FetchAllUsersResp []User

func (c *UserController) FetchAllUsers(ctx *gin.Context) {
	payload, err := c.api.FetchAllUsers(base.NewContext(ctx))

	if err != nil {
		ctx.JSON(err.ErrorCode, base.Response{
			Status: err.ErrorCode,
			Error: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, base.Response{
		Status: http.StatusOK,
		Data: payload,
	})
}

type GetUserResp User

func (c *UserController) GetUser(ctx *gin.Context) {
	// Unwrap payload to
	userID := ctx.Param("user_id")

	payload, err := c.api.GetUser(userID, base.NewContext(ctx))

	if err != nil {
		ctx.JSON(err.ErrorCode, base.Response{
			Status: err.ErrorCode,
			Error: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, base.Response{
		Status: http.StatusOK,
		Data: payload,
	})
}

func ToFetchAllUsersResp(model interface{}) *FetchAllUsersResp {
	refVal := base.ModelToSchema(model, FetchAllUsersResp{})
	pVal := make(FetchAllUsersResp, refVal.Len())
	for i := range pVal {
		pVal[i] = refVal.Index(i).Interface().(User)
	}
	return &pVal
}

func ToGetUserResp(model interface{}) *GetUserResp {
	refVal := base.ModelToSchema(model, GetUserResp{})
	pVal := refVal.Interface().(GetUserResp)
	return &pVal
}