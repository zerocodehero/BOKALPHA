/*
 @author: lynn
 @date: 2023/5/18
 @time: 23:16
*/

package api

import "github.com/zerocodehero/BOKALPHA/models"

type LoginParams struct {
	CustID   string `json:"custId"`
	CompID   string `json:"compId"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Source   string `json:"source"`
}

func (L LoginParams) Data() (data models.RequestData) {
	data.Url = "https://api.bokao2o.com/auth/merchant/v2/user/login"
	data.Params = map[string]string{
		"custId":   L.CustID,
		"compId":   L.CompID,
		"userName": L.UserName,
		"passWord": L.PassWord,
		"source":   L.Source,
	}
	data.Method = "POST"
	data.Headers = map[string]string{
		"origin":  "https://s3.boka.vc",
		"referer": "https://s3.boka.vc/",
	}
	return
}
