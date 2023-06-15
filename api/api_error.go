package api

type api_error int

const (
	API_ERR_INVALID_PARAM api_error = iota + 1
	API_ERR_DATABASE_DISCONNECT
)

var api_error_str = [...]string{
	"Invalid Parameters",
	"Database disconnect",
}

func (err api_error) string() string { return api_error_str[err-1] }
