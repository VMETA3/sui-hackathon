package request

type RequestContentType string

const (
	Json    RequestContentType = "application/json"
	FormUrl RequestContentType = "application/x-www-form-urlencoded"
)
