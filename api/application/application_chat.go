package application

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/Ewall555/MaxKB-golang-sdk/api/constant"
	"github.com/Ewall555/MaxKB-golang-sdk/api/request"
	"github.com/Ewall555/MaxKB-golang-sdk/api/response"
	"github.com/Ewall555/MaxKB-golang-sdk/client"
)

const (
	ChatOpenPostAddr                = "/chat/open"
	Chat_messageByChat_idPostAddr   = "/chat_message/%s"
	Chat_workflowOpenPostAddr       = "/chat_workflow/open"
	ProfileGetAddr                  = "/profile"
	ApplicationByAppidGetAddr       = "/%s/application"
	ChatOpenByApplication_idGetAddr = "/%s/chat/open"
	VotePutAddr                     = "/%s/chat/%s/chat_record/%s/vote"
)

type ApplicationChat struct {
	client *client.Client
}

func NewApplicationChat(cli *client.Client) *ApplicationChat {
	return &ApplicationChat{client: cli}
}

// 对话
func (c *ApplicationChat) Chat_messageByChat_id(req request.Chat_messagePostRequest, chatid *string, streamCallback func(*response.Chat_messagePostStreamResponse)) (*response.Chat_messagePostResponse, error) {
	endpoint := constant.ApplicationPath + fmt.Sprintf(Chat_messageByChat_idPostAddr, *chatid)
	if req.Stream {
		resp, err := c.client.DoRequestStream("POST", endpoint, req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				return nil, fmt.Errorf("error: %w", err)
			}
			line = bytes.TrimSpace(line)
			if len(line) == 0 {
				continue
			}
			if !strings.HasPrefix(string(line), "data: ") {
				fmt.Printf("Unexpected line: %s\n", string(line))
				continue
			}
			line = line[6:]

			var data response.Chat_messagePostStreamResponse
			if err := json.Unmarshal(line, &data); err != nil {
				return nil, fmt.Errorf("error: %w", err)
			}

			streamCallback(&data)

			if data.IsEnd {
				break
			}
		}
		return nil, nil
	}
	var resp response.ApiResponse[response.Chat_messagePostResponse]
	err := c.client.DoRequest("POST", endpoint, req, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf("API error: %s (code: %d)", resp.Message, resp.Code)
	}
	return &resp.Data, nil
}

// 获取应用相关信息
func (c *ApplicationChat) Profile() (*response.ProfileResponse, error) {
	var resp response.ApiResponse[response.ProfileResponse]
	endpoint := constant.ApplicationPath + ProfileGetAddr
	err := c.client.DoRequest("GET", endpoint, nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf("API error: %s (code: %d)", resp.Message, resp.Code)
	}
	return &resp.Data, nil
}

// 获取会话id,根据应用id
func (c *ApplicationChat) ChatOpenByApplication_id(appid string) (*string, error) {
	var resp response.ApiResponse[string]
	endpoint := constant.ApplicationPath + fmt.Sprintf(ChatOpenByApplication_idGetAddr, appid)
	err := c.client.DoRequest("GET", endpoint, nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf("API error: %s (code: %d)", resp.Message, resp.Code)
	}
	return &resp.Data, nil
}
