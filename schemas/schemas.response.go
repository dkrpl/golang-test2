package schemas

type SchemaResponses struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
type SchemaResponsesList struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Count      int         `json:"count"`
	PerPage    int         `json:"perPage"`
	PageNo     int         `json:"pageNo"`
	From       int         `json:"from"`
	To         int         `json:"to"`
	LasPage    int         `json:"lastPage"`
}

type SchemaUnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}
