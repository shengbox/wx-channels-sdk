package callbacks

import "encoding/json"

// 客服消息事件
// 文档: https://developers.weixin.qq.com/doc/store/shop/notify/kf_callback/kf_msg_event.html

func init() {
	//添加可解析的回调事件
	supportCallback(CommkfSendMsgToKf{})
}

type CommkfSendMsgToKf struct {
	CreateTime   int    `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	MsgEvent     string `json:"msg_event"`
	MsgID        string `json:"msg_id"`
	Text         struct {
		Content string `json:"content"`
	} `json:"text"`
}

func (CommkfSendMsgToKf) GetMessageType() string {
	return "event"
}

func (CommkfSendMsgToKf) GetEventType() string {
	return "commkf_send_msg_to_kf"
}

func (m CommkfSendMsgToKf) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (CommkfSendMsgToKf) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp CommkfSendMsgToKf
	err := json.Unmarshal(data, &temp)
	return temp, err
}
