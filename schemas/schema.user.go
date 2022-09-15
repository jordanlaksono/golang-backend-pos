package schemas

type SchemaUser struct {
	User_id        string `json:"user_id" validated:"uuid"`
	User_username  string `json:"user_username" validate:"required" schema:"user_username"`
	User_password  string `json:"user_password" validate:"required,gte=8" schema:"user_password"`
	UserKaryawanID string `json:"user_karyawan_id" validate:"required,uuid" schema:"supplier_id"`
}
