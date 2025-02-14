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
	ChatID           string    `json:"chat_id"`
	ChatRecordID     string    `json:"chat_record_id"`
	Operate          bool      `json:"operate"`
	Content          string    `json:"content"`
	NodeID           string    `json:"node_id"`
	UpNodeIDList     []string  `json:"up_node_id_list"`
	IsEnd            bool      `json:"is_end"`
	Usage            BaseUsage `json:"usage"`
	NodeType         string    `json:"node_type"`
	RuntimeNodeID    string    `json:"runtime_node_id"`
	ViewType         *string   `json:"view_type"`
	ChildNode        ChildNode `json:"child_node"`
	NodeIsEnd        bool      `json:"node_is_end"`
	RealNodeID       string    `json:"real_node_id"`
	ReasoningContent string    `json:"reasoning_content"`
}

type BaseUsage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChildNode struct {
	RuntimeNodeID *string    `json:"runtime_node_id"`
	ChatRecordID  *string    `json:"chat_record_id"`
	ChildNode     *ChildNode `json:"child_node"`
}

// Profile
type ProfileResponse struct {
	ID                     string            `json:"id"`
	Name                   string            `json:"name"`
	Desc                   string            `json:"desc"`
	Prologue               string            `json:"prologue"`
	DialogueNumber         int               `json:"dialogue_number"`
	Icon                   string            `json:"icon"`
	Type                   string            `json:"type"`
	STTModelID             *string           `json:"stt_model_id"`
	TTSModelID             *string           `json:"tts_model_id"`
	STTModelEnable         bool              `json:"stt_model_enable"`
	TTSModelEnable         bool              `json:"tts_model_enable"`
	TTSType                string            `json:"tts_type"`
	TTSAutoplay            bool              `json:"tts_autoplay"`
	STTAutosend            bool              `json:"stt_autosend"`
	FileUploadEnable       bool              `json:"file_upload_enable"`
	FileUploadSetting      FileUploadSetting `json:"file_upload_setting"`
	WorkFlow               WorkFlow          `json:"work_flow"`
	ShowSource             bool              `json:"show_source"`
	Language               *string           `json:"language"`
	ShowHistory            bool              `json:"show_history"`
	Draggable              bool              `json:"draggable"`
	ShowGuide              bool              `json:"show_guide"`
	Avatar                 string            `json:"avatar"`
	FloatIcon              string            `json:"float_icon"`
	Authentication         bool              `json:"authentication"`
	AuthenticationType     string            `json:"authentication_type"`
	Disclaimer             bool              `json:"disclaimer"`
	DisclaimerValue        string            `json:"disclaimer_value"`
	CustomTheme            CustomTheme       `json:"custom_theme"`
	UserAvatar             string            `json:"user_avatar"`
	FloatLocation          FloatLocation     `json:"float_location"`
	MultipleRoundsDialogue bool              `json:"multiple_rounds_dialogue"`
}

type FileUploadSetting struct {
	Audio     bool `json:"audio"`
	Image     bool `json:"image"`
	Video     bool `json:"video"`
	Document  bool `json:"document"`
	MaxFiles  int  `json:"maxFiles"`
	FileLimit int  `json:"fileLimit"`
}

type WorkFlow struct {
	Edges []Edge `json:"edges"`
	Nodes []Node `json:"nodes"`
}

type Edge struct {
	ID             string         `json:"id"`
	Type           string         `json:"type"`
	EndPoint       Point          `json:"endPoint"`
	PointsList     []Point        `json:"pointsList"`
	Properties     map[string]any `json:"properties"` // 使用 map 表示空对象
	StartPoint     Point          `json:"startPoint"`
	SourceNodeID   string         `json:"sourceNodeId"`
	TargetNodeID   string         `json:"targetNodeId"`
	SourceAnchorID string         `json:"sourceAnchorId"`
	TargetAnchorID string         `json:"targetAnchorId"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Node struct {
	X          float64    `json:"x"`
	Y          float64    `json:"y"`
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	Config             map[string]any `json:"config"`
	Height             float64        `json:"height"`
	ShowNode           bool           `json:"showNode"`
	StepName           string         `json:"stepName"`
	NodeData           NodeData       `json:"node_data"`
	InputFieldList     []Field        `json:"input_field_list"`
	APIInputFieldList  []Field        `json:"api_input_field_list"`
	UserInputFieldList []Field        `json:"user_input_field_list"`
	Condition          *string        `json:"condition"`
}

type NodeData struct {
	Desc               string             `json:"desc"`
	Name               string             `json:"name"`
	Prologue           string             `json:"prologue"`
	TtsType            string             `json:"tts_type"`
	Prompt             string             `json:"prompt"`
	System             string             `json:"system"`
	ModelID            string             `json:"model_id"`
	IsResult           bool               `json:"is_result"`
	DialogueType       string             `json:"dialogue_type"`
	ModelSetting       ModelSetting       `json:"model_setting"`
	DialogueNumber     int                `json:"dialogue_number"`
	ModelParamsSetting ModelParamsSetting `json:"model_params_setting"`
}

type ModelSetting struct {
	ReasoningContentEnd    string `json:"reasoning_content_end"`
	ReasoningContentStart  string `json:"reasoning_content_start"`
	ReasoningContentEnable bool   `json:"reasoning_content_enable"`
}

type Field struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type ModelParamsSetting struct {
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
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
	BaseUsage
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
