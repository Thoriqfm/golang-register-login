package entity

type Role struct {
	RoleID int    `json:"role_id" gorm:"type:int;primaryKey;autoIncrement"`
	Role   string `json:"role" gorm:"type:varchar(255);not null"`

	Users []User `json:"users" gorm:"foreignKey:RoleID"`
}
