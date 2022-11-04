package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Base struct {
	SimpleModelBase
	SchoolID int64 `json:"SchoolID"`
}

type SimpleModelBase struct {
	ID              int64      `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time  `json:"createdAt"`
	CreatedByUserId int64      `json:"createdByUserId"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	UpdatedByUserId int64      `json:"updatedByUserId"`
	Deleted         bool       `json:"deleted"`
	DeletedAt       *time.Time `json:"deletedAt"`
	DeletedById     int64      `json:"deletedById"`
}

func AddForeignKey(DB *gorm.DB, x interface{}, field string, dest string, onDelete string, onUpdate string) {
	DB.Model(x).AddForeignKey(field, dest, onDelete, onUpdate)
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &School{}, &Intake{}, &Stream{}, &Role{}, &RolePermission{}, &Permission{},
		&PermissionEntity{}, &PermissionGroup{})
	//users FK
	AddForeignKey(db, User{}, "current_school_id", "schools(id)", "RESTRICT", "CASCADE")
	AddForeignKey(db, User{}, "current_stream_id", "streams(id)", "RESTRICT", "CASCADE")
	//Streams FK
	AddForeignKey(db, &Stream{}, "intake_id", "intakes(id)", "RESTRICT", "CASCADE")
	//Roles FK
	AddForeignKey(db, &RolePermission{}, "role_id", "roles(id)", "CASCADE", "CASCADE")
	AddForeignKey(db, &RolePermission{}, "permission_id", "permissions(id)", "CASCADE", "CASCADE")

	return db
}

func CreateModelBase(mb Base) Base {

	mb.CreatedAt = time.Now()
	//mb.CreatedByUserId = 1
	//mb.SchoolID=1

	return mb
}

func CreateSimpleModelBase(sm SimpleModelBase) SimpleModelBase {

	sm.CreatedAt = time.Now()
	//sm.CreatedByUserId = 1

	return sm
}

func UpdateModel(sm *SimpleModelBase, mb *Base) (*SimpleModelBase, *Base) {
	if sm != nil {
		sm.UpdatedAt = time.Now()
		sm.UpdatedByUserId = 1
	}

	if mb != nil {
		mb.UpdatedAt = time.Now()
		mb.UpdatedByUserId = 1
	}
	return sm, mb
}
