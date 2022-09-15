package schemas

type SchemaKaryawan struct {
	ID                string `json:"id" validated:"uuid"`
	Karyawan_nama     string `json:"karyawan_nama" validate:"required"`
	Karyawan_alamat   string `json:"karyawan_alamat" validate:"required"`
	Karyawan_kota     string `json:"karyawan_kota" validate:"required"`
	Karyawan_negara   string `json:"karyawan_negara" validate:"required"`
	Karyawan_kode_pos string `json:"karyawan_kode_pos" validate:"required"`
	Karyawan_no_telp  string `json:"karyawan_no_telp" validate:"required"`
	Karyawan_email    string `json:"karyawan_email" validate:"required"`
}
