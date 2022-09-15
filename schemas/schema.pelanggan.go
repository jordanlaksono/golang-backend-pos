package schemas

type SchemaPelanggan struct {
	Pelanggan_id        string `json:"pelanggan_id" validated:"uuid"`
	Pelanggan_kode      string `json:"pelanggan_kode" validate:"required" schema:"pelanggan_kode"`
	Pelanggan_nama      string `json:"pelanggan_nama" validate:"required" schema:"pelanggan_nama"`
	Pelanggan_alamat    string `json:"pelanggan_alamat" validate:"required" schema:"pelanggan_alamat"`
	Pelanggan_no_telp_1 string `json:"pelanggan_no_telp_1" validate:"required" schema:"pelanggan_no_telp_1"`
	Pelanggan_no_telp_2 string `json:"pelanggan_no_telp_2" validate:"required" schema:"pelanggan_no_telp_2"`
	Pelanggan_email     string `json:"pelanggan_email" validate:"required" schema:"pelanggan_email"`
	Pelanggan_kota      string `json:"pelanggan_kota" validate:"required" schema:"pelanggan_kota"`
	Pelanggan_kode_pos  string `json:"pelanggan_kode_pos" validate:"required" schema:"pelanggan_kode_pos"`
	Pelanggan_catatan   string `json:"pelanggan_catatan" validate:"required" schema:"pelanggan_catatan"`
	Pelanggan_top_id    string `json:"pelanggan_top_id" validate:"required,uuid" schema:"pelanggan_top_id"`
	Pelanggan_status    string `json:"pelanggan_status" validate:"required" schema:"pelanggan_status"`
	Penjualan_saldo     string `json:"penjualan_saldo" validate:"required"`
}
