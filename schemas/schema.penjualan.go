package schemas

import "time"

type SchemaPenjualan struct {
	Penjualan_id                string    `json:"penjualan_id" validated:"uuid"`
	Penjualan_karyawan_id       string    `json:"penjualan_karyawan_id" validate:"required,uuid" schema:"penjualan_karyawan_id"`
	Penjualan_pelanggan_id      string    `json:"penjualan_pelanggan_id" validate:"required,uuid" schema:"penjualan_pelanggan_id"`
	Penjualan_no_faktur         string    `json:"penjualan_no_faktur" validate:"required" schema:"penjualan_no_faktur"`
	Penjualan_no_nota           string    `json:"penjualan_no_nota" validate:"required" schema:"penjualan_no_nota"`
	Penjualan_top_id            string    `json:"penjualan_top_id" validate:"required,uuid" schema:"penjualan_top_id"`
	Penjualan_metode_pembayaran string    `json:"penjualan_metode_pembayaran" validate:"required" schema:"Penjualan_metode_pembayaran"`
	Penjualan_diskon_persen     string    `json:"penjualan_diskon_persen" validate:"required"`
	Penjualan_diskon_rupiah     string    `json:"penjualan_diskon_rupiah" validate:"required"`
	Penjualan_total_harga       string    `json:"penjualan_total_harga" validate:"required"`
	Penjualan_biaya             string    `json:"penjualan_biaya" validate:"required"`
	Penjualan_biaya_debit       string    `json:"penjualan_biaya_debit" validate:"required"`
	Penjualan_selisih           string    `json:"penjualan_selisih" validate:"required"`
	Penjualan_potongan          string    `json:"penjualan_potongan" validate:"required"`
	Penjualan_biaya_kirim       string    `json:"penjualan_biaya_kirim" validate:"required"`
	Penjualan_jatuh_tempo       time.Time `json:"penjualan_jatuh_tempo"`
}
