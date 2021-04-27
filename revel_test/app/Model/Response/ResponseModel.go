package Response

type JsonResponse struct {
	Status bool `json:"status"`
	Data interface{} `json:"data"`
	Error string `json:"error"`
}