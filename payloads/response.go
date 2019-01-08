package payloads

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
