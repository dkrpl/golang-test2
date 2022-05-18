package pasien

type (
	// CityRequest is the request body for the city service
	Input struct {
		Nama_pasien string `bson:"nama_pasien"`
		Alamat      string `bson:"alamat"`
		No_rumah    int    `json:"no_rumah"`
		Kecamatan   string `bson:"kecamatan"`
		Kabupaten   string `bson:"kabupaten"`
	}
	Edit struct {
		Nama_pasien string `bson:"nama_pasien"`
		Alamat      string `bson:"alamat"`
		No_rumah    int    `json:"no_rumah"`
		Kecamatan   string `bson:"kecamatan"`
		Kabupaten   string `bson:"kabupaten"`
	}
	AddKamar struct {
		NamaKamar string `json:"namaKamar"`
		Nama      string `json:"nama"`
		Path      string `json:"path"`
	}
)
