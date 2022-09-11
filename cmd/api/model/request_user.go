package model

type RequestUser struct {
	ID            string `json:"id" gorm:"primaryKey"`
	UserType      string `json:"user_type" binding:"required"`
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Username      string `json:"username" binding:"required"`
	DocType       string `json:"doc_type" binding:"required"`
	DocNumber     string `json:"doc_number" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	PersonalEmail string `json:"personal_email" binding:"required"`
	DoB           string `json:"date_of_birth" binding:"required"`
	Credential    string `json:"credential"`
}
