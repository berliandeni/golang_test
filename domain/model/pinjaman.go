package model

type PinjamanLimit struct {
	Id         int     `gorm:"primary_key" json:"id"`
	KonsumenId int     `json:"konsumen_id"`
	Tenor1     float64 `json:"tenor1"`
	Tenor2     float64 `json:"tenor2"`
	Tenor3     float64 `json:"tenor3"`
	Tenor4     float64 `json:"tenor4"`
}

type PinjamanTrx struct {
	Id            int     `gorm:"primary_key" json:"id"`
	NoKontak      string  `json:"no_kontak"`
	OTR           float64 `json:"otr"`
	NamaAsset     string  `json:"nama_asset"`
	AdminFee      float64 `json:"admin_fee"`
	JumlahBunga   float64 `json:"jumlah_bunga"`
	JumlahCicilan float64 `json:"jumlah_cicilan"`
}

type PinjamanTrxReq struct {
	KonsumenId   int          `json:"konsumen_id"`
	Tenor        int          `json:"tenor"`
	PinjamanBody *PinjamanTrx `json:"pinjaman_body"`
}

func (PinjamanLimit) TableName() string {
	return "pinjaman_limit"
}

func (PinjamanTrx) TableName() string {
	return "pinjaman_trx"
}
