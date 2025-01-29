package memberships

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Email     string `gorm:"unique;not null"`
		Username  string `gorm:"unique;not null"`
		Password  string `gorm:"not null"`
		CreatedBy string `gorm:"not null"`
		UpdatedBy string `gorm:"not null"`
	}
)

type (
	SingUpRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
