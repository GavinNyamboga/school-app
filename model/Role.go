package model

type Role struct {
	*Base
	Name            string           `json:"name"`
	Scope           string           `gorm:"type:ENUM('GLOBAL', 'LOCAL');default:'LOCAL'" json:"scope"`
	Description     string           `json:"description"`
	RolePermissions []RolePermission `gorm:"ForeignKey:RolePermissionID" json:"rolePermissions"`
}

type RolePermission struct {
	SimpleModelBase
	RoleID       int64 `json:"roleID"`
	Role         Role  `json:"role"`
	PermissionID int64 `json:"permissionID"`
}

type Permission struct {
	SimpleModelBase
	Name               string `json:"name"`
	DisplayName        string `json:"displayName"`
	PermissionEntityID int64  `json:"permissionEntityID"`
	PermissionGroupID  int64  `json:"permissionGroupID"`
	PermissionType     string `gorm:"type:ENUM('LIST', 'CREATE', 'EDIT', 'DELETE','OTHER'); default:'OTHER'" json:"permissionType"`
}

type PermissionEntity struct {
	SimpleModelBase
	Name              string       `json:"name"`
	DisplayName       string       `json:"displayName"`
	PermissionGroupID int64        `json:"permissionGroupID"`
	Permissions       []Permission `json:"permissions"`
	Scope             string       `gorm:"type:ENUM('GLOBAL', 'LOCAL');default:'GLOBAL'" json:"scope"`
}

type PermissionGroup struct {
	SimpleModelBase
	Name               string             `json:"name"`
	DisplayName        string             `json:"displayName"`
	PermissionEntities []PermissionEntity `json:"permissionEntities"`
	Scope              string             `gorm:"type:ENUM('GLOBAL', 'LOCAL');default:'LOCAL'" json:"scope"`
}

/*func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Role{}, &RolePermission{}, &Permission{}, &PermissionEntity{}, &PermissionGroup{})
	AddForeignKey(db, &RolePermission{}, "role_id", "roles(id)", "CASCADE", "CASCADE")
	AddForeignKey(db, &RolePermission{}, "permission_id", "permissions(id)", "CASCADE", "CASCADE")
}
*/
