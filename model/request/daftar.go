package request

type Daftar struct {
	Nama     string `json:"nama"`
	Nik      string `json:"nik"`
	NoHp     string `json:"no_hp"`
	Password string `json:"password"`
	NamaBank string `json:"nama_bank"`
}
