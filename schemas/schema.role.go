package schemas

type SchemaRole struct {
	Role_id   string `json:"role_id" validated:"uuid"`
	Role_name string `json:"role_name" validate:"required"`
}
