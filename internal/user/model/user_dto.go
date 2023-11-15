package model

// dto (data transfer object) 数据传输对象

// C2S_Login Web登录请求
type C2S_Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *C2S_Login) ToDomain() (*LoginParams, error) {
	username, err := NewUsername(c.Username)
	if err != nil {
		return nil, err
	}
	password, err := NewPassword(c.Password)
	if err != nil {
		return nil, err
	}

	return &LoginParams{
		Username: username,
		Password: password,
	}, nil
}

// S2C_Login Web登录响应
type S2C_Login struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type S2C_UserInfo struct {
	UserID   int64  `json:"user_id"`  // 用戶唯一Id
	Username string `json:"username"` // 姓名
	Gender   int    `json:"gender"`   // 性別
	Height   int    `json:"height"`   // 身高

	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type C2S_Register_Base struct {
	Username string `json:"username"` // 姓名
	Gender   int    `json:"gender"`   // 性別
	Height   int    `json:"height"`   // 身高
}
type C2S_Register struct {
	C2S_Register_Base
	Password string `json:"password"` // 用戶密碼
}

func (c *C2S_Register) ToDomain() (*RegisterParams, error) {
	username, err := NewUsername(c.Username)
	if err != nil {
		return nil, err
	}
	gender, err := NewGender(c.Gender)
	if err != nil {
		return nil, err
	}
	height, err := NewHeight(c.Height)
	if err != nil {
		return nil, err
	}
	password, err := NewPassword(c.Password)
	if err != nil {
		return nil, err
	}

	return &RegisterParams{
		Username: username,
		Password: password,
		Gender:   gender,
		Height:   height,
	}, nil
}

//
type MatchPeople struct {
	TeamIndex   int64 `json:"team_index"` // 組隊編號
	TeamMemberA C2S_Register_Base
	TeamMemberB C2S_Register_Base
}

// S2C_Login Web登录响应
type S2C_MatchPeople struct {
	//MatchPeopleList []MatchPeople `json:"match_people_list"`
	TeamIndex   int               `json:"team_index"` // 組隊編號
	TeamMemberA C2S_Register_Base `json:"teamMemberA"`
	TeamMemberB C2S_Register_Base `json:"teamMemberB"`
}

type C2S_Transfer struct {
	ToUserID string `json:"to_user_id"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
