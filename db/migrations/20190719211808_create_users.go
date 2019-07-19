package migrations

import (
	"github.com/aidanlloydtucker/wso-backend-go-demo/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var CreateUsers20190719211808 = &gormigrate.Migration{
	ID: "20190719211808_create_users",
	Migrate: func(tx *gorm.DB) error {
		// It's a good pratice to copy the struct inside the function,
		// so side effects are prevented if the original struct changes during the time.
		// But, when the table already exists, it just adds new fields as columns, so just have a struct
		// with those fields.
		type User struct {
			models.BaseSchema
			Type           string
			Name           string
			CellPhone      string
			CampusPhoneExt string
			UnixID         string
			WilliamsEmail  string
			Title          string
			Visible        bool
			ClassYear      int `gorm:"size:4"`

			// Equivalent to belongs_to Department
			DepartmentID int

			DormVisible bool `gorm:"DEFAULT:true"`
			HomeTown    string
			HomeZip     string
			HomePhone   string
			HomeState   string
			HomeCountry string
			HomeVisible bool `gorm:"DEFAULT:true"`

			Major                     string
			SUBox                     string
			Entry                     string
			Admin                     bool `gorm:"DEFAULT:false"`
			FactrakAdmin              bool `gorm:"DEFAULT:false"`
			HasAcceptedFactrakPolicy  bool `gorm:"DEFAULT:false"`
			HasAcceptedDormtrakPolicy bool `gorm:"DEFAULT:false"`

			Pronoun              string
			AtWilliams           bool `gorm:"DEFAULT:true"`
			OffCycle             bool `gorm:"DEFAULT:false"`
			FactrakSurveyDeficit int

			OptOutEphcatch      bool `gorm:"DEFAULT:false"`
			EphcatchEligibility bool `gorm:"DEFAULT:false"`
		}
		return tx.AutoMigrate(&User{}).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.DropTable("users").Error
	},
}
