package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	go_httpclient "github.com/gojek/heimdall/v7/httpclient"
	"github.com/raylin666/go-utils/pkg/httpclient"
	"net/http"
	"time"
)

const (
	apiUrl = "https://oapi.dingtalk.com/"

	retryCount  = 3
	httpTimeout = 3 * time.Second
)

var _ DingTalkWarning = (*dingTalkWarning)(nil)

type DingTalkWarning interface {
	SendTextMessage(data TextWarningMessageType) (*http.Response, error)
	SendLinkMessage(data LinkWarningMessageType) (*http.Response, error)
	SendMarkdownMessage(data MarkdownWarningMessageType) (*http.Response, error)
	SendAllActionCardMessage(data AllActionCardWarningMessageType) (*http.Response, error)
	SendFirstActionCardMessage(data FirstActionCardWarningMessageType) (*http.Response, error)
	SendFeedCardMessage(data FeedCardWarningMessageType) (*http.Response, error)
}

type dingTalkWarning struct {
	url         string
	accessToken string

	client httpclient.Client
}

func NewDingTalkWarning(accessToken string) DingTalkWarning {
	var dtw = new(dingTalkWarning)
	dtw.accessToken = accessToken
	dtw.url = fmt.Sprintf(apiUrl+"/robot/send?access_token=%s", dtw.accessToken)
	dtw.client = httpclient.NewClient(
		go_httpclient.WithHTTPTimeout(httpTimeout),
		go_httpclient.WithRetryCount(retryCount),
	)
	return dtw
}

type textWarningMessageType struct {
	TextWarningMessageType
	Msgtype string `json:"msgtype"`
}

type TextWarningMessageType struct {
	At struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func (dtw *dingTalkWarning) withHeader() http.Header {
	var headers = http.Header{}
	headers.Add("Content-Type", "application/json")
	return headers
}

func (dtw *dingTalkWarning) SendTextMessage(data TextWarningMessageType) (*http.Response, error) {
	var info = new(textWarningMessageType)
	info.Msgtype = "text"
	info.TextWarningMessageType = data
	jsondata, _ := json.Marshal(info)
	fmt.Println(bytes.NewBuffer(jsondata))
	return dtw.client.POST(dtw.url, bytes.NewBuffer(jsondata), dtw.withHeader())
}

type linkWarningMessageType struct {
	LinkWarningMessageType
	Msgtype string `json:"msgtype"`
}

type LinkWarningMessageType struct {
	Msgtype string `json:"msgtype"`
	Link    struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicUrl     string `json:"picUrl"`
		MessageUrl string `json:"messageUrl"`
	} `json:"link"`
}

func (dtw *dingTalkWarning) SendLinkMessage(data LinkWarningMessageType) (*http.Response, error) {
	var info = new(linkWarningMessageType)
	info.Msgtype = "link"
	info.LinkWarningMessageType = data
	jsondata, _ := json.Marshal(info)
	return dtw.client.POST(dtw.url, bytes.NewBuffer(jsondata), dtw.withHeader())
}

type markdownWarningMessageType struct {
	MarkdownWarningMessageType
	Msgtype string `json:"msgtype"`
}

type MarkdownWarningMessageType struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

func (dtw *dingTalkWarning) SendMarkdownMessage(data MarkdownWarningMessageType) (*http.Response, error) {
	var info = new(markdownWarningMessageType)
	info.Msgtype = "markdown"
	info.MarkdownWarningMessageType = data
	jsondata, _ := json.Marshal(info)
	return dtw.client.POST(dtw.url, bytes.NewBuffer(jsondata), dtw.withHeader())
}

type allActionCardWarningMessageType struct {
	AllActionCardWarningMessageType
	Msgtype string `json:"msgtype"`
}

type AllActionCardWarningMessageType struct {
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation string `json:"btnOrientation"`
		SingleTitle    string `json:"singleTitle"`
		SingleURL      string `json:"singleURL"`
	} `json:"actionCard"`
	Msgtype string `json:"msgtype"`
}

func (dtw *dingTalkWarning) SendAllActionCardMessage(data AllActionCardWarningMessageType) (*http.Response, error) {
	var info = new(allActionCardWarningMessageType)
	info.Msgtype = "actionCard"
	info.AllActionCardWarningMessageType = data
	jsondata, _ := json.Marshal(info)
	return dtw.client.POST(dtw.url, bytes.NewBuffer(jsondata), dtw.withHeader())
}

type firstActionCardWarningMessageType struct {
	FirstActionCardWarningMessageType
	Msgtype string `json:"msgtype"`
}

type FirstActionCardWarningMessageType struct {
	Msgtype    string `json:"msgtype"`
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation string `json:"btnOrientation"`
		Btns           []struct {
			Title     string `json:"title"`
			ActionURL string `json:"actionURL"`
		} `json:"btns"`
	} `json:"actionCard"`
}

func (dtw *dingTalkWarning) SendFirstActionCardMessage(data FirstActionCardWarningMessageType) (*http.Response, error) {
	var info = new(firstActionCardWarningMessageType)
	info.Msgtype = "actionCard"
	info.FirstActionCardWarningMessageType = data
	jsondata, _ := json.Marshal(info)
	return dtw.client.POST(dtw.url, bytes.NewBuffer(jsondata), dtw.withHeader())
}

type feedCardWarningMessageType struct {
	FeedCardWarningMessageType
	Msgtype string `json:"msgtype"`
}

type FeedCardWarningMessageType struct {
	Msgtype  string `json:"msgtype"`
	FeedCard struct {
		Links []struct {
			Title      string `json:"title"`
			MessageURL string `json:"messageURL"`
			PicURL     string `json:"picURL"`
		} `json:"links"`
	} `json:"feedCard"`
}

func (dtw *dingTalkWarning) SendFeedCardMessage(data FeedCardWarningMessageType) (*http.Response, error) {
	var info = new(feedCardWarningMessageType)
	info.Msgtype = "feedCard"
	info.FeedCardWarningMessageType = data
	jsondata, _ := json.Marshal(info)
	return dtw.client.POST(dtw.url, bytes.NewBuffer(jsondata), dtw.withHeader())
}
