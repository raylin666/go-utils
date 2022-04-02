package dingtalk

type robotTextMessageType struct {
	RobotTextMessageType
	Msgtype string `json:"msgtype"`
}

type RobotTextMessageType struct {
	At struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

type robotLinkMessageType struct {
	RobotLinkMessageType
	Msgtype string `json:"msgtype"`
}

type RobotLinkMessageType struct {
	Msgtype string `json:"msgtype"`
	Link    struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicUrl     string `json:"picUrl"`
		MessageUrl string `json:"messageUrl"`
	} `json:"link"`
}

type robotMarkdownMessageType struct {
	RobotMarkdownMessageType
	Msgtype string `json:"msgtype"`
}

type RobotMarkdownMessageType struct {
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

type robotAllActionCardMessageType struct {
	RobotAllActionCardMessageType
	Msgtype string `json:"msgtype"`
}

type RobotAllActionCardMessageType struct {
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation string `json:"btnOrientation"`
		SingleTitle    string `json:"singleTitle"`
		SingleURL      string `json:"singleURL"`
	} `json:"actionCard"`
	Msgtype string `json:"msgtype"`
}

type robotFirstActionCardMessageType struct {
	RobotFirstActionCardMessageType
	Msgtype string `json:"msgtype"`
}

type RobotFirstActionCardMessageType struct {
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

type robotFeedCardMessageType struct {
	RobotFeedCardMessageType
	Msgtype string `json:"msgtype"`
}

type RobotFeedCardMessageType struct {
	Msgtype  string `json:"msgtype"`
	FeedCard struct {
		Links []struct {
			Title      string `json:"title"`
			MessageURL string `json:"messageURL"`
			PicURL     string `json:"picURL"`
		} `json:"links"`
	} `json:"feedCard"`
}
