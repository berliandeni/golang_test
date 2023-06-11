package model

type Konsumen struct {
	Id           int    `gorm:"primary_key" json:"id"`
	NIK          string `json:"nik"`
	FullName     string `json:"full_name"`
	LegalName    string `json:"Legal_name"`
	TempatLahir  string `json:"Tempat_lahir"`
	TanggalLahir string `json:"Tanggal_lahir"`
	Gaji         string `json:"gaji"`
	FotoKTP      string `json:"foto_ktp"`
	FotoSelfie   string `json:"foto_selfie"`
}

type KonsumenResp struct {
	Konsumen      Konsumen      `json:"konsumen"`
	PinjamanLimit PinjamanLimit `json:"pinjaman_limit"`
}

func (Konsumen) TableName() string {
	return "konsumen"
}
