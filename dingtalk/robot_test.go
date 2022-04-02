package dingtalk

import (
	"github.com/raylin666/go-utils/httpclient"
	"io/ioutil"
	"testing"
	"time"
)

func TestRobotSendTextMessage(t *testing.T) {
	var rb = NewRobot("1if7D3f8D", httpclient.WithHTTPTimeout(3*time.Second))
	var message = RobotTextMessageType{}
	message.At.AtMobiles = []string{"xxxxxx"}
	message.Text.Content = "测试通知"
	resp, err := rb.SendTextMessage(message)
	if err != nil {
		t.Fatal(err)
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(res))
	defer resp.Body.Close()
}