package schemas

type SchemaKelompok struct {
	Kelompok_id   string `json:"kelompok_id" validated:"uuid"`
	Kelompok_kode string `json:"kelompok_kode" validate:"required" schema:"kelompok_kode"`
	Kelompok_nama string `json:"kelompok_nama" validate:"required" schema:"kelompok_nama"`
}
