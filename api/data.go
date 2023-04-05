package api

type LoginData struct {
	ID string `json:"id"`
	PW string `json:"pw"`
}

type SignInData struct {
	ID    string `json:"id"`
	PW    string `json:"pw"`
	Email string `json:"email"`
}

type ExecuteData struct {
	Node      string `json:"node"`
	ExecuteID string `json:executeid`
}
