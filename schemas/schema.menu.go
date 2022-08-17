package schemas

type SchemaMenu struct {
	Menu_id   string `json:"menu_id" validated:"uuid"`
	Menu_nama string `json:"menu_nama" validate:"required"`
	Menu_link string `json:"menu_link" validate:"required"`
}
