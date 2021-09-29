package base

type Option struct {
	Text    string `json:"text"`
	Value   string `json:"value"`
	Comment string `json:"comment"`
	Color   string `json:"color"`
}

//用户
type LoginUser struct {
	Name string `json:"name,omitempty"` //名称
	Role int    `json:"role,omitempty"` //角色类型
}

//用户登录参数
type UserLoginReqeust struct {
	Name string `json:"name,omitempty"` //登录名
	Auth string `json:"auth,omitempty"` //密码
}

//Session参数
type SessionParam struct {
	User           *LoginUser `json:"user,omitempty"`           //登录用户
	IsSuperManager bool       `json:"isSuperManager,omitempty"` //是超管
	IsManager      bool       `json:"isManager,omitempty"`      //是管理员
	IsViewer       bool       `json:"isViewer,omitempty"`       //是观察员
}
