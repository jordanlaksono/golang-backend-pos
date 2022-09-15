package schemas

type SchemaLokasi struct {
	Lokasi_id   string `json:"lokasi_id" validated:"uuid"`
	Lokasi_kode string `json:"lokasi_kode" validate:"required" schema:"lokasi_kode"`
	Lokasi_nama string `json:"lokasi_nama" validate:"required" schema:"lokasi_nama"`
}
