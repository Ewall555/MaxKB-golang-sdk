package request

type Chat_messagePostRequest struct {
	Message string `json:"message"`
	ReChat  bool   `json:"re_chat"`
	Stream  bool   `json:"stream"`
}

type ChatCompletionsRequest struct {
	Messages []Message `json:"messages"`
	ChatID   string    `json:"chat_id,omitempty"`
	ReChat   bool      `json:"re_chat"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}
