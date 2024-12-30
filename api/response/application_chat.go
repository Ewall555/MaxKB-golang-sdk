package response

// Chat_message, stream false

type Chat_messagePostResponse struct {
	ChatID           string `json:"chat_id"`
	ID               string `json:"id"`
	Operate          bool   `json:"operate"`
	Content          string `json:"content"`
	IsEnd            bool   `json:"is_end"`
	CompletionTokens int    `json:"completion_tokens"`
	PromptTokens     int    `json:"prompt_tokens"`
}

// Chat_message, stream true

type Chat_messagePostStreamResponse struct {
	ChatID       string     `json:"chat_id"`
	ChatRecordID string     `json:"chat_record_id"`
	Operate      bool       `json:"operate"`
	Content      string     `json:"content"`
	NodeID       string     `json:"node_id"`
	UpNodeIDList []string   `json:"up_node_id_list"`
	IsEnd        bool       `json:"is_end"`
	Usage        TokenUsage `json:"usage"`
	NodeIsEnd    bool       `json:"node_is_end"`
	ViewType     string     `json:"view_type"`
	NodeType     string     `json:"node_type"`
}

type TokenUsage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Profile

type ProfileResponse struct {
	ID                     string                 `json:"id"`
	Name                   string                 `json:"name"`
	Desc                   string                 `json:"desc"`
	Prologue               string                 `json:"prologue"`
	DialogueNumber         int                    `json:"dialogue_number"`
	Icon                   string                 `json:"icon"`
	Type                   string                 `json:"type"`
	STTModelID             *string                `json:"stt_model_id"`
	TTSModelID             *string                `json:"tts_model_id"`
	STTModelEnable         bool                   `json:"stt_model_enable"`
	TTSModelEnable         bool                   `json:"tts_model_enable"`
	TTSType                string                 `json:"tts_type"`
	FileUploadEnable       bool                   `json:"file_upload_enable"`
	FileUploadSetting      map[string]interface{} `json:"file_upload_setting"`
	WorkFlow               map[string]interface{} `json:"work_flow"`
	ShowSource             bool                   `json:"show_source"`
	ShowHistory            bool                   `json:"show_history"`
	Draggable              bool                   `json:"draggable"`
	ShowGuide              bool                   `json:"show_guide"`
	Avatar                 string                 `json:"avatar"`
	FloatIcon              string                 `json:"float_icon"`
	Authentication         bool                   `json:"authentication"`
	AuthenticationType     string                 `json:"authentication_type"`
	Disclaimer             bool                   `json:"disclaimer"`
	DisclaimerValue        string                 `json:"disclaimer_value"`
	CustomTheme            CustomTheme            `json:"custom_theme"`
	UserAvatar             string                 `json:"user_avatar"`
	FloatLocation          FloatLocation          `json:"float_location"`
	MultipleRoundsDialogue bool                   `json:"multiple_rounds_dialogue"`
}

type CustomTheme struct {
	ThemeColor      string `json:"theme_color"`
	HeaderFontColor string `json:"header_font_color"`
}

type FloatLocation struct {
	X FloatLocationValue `json:"x"`
	Y FloatLocationValue `json:"y"`
}

type FloatLocationValue struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

// ChatCompletions, stream false
type ChatCompletionsResponse struct {
	ID                string   `json:"id"`
	Choices           []Choice `json:"choices"`
	Created           int      `json:"created"`
	Model             string   `json:"model"`
	Object            string   `json:"object"`
	ServiceTier       *string  `json:"service_tier"`
	SystemFingerprint *string  `json:"system_fingerprint"`
	Usage             Usage    `json:"usage"`
}

type Choice struct {
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
	Logprobs     *string `json:"logprobs"`
	Message      Message `json:"message"`
	ChatID       string  `json:"chat_id"`
}

type Message struct {
	Content      string  `json:"content"`
	Refusal      *string `json:"refusal"`
	Role         string  `json:"role"`
	Audio        *string `json:"audio"`
	FunctionCall *string `json:"function_call"`
	ToolCalls    *string `json:"tool_calls"`
}

type Usage struct {
	CompletionTokens        int     `json:"completion_tokens"`
	PromptTokens            int     `json:"prompt_tokens"`
	TotalTokens             int     `json:"total_tokens"`
	CompletionTokensDetails *string `json:"completion_tokens_details"`
	PromptTokensDetails     *string `json:"prompt_tokens_details"`
}

// ChatCompletions, stream true

type ChatCompletionsStreamResponse struct {
	ID                string         `json:"id"`
	Choices           []StreamChoice `json:"choices"`
	Created           int            `json:"created"`
	Model             string         `json:"model"`
	Object            string         `json:"object"`
	ServiceTier       *string        `json:"service_tier"`
	SystemFingerprint *string        `json:"system_fingerprint"`
	Usage             Usage          `json:"usage"`
}

type StreamChoice struct {
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
	Logprobs     *string `json:"logprobs"`
	Delta        Delta   `json:"delta"`
}

type Delta struct {
	Content      string  `json:"content"`
	Role         string  `json:"role"`
	ToolCalls    *string `json:"tool_calls"`
	FunctionCall *string `json:"function_call"`
	Refusal      *string `json:"refusal"`
	ChatID       string  `json:"chat_id"`
}
