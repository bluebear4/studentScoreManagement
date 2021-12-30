package handle

type LoginFrom struct {
	UserName string `json:"user_name,omitempty"`
	PassWord string `json:"pass_word,omitempty"`
}

type RegisterFrom struct {
	UserName   string `json:"user_name,omitempty"`
	PassWord   string `json:"pass_word,omitempty"`
	RoleID     string `json:"role_id,omitempty"`
	VerifyCode string `json:"verify_code,omitempty"`
}
