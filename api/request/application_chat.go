package request

type Chat_messagePostRequest struct {
	Message string `json:"message"`
	ReChat  bool   `json:"re_chat"`
	Stream  bool   `json:"stream"`
}
