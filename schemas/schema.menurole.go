package schemas

type SchemaMenuRole struct {
	Menu_role_id string `json:"menu_role_id" validate:"uuid"`
	MenuID       string `json:"menu_id" validate:"required,uuid"`
	Role_id      string `json:"role_id" validate:"required,uuid"`
	User_id      string `json:"user_id" validate:"required,uuid"`
	Menu_r       string `json:"menu_r" validate:"required,uuid"`
}
