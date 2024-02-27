package api

type LoginRequestBody struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	AuthToken string `json:"auth_token"`
}
