package schemas

type SchemaTop struct {
	Top_id     string `json:"top_id" validated:"uuid"`
	Top_kode   string `json:"top_kode" validate:"required" schema:"top_kode"`
	Top_nama   string `json:"top_nama" validate:"required" schema:"top_nama"`
	Top_jumlah string `json:"top_jumlah" validate:"required" schema:"top_jumlah"`
}
