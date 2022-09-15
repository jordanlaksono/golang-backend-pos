package schemas

type SchemaBarang struct {
	Barang_id          string `json:"barang_id" validated:"uuid"`
	Barang_kode        string `json:"barang_kode" validate:"required" schema:"barang_kode"`
	Barang_kelompok_id string `json:"barang_kelompok_id" validate:"required" schema:"barang_kelompok_id"`
	Barang_supplier_id string `json:"barang_supplier_id" validate:"required" schema:"barang_supplier_id"`
	Barang_nama        string `json:"barang_nama" validate:"required" schema:"barang_nama"`
	Barang_image       string `json:"barang_image" validate:"required" schema:"barang_image"`
	Barang_keterangan  string `json:"barang_keterangan" validate:"required" schema:"barang_keterangan"`
}
