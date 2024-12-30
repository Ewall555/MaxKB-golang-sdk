# MaxKB SDK for Go

![version](https://img.shields.io/badge/version-v1-green)

使用 Golang 开发的 MaxKB SDK，简单、易用。

## 文档 && 例子

## 快速开始

下面是与应用对话的的例子：

```go
package main

import (
	"fmt"

	mk "github.com/Ewall555/MaxKB-golang-sdk/maxkb"            // 引入包
	mkreq "github.com/Ewall555/MaxKB-golang-sdk/api/request"   // 请求参数
	mkresp "github.com/Ewall555/MaxKB-golang-sdk/api/response" // 返回参数
)

func main() {
	var baseURL = "https://maxkb.example.com"
	var apiKey = "application-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	// 初始化
	mkClient := mk.New(baseURL, apiKey)
	// 获取应用相关信息
	profileresp, err := mkClient.ApplicationChat.Profile()
	if err != nil {
		panic(err)
	}
	// 获取会话id,根据应用id
	chatid, err := mkClient.ApplicationChat.ChatOpenByApplication_id(profileresp.ID)
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
		fmt.Printf("%+v", data.Content)
	}
	// 对话
	resp, err := mkClient.ApplicationChat.Chat_messageByChat_id(req, chatid, streamCallback)
	if err != nil {
		panic(err)
	}
	// 非流式响应
	if resp != nil {
		fmt.Printf("%+v", resp.Content)
	}
}

```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
