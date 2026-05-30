// Package dingtalk provides DingTalk robot message pushing utilities.
// It supports various message types: text, link, markdown, actionCard, feedCard.
package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	httpclient "github.com/raylin666/go-utils/http"
	"net/http"
)

const (
	apiUrl = "https://oapi.dingtalk.com/"
)

var _ Robot = (*robot)(nil)

type Robot interface {
	SendTextMessage(data RobotTextMessageType) (*http.Response, error)
	SendLinkMessage(data RobotLinkMessageType) (*http.Response, error)
	SendMarkdownMessage(data RobotMarkdownMessageType) (*http.Response, error)
	SendAllActionCardMessage(data RobotAllActionCardMessageType) (*http.Response, error)
	SendFirstActionCardMessage(data RobotFirstActionCardMessageType) (*http.Response, error)
	SendFeedCardMessage(data RobotFeedCardMessageType) (*http.Response, error)
}

type robot struct {
	accessToken string
	api         string
	client      httpclient.Client
}

func NewRobot(accessToken string, opts ...httpclient.ClientOptions) Robot {
	var r = new(robot)
	r.accessToken = accessToken
	r.api = fmt.Sprintf("%s/robot/send?access_token=%s", apiUrl, accessToken)
	r.client = httpclient.NewClient(opts...)
	return r
}

func (r *robot) withHeader() http.Header {
	var headers = http.Header{}
	headers.Add("Content-Type", "application/json")
	return headers
}

// sendMessage 通用消息发送方法，减少代码重复
func (r *robot) sendMessage(msgType string, data interface{}) (*http.Response, error) {
	message := map[string]interface{}{
		"msgtype": msgType,
	}
	// 根据消息类型设置对应字段
	switch msgType {
	case "text":
		message[msgType] = data
	case "link":
		message[msgType] = data
	case "markdown":
		message[msgType] = data
	case "actionCard":
		message[msgType] = data
	case "feedCard":
		message[msgType] = data
	}
	
	jsondata, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	return r.client.POST(r.api, bytes.NewBuffer(jsondata), r.withHeader())
}

func (r *robot) SendTextMessage(data RobotTextMessageType) (*http.Response, error) {
	return r.sendMessage("text", data)
}

func (r *robot) SendLinkMessage(data RobotLinkMessageType) (*http.Response, error) {
	return r.sendMessage("link", data)
}

func (r *robot) SendMarkdownMessage(data RobotMarkdownMessageType) (*http.Response, error) {
	return r.sendMessage("markdown", data)
}

func (r *robot) SendAllActionCardMessage(data RobotAllActionCardMessageType) (*http.Response, error) {
	return r.sendMessage("actionCard", data)
}

func (r *robot) SendFirstActionCardMessage(data RobotFirstActionCardMessageType) (*http.Response, error) {
	return r.sendMessage("actionCard", data)
}

func (r *robot) SendFeedCardMessage(data RobotFeedCardMessageType) (*http.Response, error) {
	return r.sendMessage("feedCard", data)
}
