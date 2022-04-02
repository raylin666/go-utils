package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/raylin666/go-utils/httpclient"
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

func NewRobot(accessToken string, opts ...httpclient.Options) Robot {
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

func (r *robot) SendTextMessage(data RobotTextMessageType) (*http.Response, error) {
	var message = new(robotTextMessageType)
	message.Msgtype = "text"
	message.RobotTextMessageType = data
	jsondata, _ := json.Marshal(message)
	return r.client.POST(r.api, bytes.NewBuffer(jsondata), r.withHeader())
}

func (r *robot) SendLinkMessage(data RobotLinkMessageType) (*http.Response, error) {
	var message = new(robotLinkMessageType)
	message.Msgtype = "link"
	message.RobotLinkMessageType = data
	jsondata, _ := json.Marshal(message)
	return r.client.POST(r.api, bytes.NewBuffer(jsondata), r.withHeader())
}

func (r *robot) SendMarkdownMessage(data RobotMarkdownMessageType) (*http.Response, error) {
	var message = new(robotMarkdownMessageType)
	message.Msgtype = "markdown"
	message.RobotMarkdownMessageType = data
	jsondata, _ := json.Marshal(message)
	return r.client.POST(r.api, bytes.NewBuffer(jsondata), r.withHeader())
}

func (r *robot) SendAllActionCardMessage(data RobotAllActionCardMessageType) (*http.Response, error) {
	var message = new(robotAllActionCardMessageType)
	message.Msgtype = "actionCard"
	message.RobotAllActionCardMessageType = data
	jsondata, _ := json.Marshal(message)
	return r.client.POST(r.api, bytes.NewBuffer(jsondata), r.withHeader())
}

func (r *robot) SendFirstActionCardMessage(data RobotFirstActionCardMessageType) (*http.Response, error) {
	var message = new(robotFirstActionCardMessageType)
	message.Msgtype = "actionCard"
	message.RobotFirstActionCardMessageType = data
	jsondata, _ := json.Marshal(message)
	return r.client.POST(r.api, bytes.NewBuffer(jsondata), r.withHeader())
}

func (r *robot) SendFeedCardMessage(data RobotFeedCardMessageType) (*http.Response, error) {
	var message = new(robotFeedCardMessageType)
	message.Msgtype = "feedCard"
	message.RobotFeedCardMessageType = data
	jsondata, _ := json.Marshal(message)
	return r.client.POST(r.api, bytes.NewBuffer(jsondata), r.withHeader())
}
