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

type ExecuteProcessData struct {
	Node      string `json:"node"`
	ExecuteID string `json:"executeid"`
}

type StopProcessData struct {
	PID uint64 `json:"pid"`
}

type AddNodeData struct {
	IP   string `json:"ip"`
	Node string `json:"node"`
}
