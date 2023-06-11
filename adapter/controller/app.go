package controller

type AppController struct {
	Konsumen interface{ Konsumen }
	Pinjaman interface{ Pinjaman }
}
