package api

import (
	"fmt"
	"github.com/aidanlloydtucker/wso-backend-go-demo/models"
	"testing"
)

type Schema struct {
	ID uint `json:"id" api:"BaseSchema.ID"`
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

type FetchAllSchemaResp []Schema
type GetSchemaResp Schema

func TestModelToSchema(t *testing.T) {
	s := ModelToSchema(models.User{
		BaseSchema: models.BaseSchema{
			ID: 3,
		},
		Type: "student",
		DepartmentID: 6,
		Department: models.Department{
			Name: "hi",
		},
	}, GetSchemaResp{})

	fmt.Printf("%#+v\n", s)
}

func TestModelToSchemaArray(t *testing.T) {
	users := make([]models.User, 2)
	users[0] = models.User{
		BaseSchema: models.BaseSchema{
			ID: 3,
		},
		Type: "student",
		DepartmentID: 6,
		Department: models.Department{
			Name: "hi",
		},
	}
	users[1] = models.User{
		BaseSchema: models.BaseSchema{
			ID: 5,
		},
		Type: "alumni",
		DepartmentID: 256,
		Department: models.Department{
			Name: "t346",
		},
		Name: "foo",
	}

	s := ModelToSchema(users, FetchAllSchemaResp{})

	fmt.Printf("%#+v\n", s)
}