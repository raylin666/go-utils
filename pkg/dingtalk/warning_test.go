package dingtalk

import (
	"io/ioutil"
	"testing"
)

func TestSendTextMessage(t *testing.T) {
	var dt = NewDingTalkWarning("")
	tmt := TextWarningMessageType{}
	tmt.At.AtMobiles = []string{"xxxxxx"}
	tmt.Text.Content = "测试通知"
	resp, err := dt.SendTextMessage(tmt)
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