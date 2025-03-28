# MaxKB SDK for Go

![version](https://img.shields.io/badge/version-v1-green)

使用 Golang 开发的 MaxKB SDK，简单、易用。

## 文档 && 例子

## 快速开始

下面是与应用对话的原接口风格例子(支持 openai 接口风格对话)：
```bash
go get package github.com/Ewall555/MaxKB-golang-sdk
```

```go
package main

import (
	"fmt"
	mk "github.com/Ewall555/MaxKB-golang-sdk/maxkb"            // 引入包
	mkreq "github.com/Ewall555/MaxKB-golang-sdk/api/request"   // 请求参数
	mkresp "github.com/Ewall555/MaxKB-golang-sdk/api/response" // 返回参数
	mkconfig "github.com/Ewall555/MaxKB-golang-sdk/config"     // 配置参数
)

func main() {
	var baseURL = "http(s)://maxkb.example.com"
	var apiKey = "application(user)-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	// 初始化
	mkClient := mk.NewMaxKB(&mkconfig.Config{
		BaseURL: baseURL,
		ApiKey:  apiKey,
	})
	// 获取应用聊天接口
	applicationChat := mkClient.GetApplicationChat()
	// 获取应用相关信息
	profileresp, err := applicationChat.Profile()
	if err != nil {
		panic(err)
	}
	// 获取会话id,根据应用id
	chatid, err := applicationChat.ChatOpenByApplication_id(profileresp.ID)
	if err != nil {
		panic(err)
	}
	// 请求参数
	req := mkreq.Chat_messagePostRequest{
		Message: "你好，你是谁？",
		ReChat:  false,
		Stream:  true, // 是否流式返回
	}
	// 流式回调
	streamCallback := func(data *mkresp.Chat_messagePostStreamResponse) {
		fmt.Printf("%s", data.ReasoningContent) //思考过程
		fmt.Printf("%s", data.Content)          //思考结果
	}
	// 对话
	resp, err := applicationChat.Chat_messageByChat_id(req, chatid, streamCallback)
	if err != nil {
		panic(err)
	}
	// 非流式响应
	if resp != nil {
		fmt.Printf("%s", resp.Content)
	}
}

```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
