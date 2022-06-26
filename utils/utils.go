package utils

type ResponseAPI struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Com     Comment `json:"comment"`
}

type Comment struct {
	Text string `form:"text" json:"text"`
}
