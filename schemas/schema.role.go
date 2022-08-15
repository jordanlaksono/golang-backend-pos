package schemas

type SchemaRole struct {
	Role_id   string `json:"Role_id" validated:"uuid"`
	Role_name string `json:"Role_name" validate:"required,lowercase"`
}
