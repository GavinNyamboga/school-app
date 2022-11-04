package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type School struct {
	SimpleModelBase
	Name           string `json:"name"`
	Registration   string `json:"registration"`
	SchoolType     string `gorm:"type:ENUM('KCSE','O_LEVELS', 'KCPE', 'IGCSE')" json:"schoolType"`
	SchoolCode     string `gorm:"unique" json:"schoolCode"`
	GenderType     string `gorm:"type:ENUM('BOYS','GIRLS','MIXED')" json:"genderType"`
	BoardingStatus string `gorm:"type:ENUM('BOARDING','DAY','MIXED')" json:"boardingStatus"`
}

func GetSchoolTotalStudents(db *gorm.DB, ID int64) int64 {
	var totalStudents int64
	st := fmt.Sprintf("select count(s.id) from users s where s.school_id=%v and user_type='STUDENT' and s.deleted<>1", ID)
	db.Raw(st).Scan(&totalStudents)
	return totalStudents
}

/*func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&School{})
}
*/
