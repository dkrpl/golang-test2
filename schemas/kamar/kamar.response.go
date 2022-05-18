package schemas

type (
	Detail struct {
		ID         string `json:"_id" bson:"_id,omitempty"`
		Nama_kamar string `json:"nama_kamar"`
		Desc       string `json:"desc"`
	}

	List Detail

	Input_Response struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}
	Edit_Response   Input_Response
	Delete_Response Input_Response
)
