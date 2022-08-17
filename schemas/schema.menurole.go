package schemas

type SchemaMenuRole struct {
	Menu_role_id string `json:"menu_role_id" validate:"uuid"`
	MenuID       string `json:"menu_id" validate:"required,uuid"`
	Role_id      string `json:"role_id" validate:"required,uuid"`
}
