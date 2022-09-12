package request

type RequestUser struct {
	ID            string `json:"id,omitempty"`
	UserType      string `json:"user_type,omitempty" binding:"required"`
	FirstName     string `json:"first_name,omitempty" binding:"required"`
	LastName      string `json:"last_name,omitempty" binding:"required"`
	Username      string `json:"username,omitempty"`
	DocType       string `json:"doc_type,omitempty" binding:"required"`
	DocNumber     string `json:"doc_number,omitempty" binding:"required"`
	Phone         string `json:"phone,omitempty" binding:"required"`
	PersonalEmail string `json:"personal_email,omitempty" binding:"required"`
	DoB           string `json:"date_of_birth,omitempty" binding:"required"`
	Credential    string `json:"credential,omitempty"`
}

type RequestUserEdit struct {
	ID            string `json:"id,omitempty"`
	UserType      string `json:"user_type,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Username      string `json:"username,omitempty" binding:"required"`
	DocType       string `json:"doc_type,omitempty"`
	DocNumber     string `json:"doc_number,omitempty"`
	Phone         string `json:"phone,omitempty"`
	PersonalEmail string `json:"personal_email,omitempty"`
	DoB           string `json:"date_of_birth,omitempty"`
	Credential    string `json:"credential,omitempty"`
}
