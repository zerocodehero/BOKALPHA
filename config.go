package bokalpha

type BedResponse struct{

	
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`

}


type LoginParams struct {
	CustID   string `json:"custId"`
	CompID   string `json:"compId"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Source   string `json:"source"`
}

func Login(params LoginParams) {
	url := "https://api.bokao2o.com/auth/merchant/v2/user/login"
	params := map[string]string{
		"custId":   params.CustID,
		"compId":   params.CompID,
		"userName": params.UserName,
		"passWord": params.PassWord,
		"source":   params.Source
	}

}
