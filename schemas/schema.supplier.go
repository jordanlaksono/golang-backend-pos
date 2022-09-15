package schemas

type SchemaSupplier struct {
	Supplier_id       string `json:"supplier_id" validated:"uuid"`
	Supplier_kode     string `json:"supplier_kode" validate:"required" schema:"supplier_kode"`
	Supplier_nama     string `json:"supplier_nama" validate:"required" schema:"supplier_nama"`
	Supplier_alamat   string `json:"supplier_alamat" validate:"required" schema:"supplier_alamat"`
	Supplier_kota     string `json:"supplier_kota" validate:"required" schema:"supplier_kota"`
	Supplier_kode_pos string `json:"supplier_kode_pos" validate:"required" schema:"supplier_kode_pos"`
	Supplier_email    string `json:"supplier_email" validate:"required" schema:"supplier_email"`
	Supplier_no_telp  string `json:"supplier_no_telp" validate:"required" schema:"supplier_no_telp"`
	Supplier_top_id   string `json:"supplier_top_id" validate:"required,uuid" schema:"supplier_top_id"`
}
