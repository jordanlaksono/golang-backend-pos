package schemas

type SchemaUser struct {
	User_id         string `json:"user_id" validated:"uuid"`
	User_first_name string `json:"user_first_name" validate:"required,lowercase"`
	User_last_name  string `json:"user_last_name" validate:"required,lowercase"`
	User_email      string `json:"email" validate:"required,email"`
	User_password   string `json:"password" validate:"required,gte=8"`
}
