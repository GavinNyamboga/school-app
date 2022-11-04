package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Base
	UserType        string         `gorm:"type:ENUM('STUDENT', 'TEACHER', 'ADMIN', 'PARENT', 'USER');default:'USER'" json:"userType"`
	Username        string         `gorm:"unique" json:"username"`
	Email           string         `json:"email"`
	Password        string         `json:"password"`
	CurrentSchoolID int64          `json:"currentSchoolID"`
	CurrentSchool   School         `gorm:"foreignKey:CurrentSchoolID" json:"currentSchool"`
	Person          Person         `gorm:"embedded"`
	StudentProfile  StudentProfile `gorm:"embedded"`
	RoleID          int64          `json:"roleID"`
	Role            Role
}

type StudentProfile struct {
	AdmNo           string `json:"admNo"`
	CurrentStreamId int64  `json:"currentStreamId"`
	CurrentStream   int64  `gorm:"" json:"currentStream"`
	GuardianPhone   string `json:"guardianPhone"`
}

type Person struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `gorm:"type:ENUM('MALE','FEMALE','NOT_SPECIFIED','N/A')" json:"gender"`
}

/*func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
	AddForeignKey(db, User{}, "current_school_id", "schools(id)", "RESTRICT", "CASCADE")
	AddForeignKey(db, User{}, "current_stream_id", "streams(id)", "RESTRICT", "CASCADE")
}*/

func GetUsers(db *gorm.DB) []User {
	var Users []User
	db.Find(&Users)
	return Users
}
func GetUser(db *gorm.DB, Id int64) *User {
	var getUser User
	db.Where("ID=?", Id).Find(&getUser)
	return &getUser
}
func (u *User) CreateUser(db *gorm.DB) *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func DeleteUser(db *gorm.DB, Id int64) User {
	var user User
	db.Where("ID=?", Id).Delete(&user)
	return user
}
