package schemas

type (
	Input struct {
		Nama_kamar string `json:"nama_kamar"`
		Desc       string `json:"desc"`
	}
	Edit struct {
		ID         string `json:"_id"`
		Nama_kamar string `json:"nama_kamar"`
		Desc       string `json:"desc"`
	}
)
