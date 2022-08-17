package schemas

type SchemaMenu struct {
	ID        string `json:"id" validated:"uuid"`
	Menu_nama string `json:"menu_nama" validate:"required"`
	Menu_link string `json:"menu_link" validate:"required"`
}
