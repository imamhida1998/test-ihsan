package request

type Login struct {
	NoHp     string `json:"no_hp"`
	Password string `json:"password"`
}
