package callbacks

import "encoding/json"

// 订单待发货
// 文档: https://developers.weixin.qq.com/doc/store/shop/notify/order_callback/channels_ec_order_wait_shipping.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcOrderWaitShipping{})
}

type ChannelsEcOrderWaitShipping struct {
	CreateTime   int  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	OrderInfo    struct {
		OrderID int `json:"order_id"`
	} `json:"order_info"`
}

func (ChannelsEcOrderWaitShipping) GetMessageType() string {
	return "event"
}

func (ChannelsEcOrderWaitShipping) GetEventType() string {
	return "channels_ec_order_wait_shipping"
}

func (m ChannelsEcOrderWaitShipping) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcOrderWaitShipping) ParseFromJson(data []byte) (CallbackExtraInfoInterface,error) {
	var temp ChannelsEcOrderWaitShipping
	err := json.Unmarshal(data, &temp)
	return temp, err
}
