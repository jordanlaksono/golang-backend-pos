package schemas

type SchemaSatuan struct {
	Satuan_id   string `json:"satuan_id" validated:"uuid"`
	Satuan_kode string `json:"satuan_kode" validate:"required" schema:"satuan_kode"`
	Satuan_nama string `json:"satuan_nama" validate:"required" schema:"satuan_nama"`
}
