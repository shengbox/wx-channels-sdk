package apis

import (
	"encoding/json"
)

// 解码订单包含的敏感数据
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/decode_order_sensitive_info.html

type ReqOrderSensitiveinfoDecode struct {
	OrderID string `json:"order_id"`
}

var _ bodyer = ReqOrderSensitiveinfoDecode{}

func (x ReqOrderSensitiveinfoDecode) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespOrderSensitiveinfoDecode struct {
	CommonResp
	AddressInfo struct {
		CityName     string  `json:"city_name"`
		CountyName   string  `json:"county_name"`
		DetailInfo   string  `json:"detail_info"`
		HouseNumber  string  `json:"house_number"`
		Lat          float64 `json:"lat"`
		Lng          float64 `json:"lng"`
		PostalCode   string  `json:"postal_code"`
		ProvinceName string  `json:"province_name"`
		TelNumber    string  `json:"tel_number"`
		UserName     string  `json:"user_name"`
	} `json:"address_info"`
}

var _ bodyer = RespOrderSensitiveinfoDecode{}

func (x RespOrderSensitiveinfoDecode) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderSensitiveinfoDecode(req ReqOrderSensitiveinfoDecode) (RespOrderSensitiveinfoDecode, error) {
	var resp RespOrderSensitiveinfoDecode
	err := c.executeWXApiPost("/channels/ec/order/sensitiveinfo/decode", req, &resp, true)
	if err != nil {
		return RespOrderSensitiveinfoDecode{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderSensitiveinfoDecode{}, bizErr
	}
	return resp, nil
}
