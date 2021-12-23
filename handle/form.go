package handle

type LoginFrom struct {
	UserName string `json:"user_name,omitempty"`
	PassWord string `json:"pass_word,omitempty"`
}